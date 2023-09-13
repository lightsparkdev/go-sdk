package utils

import (
	"errors"
	"github.com/lightsparkdev/go-sdk/objects"
)

func ValueMilliSatoshi(amount objects.CurrencyAmount) (int64, error) {
	if amount.OriginalUnit == objects.CurrencyUnitMillisatoshi {
		return amount.OriginalValue, nil
	} else if amount.OriginalUnit == objects.CurrencyUnitSatoshi {
		return amount.OriginalValue * 1000, nil
	} else if amount.OriginalUnit == objects.CurrencyUnitBitcoin {
		return amount.OriginalValue * 100_000_000_000, nil
	} else if amount.OriginalUnit == objects.CurrencyUnitMicrobitcoin {
		return amount.OriginalValue * 100_000, nil
	} else if amount.OriginalUnit == objects.CurrencyUnitMillibitcoin {
		return amount.OriginalValue * 100_000_000, nil
	} else if amount.OriginalUnit == objects.CurrencyUnitNanobitcoin {
		return amount.OriginalValue * 100, nil
	}
	return -1, errors.New("invalid currency conversion")
}
