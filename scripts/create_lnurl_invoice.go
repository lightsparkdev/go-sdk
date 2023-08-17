// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CREATE_LNURL_INVOICE_MUTATION = `
mutation CreateLnurlInvoice(
    $node_id: ID!
    $amount_msats: Long!
	$metadata_hash: String!
	$payment_hash: String
    $preimage_nonce: String
	$expiry_secs: Int
) {
    create_lnurl_invoice(input: {
        node_id: $node_id
        amount_msats: $amount_msats
        metadata_hash: $metadata_hash
		payment_hash: $payment_hash
		preimage_nonce: $preimage_nonce
		expiry_secs: $expiry_secs
    }) {
        invoice {
            ...InvoiceFragment
        }
    }
}

` + objects.InvoiceFragment
