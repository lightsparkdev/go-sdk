package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const OUTGOING_PAYMENTS_FOR_INVOICE_QUERY = `
query OutgoingPaymentsForInvoice(
	$encoded_invoice: String!
	$statuses: [PaymentStatus!]
} {
	outgoing_payments_for_invoice_query(input: {
		encoded_invoice: $encoded_invoice
		statuses: $statuses
	}) {
		...OutgoingPaymentsForInvoiceQueryOutputFragment
	}
}

` + objects.OutgoingPaymentsForInvoiceQueryOutputFragment