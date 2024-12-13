package remotesigning

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func GetPaymentHashFromScript(scriptHex string) (*string, error) {
	pattern := `OP_HASH160 ([a-fA-F0-9]{40}) OP_EQUALVERIFY`

	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		return nil, err
	}

	disassembled, err := txscript.DisasmString(script)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(disassembled)
	if len(match) > 0 {
		return &match[1], nil
	} else {
		return nil, errors.New("no match found")
	}
}

func CalculateWitnessHash(amount int64, script string, transaction string) (*string, error) {
	decodedTx, err := hex.DecodeString(transaction)
	if err != nil {
		return nil, err
	}

	tx, err := btcutil.NewTxFromBytes(decodedTx)
	if err != nil {
		return nil, err
	}

	decodedScript, err := hex.DecodeString(script)
	if err != nil {
		return nil, err
	}

	prevOutFetcher := txscript.NewCannedPrevOutputFetcher(
		decodedScript, amount,
	)

	txhash := txscript.NewTxSigHashes(tx.MsgTx(), prevOutFetcher)
	hash, err := txscript.CalcWitnessSigHash(decodedScript, txhash, txscript.SigHashAll, tx.MsgTx(), 0, amount)
	if err != nil {
		return nil, err
	}

	result := hex.EncodeToString(hash)

	return &result, nil
}

func CalculateWitnessHashPSBT(transaction string) (*string, error) {
	transactionBytes, err := hex.DecodeString(transaction)
	if err != nil {
		return nil, err
	}
	// Reader for the PSBT.
	psbtBytes := []byte(transactionBytes)
	r := bytes.NewReader(psbtBytes)

	// Create instance of a PSBT.
	p, err := psbt.NewFromRawBytes(r, false)
	if err != nil {
		return nil, err
	}
	prevOutFetcher := txscript.NewCannedPrevOutputFetcher(
		p.Inputs[0].WitnessUtxo.PkScript, int64(p.Inputs[0].WitnessUtxo.Value),
	)
	sigHashes := txscript.NewTxSigHashes(p.UnsignedTx, prevOutFetcher)
	hash, err := txscript.CalcWitnessSigHash(p.Inputs[0].WitnessScript, sigHashes, txscript.SigHashAll, p.UnsignedTx, 0, int64(p.Inputs[0].WitnessUtxo.Value))
	if err != nil {
		return nil, err
	}
	result := hex.EncodeToString(hash)
	return &result, nil
}

func DerivationPathFromString(path string) ([]uint32, error) {
	if !strings.HasPrefix(path, "m/") {
		return nil, fmt.Errorf("invalid derivation path: derivation path must start with 'm/'")
	}

	components := strings.Split(path[2:], "/")
	if len(components) == 0 {
		return nil, fmt.Errorf("invalid derivation path: empty component")
	}

	// Validate no empty components
	for _, component := range components {
		if component == "" {
			return nil, fmt.Errorf("invalid derivation path: empty component")
		}
	}
	derivationPath := make([]uint32, 0)

	for _, component := range components {
		isHardened := strings.HasSuffix(component, "'")
		if isHardened {
			component = component[:len(component)-1]
		}

		num, err := strconv.ParseUint(component, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid path: %s", component)
		}

		if isHardened {
			num += 0x80000000
		}
		derivationPath = append(derivationPath, uint32(num))
	}

	return derivationPath, nil
}

func DeriveKey(masterSeed []byte, derivationPath []uint32, networkParams *chaincfg.Params) (*hdkeychain.ExtendedKey, error) {
	masterKey, err := hdkeychain.NewMaster(masterSeed, networkParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create master key: %v", err)
	}
	key := masterKey
	for _, index := range derivationPath {
		key, err = key.Derive(index)
		if err != nil {
			return nil, fmt.Errorf("failed to derive key: %v", err)
		}
	}
	return key, nil
}

func DerivePublicKey(masterSeed []byte, derivationPath string, networkParams *chaincfg.Params) (*secp256k1.PublicKey, error) {
	derivationPathComponents, err := DerivationPathFromString(derivationPath)
	if err != nil {
		return nil, err
	}
	key, err := DeriveKey(masterSeed, derivationPathComponents, networkParams)
	if err != nil {
		return nil, err
	}
	return key.ECPubKey()
}

func ValidateScript(signing *SigningJob, publicKey *secp256k1.PublicKey) (bool, error) {
	generated_script, err := GenerateP2WPKHFromPubkey(publicKey.SerializeCompressed())
	if err != nil {
		return false, err
	}

	// Step 2: Obtain Tx Script from Change Address (Directly from Transaction)
	txHex := *signing.Transaction
	rawTxBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return false, fmt.Errorf("failed to decode transaction hex: %v", err)
	}

	var tx wire.MsgTx
	if err := tx.Deserialize(bytes.NewReader(rawTxBytes)); err != nil {
		return false, fmt.Errorf("failed to deserialize transaction: %v", err)
	}

	if len(tx.TxOut) < 2 {
		// TODO: May need to modify this to validate non-withdrawal L1 transactions.
		return false, fmt.Errorf("no change output found")
	}
	expected_script := tx.TxOut[1].PkScript

	// Step 3: Compare the two scripts
	if !bytes.Equal(generated_script, expected_script) {
		return false, fmt.Errorf("scripts do not match")
	}

	return true, nil
}

func SplitDerivationPath(path string) (hardenedPath []uint32, remainingPath []uint32, err error) {
	derivationPath, err := DerivationPathFromString(path)
	if err != nil {
		return nil, nil, err
	}

	hardenedPath = make([]uint32, 0)
	remainingPath = make([]uint32, 0)

	for _, component := range derivationPath {
		if component >= 0x80000000 {
			hardenedPath = append(hardenedPath, component)
		} else {
			remainingPath = append(remainingPath, component)
		}
	}

	return hardenedPath, remainingPath, nil
}

func GenerateP2WPKHFromPubkey(child_pubkey []byte) ([]byte, error) {
	pkHash := btcutil.Hash160(child_pubkey)
	// Create P2WPKH script: OP_0 <20-byte-key-hash>
	return txscript.NewScriptBuilder().
		AddOp(txscript.OP_0).
		AddData(pkHash).
		Script()
}
