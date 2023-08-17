// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts


const SCREEN_BITCOIN_ADDRESSES_MUTATION = `
mutation ScreenBitcoinAddresses(
    $provider: CryptoSanctionsScreeningProvider!
    $addresses: [String!]!
) {
    screen_bitcoin_addresses(input: {
        provider: $provider
        addresses: $addresses
    }) {
        ratings
    }
}

`
