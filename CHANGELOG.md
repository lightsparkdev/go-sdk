# Changelog

# 0.9.0
- Remove the deprecated `payment` field from the `CreateTestModePayment` request.

# 0.8.1
- Add 'NodeId' field to remote signing requests.

# 0.8.0
- Add 'OutgoingPaymentsForInvoice' query to the client.
- Add 'WithdrawalFeeEstimate' query to the client.

# 0.7.4
- Add `DailyLiquidityForecast` to objects.

# 0.7.3
- Add ability to load operation signing key directly.

# 0.7.2
- Fix an issue with 0.7.1's remote signing DER encoding.
- Some minor build fixes to get integration tests working again.

# 0.7.1
- Make remote signing encode OSK signatures in DER format
- Some minor security hardening (CVE-2023-39325, CVE-2022-28948, SSRF protection in demos)

# 0.7.0
- Add a function for cancelling unpaid invoices.
- Add UMA invites support.

# 0.6.1
- Fix serialization of interfaces by including typenames.

# v0.6.0
- Add a more human-readable `balances` field to nodes and wallets.
- Add `deprecated` tags where relevant.

# v0.5.1
- Remove is_raw field in DeriveKeyAndSign request.

# v0.5.0
- Expose remote signing requests and responses for the ability to handle them in custom ways.

# v0.4.1

- Fixed an encoding issue with signing GraphQL requests.

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
