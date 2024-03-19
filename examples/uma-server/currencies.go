package main

import (
	umaprotocol "github.com/uma-universal-money-address/uma-go-sdk/uma/protocol"
)

var SatsCurrency = umaprotocol.Currency{
	Code:                "SAT",
	Name:                "Satoshis",
	Symbol:              "SAT",
	MillisatoshiPerUnit: 1000,
	Convertible: umaprotocol.ConvertibleCurrency{
		MinSendable: 1,
		MaxSendable: 100_000_000,
	},
	Decimals: 0,
}

var UsdCurrency = umaprotocol.Currency{
	Code:                "USD",
	Name:                "US Dollars",
	Symbol:              "$",
	MillisatoshiPerUnit: MillisatoshiPerUsd,
	Convertible: umaprotocol.ConvertibleCurrency{
		MinSendable: 1,
		MaxSendable: 1_000,
	},
	Decimals: 2,
}
