package scripts

const FAIL_HTLCS_MUTATION = `
mutation FailHtlcs($invoice_id: ID!) {
	fail_htlcs(input: { invoice_id: $invoice_id })
}`
