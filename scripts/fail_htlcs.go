package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const FAIL_HTLCS_MUTATION = `
mutation FailHtlcs($invoice_id: ID!, $cancel_invoice: Boolean!) {
	fail_htlcs(input: { invoice_id: $invoice_id, cancel_invoice: $cancel_invoice}) {
		...FailHtlcsOutputFragment
	}
}

` + objects.FailHtlcsOutputFragment
