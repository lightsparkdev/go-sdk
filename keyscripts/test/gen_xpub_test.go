// Copyright ©, 2024-present, Lightspark Group, Inc. - All Rights Reserved
package keyscripts_test

import (
	"testing"

	keyscripts "github.com/lightsparkdev/go-sdk/keyscripts"
	"github.com/stretchr/testify/require"
)

func TestGenerateXpub(t *testing.T) {
	tests := []struct {
		name           string
		seed           string
		derivationPath []uint32
		expectedXpub   string
		bitcoinNetwork string
	}{
		// Test vectors from https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#user-content-Test_Vectors

		{
			name:           "Test Vector 1 - m/0'",
			seed:           "000102030405060708090a0b0c0d0e0f",
			derivationPath: []uint32{0x80000000},
			expectedXpub:   "xpub68Gmy5EdvgibQVfPdqkBBCHxA5htiqg55crXYuXoQRKfDBFA1WEjWgP6LHhwBZeNK1VTsfTFUHCdrfp1bgwQ9xv5ski8PX9rL2dZXvgGDnw",
			bitcoinNetwork: "mainnet",
		},
		{
			name:           "Test Vector 1 - m/0'/1",
			seed:           "000102030405060708090a0b0c0d0e0f",
			derivationPath: []uint32{0x80000000, 1},
			expectedXpub:   "xpub6ASuArnXKPbfEwhqN6e3mwBcDTgzisQN1wXN9BJcM47sSikHjJf3UFHKkNAWbWMiGj7Wf5uMash7SyYq527Hqck2AxYysAA7xmALppuCkwQ",
			bitcoinNetwork: "mainnet",
		},
		{
			name:           "Test Vector 1 - m/0'/1/2'",
			seed:           "000102030405060708090a0b0c0d0e0f",
			derivationPath: []uint32{0x80000000, 1, 0x80000002},
			expectedXpub:   "xpub6D4BDPcP2GT577Vvch3R8wDkScZWzQzMMUm3PWbmWvVJrZwQY4VUNgqFJPMM3No2dFDFGTsxxpG5uJh7n7epu4trkrX7x7DogT5Uv6fcLW5",
			bitcoinNetwork: "mainnet",
		},
		// Test Vector 2
		{
			name:           "Test Vector 2 - m/0",
			seed:           "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			derivationPath: []uint32{0},
			expectedXpub:   "xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH",
			bitcoinNetwork: "mainnet",
		},
		{
			name:           "Test Vector 2 - m/0/2147483647'",
			seed:           "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			derivationPath: []uint32{0, 0x80000000 + 2147483647},
			expectedXpub:   "xpub6ASAVgeehLbnwdqV6UKMHVzgqAG8Gr6riv3Fxxpj8ksbH9ebxaEyBLZ85ySDhKiLDBrQSARLq1uNRts8RuJiHjaDMBU4Zn9h8LZNnBC5y4a",
			bitcoinNetwork: "mainnet",
		},
		{
			name:           "Test Vector 2 - m/0/2147483647'/1",
			seed:           "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			derivationPath: []uint32{0, 0x80000000 + 2147483647, 1},
			expectedXpub:   "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon",
			bitcoinNetwork: "mainnet",
		},
	}

	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			xpub, err := keyscripts.GenHardenedXPub(tv.seed, tv.derivationPath, tv.bitcoinNetwork)
			require.NoError(t, err)
			require.Equal(t, tv.expectedXpub, xpub)
		})
	}
}
