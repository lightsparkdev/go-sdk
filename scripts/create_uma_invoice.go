// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CREATE_UMA_INVOICE_MUTATION = `
mutation CreateUmaInvoice(
    $node_id: ID!
    $amount_msats: Long!
	$metadata_hash: String!
	$expiry_secs: Int
    $receiver_hash: String = null
) {
    create_uma_invoice(input: {
        node_id: $node_id
        amount_msats: $amount_msats
        metadata_hash: $metadata_hash
		expiry_secs: $expiry_secs
        receiver_hash: $receiver_hash
    }) {
        invoice {
            ...InvoiceFragment
        }
    }
}

` + objects.InvoiceFragment
