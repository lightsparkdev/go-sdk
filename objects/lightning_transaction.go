// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// LightningTransaction This is an object representing a transaction made over the Lightning Network. You can retrieve this object to receive information about a specific transaction made over Lightning for a Lightspark node.
type LightningTransaction interface {
	Transaction
	Entity
}

func LightningTransactionUnmarshal(data map[string]interface{}) (LightningTransaction, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "IncomingPayment":
		var incomingPayment IncomingPayment
		if err := json.Unmarshal(dataJSON, &incomingPayment); err != nil {
			return nil, err
		}
		return incomingPayment, nil
	case "OutgoingPayment":
		var outgoingPayment OutgoingPayment
		if err := json.Unmarshal(dataJSON, &outgoingPayment); err != nil {
			return nil, err
		}
		return outgoingPayment, nil
	case "RoutingTransaction":
		var routingTransaction RoutingTransaction
		if err := json.Unmarshal(dataJSON, &routingTransaction); err != nil {
			return nil, err
		}
		return routingTransaction, nil

	default:
		return nil, fmt.Errorf("unknown LightningTransaction type: %s", data["__typename"])
	}
}
