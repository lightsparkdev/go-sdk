// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// Entity This interface is used by all the entities in the Lightspark system. It defines a few core fields that are available everywhere. Any object that implements this interface can be queried using the `entity` query and its ID.
type Entity interface {

	// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	GetId() string

	// GetCreatedAt The date and time when the entity was first created.
	GetCreatedAt() time.Time

	// GetUpdatedAt The date and time when the entity was last updated.
	GetUpdatedAt() time.Time

	// GetTypename The typename of the object
	GetTypename() string
}

func EntityUnmarshal(data map[string]interface{}) (Entity, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "Account":
		var account Account
		if err := json.Unmarshal(dataJSON, &account); err != nil {
			return nil, err
		}
		return account, nil
	case "ApiToken":
		var apiToken ApiToken
		if err := json.Unmarshal(dataJSON, &apiToken); err != nil {
			return nil, err
		}
		return apiToken, nil
	case "Channel":
		var channel Channel
		if err := json.Unmarshal(dataJSON, &channel); err != nil {
			return nil, err
		}
		return channel, nil
	case "ChannelClosingTransaction":
		var channelClosingTransaction ChannelClosingTransaction
		if err := json.Unmarshal(dataJSON, &channelClosingTransaction); err != nil {
			return nil, err
		}
		return channelClosingTransaction, nil
	case "ChannelOpeningTransaction":
		var channelOpeningTransaction ChannelOpeningTransaction
		if err := json.Unmarshal(dataJSON, &channelOpeningTransaction); err != nil {
			return nil, err
		}
		return channelOpeningTransaction, nil
	case "ChannelSnapshot":
		var channelSnapshot ChannelSnapshot
		if err := json.Unmarshal(dataJSON, &channelSnapshot); err != nil {
			return nil, err
		}
		return channelSnapshot, nil
	case "Deposit":
		var deposit Deposit
		if err := json.Unmarshal(dataJSON, &deposit); err != nil {
			return nil, err
		}
		return deposit, nil
	case "GraphNode":
		var graphNode GraphNode
		if err := json.Unmarshal(dataJSON, &graphNode); err != nil {
			return nil, err
		}
		return graphNode, nil
	case "Hop":
		var hop Hop
		if err := json.Unmarshal(dataJSON, &hop); err != nil {
			return nil, err
		}
		return hop, nil
	case "IncomingPayment":
		var incomingPayment IncomingPayment
		if err := json.Unmarshal(dataJSON, &incomingPayment); err != nil {
			return nil, err
		}
		return incomingPayment, nil
	case "IncomingPaymentAttempt":
		var incomingPaymentAttempt IncomingPaymentAttempt
		if err := json.Unmarshal(dataJSON, &incomingPaymentAttempt); err != nil {
			return nil, err
		}
		return incomingPaymentAttempt, nil
	case "Invoice":
		var invoice Invoice
		if err := json.Unmarshal(dataJSON, &invoice); err != nil {
			return nil, err
		}
		return invoice, nil
	case "LightsparkNodeWithOSK":
		var lightsparkNodeWithOSK LightsparkNodeWithOSK
		if err := json.Unmarshal(dataJSON, &lightsparkNodeWithOSK); err != nil {
			return nil, err
		}
		return lightsparkNodeWithOSK, nil
	case "LightsparkNodeWithRemoteSigning":
		var lightsparkNodeWithRemoteSigning LightsparkNodeWithRemoteSigning
		if err := json.Unmarshal(dataJSON, &lightsparkNodeWithRemoteSigning); err != nil {
			return nil, err
		}
		return lightsparkNodeWithRemoteSigning, nil
	case "Offer":
		var offer Offer
		if err := json.Unmarshal(dataJSON, &offer); err != nil {
			return nil, err
		}
		return offer, nil
	case "OfferData":
		var offerData OfferData
		if err := json.Unmarshal(dataJSON, &offerData); err != nil {
			return nil, err
		}
		return offerData, nil
	case "OutgoingPayment":
		var outgoingPayment OutgoingPayment
		if err := json.Unmarshal(dataJSON, &outgoingPayment); err != nil {
			return nil, err
		}
		return outgoingPayment, nil
	case "OutgoingPaymentAttempt":
		var outgoingPaymentAttempt OutgoingPaymentAttempt
		if err := json.Unmarshal(dataJSON, &outgoingPaymentAttempt); err != nil {
			return nil, err
		}
		return outgoingPaymentAttempt, nil
	case "RoutingTransaction":
		var routingTransaction RoutingTransaction
		if err := json.Unmarshal(dataJSON, &routingTransaction); err != nil {
			return nil, err
		}
		return routingTransaction, nil
	case "Signable":
		var signable Signable
		if err := json.Unmarshal(dataJSON, &signable); err != nil {
			return nil, err
		}
		return signable, nil
	case "SignablePayload":
		var signablePayload SignablePayload
		if err := json.Unmarshal(dataJSON, &signablePayload); err != nil {
			return nil, err
		}
		return signablePayload, nil
	case "UmaInvitation":
		var umaInvitation UmaInvitation
		if err := json.Unmarshal(dataJSON, &umaInvitation); err != nil {
			return nil, err
		}
		return umaInvitation, nil
	case "Wallet":
		var wallet Wallet
		if err := json.Unmarshal(dataJSON, &wallet); err != nil {
			return nil, err
		}
		return wallet, nil
	case "Withdrawal":
		var withdrawal Withdrawal
		if err := json.Unmarshal(dataJSON, &withdrawal); err != nil {
			return nil, err
		}
		return withdrawal, nil
	case "WithdrawalRequest":
		var withdrawalRequest WithdrawalRequest
		if err := json.Unmarshal(dataJSON, &withdrawalRequest); err != nil {
			return nil, err
		}
		return withdrawalRequest, nil

	default:
		return nil, fmt.Errorf("unknown Entity type: %s", data["__typename"])
	}
}
