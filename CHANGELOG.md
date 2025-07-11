# Changelog

# v0.16.2
- Improve webhook request validation and test coverage

# v0.16.1
- Replace go-bip32 library for key derivation

# v0.16.0
- Add `LightsparkClient.CreateInvoiceWithPaymentHash`
- Add util functions for preimage generation
- Add L1 wallet validation
- Add support for offers

# v0.15.1
- Fix: Remove default null param on `PayUmaInvoice` graphql

# v0.15.0
- Add `is_uma` and `is_lnurl` to the Invoice object and `RELEASE_PAYMENT_PREIMAGE` webhook data.
- Add extra params to `CreateUmaInvoice` and `PayUmaInvoice` for analytics

# v0.14.0
- Add a function to specifiy fess for request withdrawal.

# v0.13.0
- Add idempotency where useful
  - Add optional PayInvoiceWithIdempotencyKey and RequestWithdrawalWithIdempotencyKey functions to the client.
  - Add IdempotencyKey field to OutgoingPayment and Withdrawal objects.
  - Add new FetchOutgoingPaymentsByIdempotencyKey query.
- Add FailHtlcs function to cancel pending htlcs (for example for HODL invoices).
- Add FetchInvoiceByPaymentHash and FetchOutgoingPaymentsByPaymentHash to get all outgoing payments for a specific hash.
- Update objects from graphql schema.

# v0.12.0
- Add `RequestError`, `GraphQLInternalError` and, `GraphQLError` to the client to better differentiate between different types of errors.

# v0.11.1
- Update "lightspark-crypto-go" to v0.2.0 to fix build on some linux versions.

# v0.11.0
- Add `CreateNodeWalletAddressWithKeys` to the client to return the public keys for the L1 address.
- Add a utility function to create a 2-of-2 multisig L1 address based on 2 public keys.

# 0.10.0
- Add `FetchIncomingPaymentsByInvoice` to the client.
- Return the whole payment object when fetching payments for an invoice.
- Some tweaks to the `ChannelSnapshot` object.
- Adding more documentation.

# 0.9.1
- Fix a bug for fetching outgoing payments for an invoice.

# 0.9.0
- Remove the deprecated `payment` field from the `CreateTestModePayment` request.
- Remove the unused fragment

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
