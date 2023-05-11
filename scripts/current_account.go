package scripts

import "lightspark/objects"


const CURRENT_ACCOUNT_QUERY = `
query GetCurrentAccount {
    current_account {
        ...AccountFragment
    }
}

` + objects.AccountFragment