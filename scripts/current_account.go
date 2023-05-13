// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "lightspark/objects"

const CURRENT_ACCOUNT_QUERY = `
query GetCurrentAccount {
    current_account {
        ...AccountFragment
    }
}

` + objects.AccountFragment
