// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type CurrencyUnit int

const (
	CurrencyUnitUndefined CurrencyUnit = iota

	// Bitcoin is the cryptocurrency native to the Bitcoin network. It is used as the native medium for value transfer for the Lightning Network.
	CurrencyUnitBitcoin
	// 0.00000001 (10e-8) Bitcoin or one hundred millionth of a Bitcoin. This is the unit most commonly used in Lightning transactions.
	CurrencyUnitSatoshi
	// 0.001 Satoshi, or 10e-11 Bitcoin. We recommend using the Satoshi unit instead when possible.
	CurrencyUnitMillisatoshi
	// United States Dollar.
	CurrencyUnitUsd
	// 0.000000001 (10e-9) Bitcoin or a billionth of a Bitcoin. We recommend using the Satoshi unit instead when possible.
	CurrencyUnitNanobitcoin
	// 0.000001 (10e-6) Bitcoin or a millionth of a Bitcoin. We recommend using the Satoshi unit instead when possible.
	CurrencyUnitMicrobitcoin
	// 0.001 (10e-3) Bitcoin or a thousandth of a Bitcoin. We recommend using the Satoshi unit instead when possible.
	CurrencyUnitMillibitcoin
)

func (a *CurrencyUnit) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = CurrencyUnitUndefined
	case "BITCOIN":
		*a = CurrencyUnitBitcoin
	case "SATOSHI":
		*a = CurrencyUnitSatoshi
	case "MILLISATOSHI":
		*a = CurrencyUnitMillisatoshi
	case "USD":
		*a = CurrencyUnitUsd
	case "NANOBITCOIN":
		*a = CurrencyUnitNanobitcoin
	case "MICROBITCOIN":
		*a = CurrencyUnitMicrobitcoin
	case "MILLIBITCOIN":
		*a = CurrencyUnitMillibitcoin

	}
	return nil
}

func (a CurrencyUnit) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case CurrencyUnitBitcoin:
		s = "BITCOIN"
	case CurrencyUnitSatoshi:
		s = "SATOSHI"
	case CurrencyUnitMillisatoshi:
		s = "MILLISATOSHI"
	case CurrencyUnitUsd:
		s = "USD"
	case CurrencyUnitNanobitcoin:
		s = "NANOBITCOIN"
	case CurrencyUnitMicrobitcoin:
		s = "MICROBITCOIN"
	case CurrencyUnitMillibitcoin:
		s = "MILLIBITCOIN"

	}
	return s
}

func (a CurrencyUnit) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
