package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const INCOMING_PAYMENTS_FOR_INVOICE_QUERY = `
query IncomingPaymentsForInvoice(
	$invoice_id: ID!
	$statuses: [TransactionStatus!]
) {
	incoming_payments_for_invoice(input: {
		invoice_id: $invoice_id
		statuses: $statuses
	}) {
		...IncomingPaymentsForInvoiceQueryOutputFragment
	}
}

` + objects.IncomingPaymentsForInvoiceQueryOutputFragment
