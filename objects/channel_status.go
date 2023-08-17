// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// This is an enum representing the status of a channel on the Lightning Network.
type ChannelStatus int

const (
	ChannelStatusUndefined ChannelStatus = iota

	// The channel is online and ready to send and receive funds.
	ChannelStatusOk
	// The channel has been created, but the Bitcoin transaction that initiates it still needs to be confirmed on the Bitcoin blockchain.
	ChannelStatusPending
	// The channel is not available, likely because the peer is not online.
	ChannelStatusOffline
	// The channel is behaving properly, but its remote balance is much higher than its local balance so it is not balanced properly for sending funds out.
	ChannelStatusUnbalancedForSend
	// The channel is behaving properly, but its remote balance is much lower than its local balance so it is not balanced properly for receiving funds.
	ChannelStatusUnbalancedForReceive
	// The channel has been closed. Information about the channel is still available for historical purposes but the channel cannot be used anymore.
	ChannelStatusClosed
	// Something unexpected happened and we cannot determine the status of this channel. Please try again later or contact the support.
	ChannelStatusError
)

func (a *ChannelStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = ChannelStatusUndefined
	case "OK":
		*a = ChannelStatusOk
	case "PENDING":
		*a = ChannelStatusPending
	case "OFFLINE":
		*a = ChannelStatusOffline
	case "UNBALANCED_FOR_SEND":
		*a = ChannelStatusUnbalancedForSend
	case "UNBALANCED_FOR_RECEIVE":
		*a = ChannelStatusUnbalancedForReceive
	case "CLOSED":
		*a = ChannelStatusClosed
	case "ERROR":
		*a = ChannelStatusError

	}
	return nil
}

func (a ChannelStatus) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case ChannelStatusOk:
		s = "OK"
	case ChannelStatusPending:
		s = "PENDING"
	case ChannelStatusOffline:
		s = "OFFLINE"
	case ChannelStatusUnbalancedForSend:
		s = "UNBALANCED_FOR_SEND"
	case ChannelStatusUnbalancedForReceive:
		s = "UNBALANCED_FOR_RECEIVE"
	case ChannelStatusClosed:
		s = "CLOSED"
	case ChannelStatusError:
		s = "ERROR"

	}
	return s
}

func (a ChannelStatus) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
