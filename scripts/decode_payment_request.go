// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "lightspark/objects"

const DECODE_PAYMENT_REQUEST_QUERY = `
query DecodedPaymentRequest(
    $encoded_payment_request: String!
) {
    decoded_payment_request(encoded_payment_request: $encoded_payment_request) {
        __typename
        ... on InvoiceData {
            ...InvoiceDataFragment
        }
    }
}

` + objects.InvoiceDataFragment
