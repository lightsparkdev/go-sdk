module github.com/lightsparkdev/go-sdk

go 1.21

toolchain go1.21.0

require (
	github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go v0.2.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.17.0
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/uma-universal-money-address/uma-go-sdk v0.6.3-0.20240326175044-056b9d482661 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

require (
	github.com/btcsuite/btcd v0.23.4
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Break dependency cycle with objx.
// See https://github.com/stretchr/objx/pull/140
exclude github.com/stretchr/testify v1.8.2
