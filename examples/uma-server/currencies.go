package main

import "github.com/uma-universal-money-address/uma-go-sdk/uma"

var SatsCurrency = uma.Currency{
	Code:                "SAT",
	Name:                "Satoshis",
	Symbol:              "SAT",
	MillisatoshiPerUnit: 1000,
	Convertible: uma.ConvertibleCurrency{
		MinSendable: 1,
		MaxSendable: 100_000_000,
	},
	Decimals: 0,
}

var UsdCurrency = uma.Currency{
	Code:                "USD",
	Name:                "US Dollars",
	Symbol:              "$",
	MillisatoshiPerUnit: MillisatoshiPerUsd,
	Convertible: uma.ConvertibleCurrency{
		MinSendable: 1,
		MaxSendable: 1_000,
	},
	Decimals: 2,
}
