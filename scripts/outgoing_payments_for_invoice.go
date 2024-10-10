package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const OUTGOING_PAYMENTS_FOR_INVOICE_QUERY = `
query OutgoingPaymentsForInvoice(
	$encoded_invoice: String!
	$statuses: [TransactionStatus!]
) {
	outgoing_payments_for_invoice(input: {
		encoded_invoice: $encoded_invoice
		statuses: $statuses
	}) {
		...OutgoingPaymentsForInvoiceQueryOutputFragment
	}
}

` + objects.OutgoingPaymentsForInvoiceQueryOutputFragment
