package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const INVOICE_FOR_PAYMENT_HASH_QUERY = `
query InvoiceForPaymentHash($payment_hash: Hash32!) {
	invoice_for_payment_hash(input: {
		payment_hash: $payment_hash
	}) {
		invoice {
			...InvoiceFragment
		}
	}
}

` + objects.InvoiceFragment
