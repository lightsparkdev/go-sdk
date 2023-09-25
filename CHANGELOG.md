# Changelog

# v0.4.0

- Moved the UMA protocol out to its own Lightspark-agnostic
  repo: https://github.com/uma-universal-money-address/umd-go-sdk. Fixed some bugs along the way.

# v0.3.0

- Breaking change! Adjusting the API to support both remote signing and OSK side-by-side. See `SigningKeyLoader`
  and `client.LoadNodeSigningKey` for more details.
- Added full support for the UMA protocol. See the `uma` package and `examples/uma-server`.

# v0.2.0

Breaking change! Migrating to remote signing.

# v0.1.5

Added support for generating invoices for LNURLs.

# v0.1.4

Use CSPRNG to generate nonce.

# v0.1.3

Add two functions for test mode.

- CreateTestModeInvoice for creating an test invoice.
- CreateTestModePayment for sending a test payment to an invoice.

# v0.1.2

Add webhook.

# v0.1.1

Fixed payment related bugs.

- Fixed time format for sending the signed requests
- Fixed random number seed for sending the signed requests

# v0.1.0

First draft of the SDK.
