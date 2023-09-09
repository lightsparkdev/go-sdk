// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

const NODE_CHANNEL_UTXO_QUERY = `
query getNodeChannelUtxos($node_id: ID!) {
	entity(id: $node_id) { 
		...on LightsparkNode {
			channels {
				entities {
					funding_transaction {
						transaction_hash
					}
				}
			}
		}
	}
}`
