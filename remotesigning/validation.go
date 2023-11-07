package remotesigning

import (
	"encoding/hex"
	"errors"
	"regexp"

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
