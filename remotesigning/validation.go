package remotesigning

import (
	"bytes"
	"encoding/hex"
	"errors"
	"regexp"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/txscript"
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
		return nil, errors.New("No match found")
	}
}

func CalculateWitnessHash(amount int64, script string, transaction string) (*string, error) {
	decodedTx, err := hex.DecodeString(transaction)
	if err != nil {
		return nil, err
	}

	tx, err := btcutil.NewTxFromBytes(decodedTx)
	if err != nil {
		return nil ,err
	}

	decodedScript, err := hex.DecodeString(script)
	if err != nil {
		return nil ,err
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