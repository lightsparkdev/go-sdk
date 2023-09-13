package uma

import "encoding/json"

type KycStatus int

const (
	KycStatusUnknown KycStatus = iota
	KycStatusNotVerified
	KycStatusPending
	KycStatusVerified
)

func (k *KycStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*k = KycStatusUnknown
	case "UNKNOWN":
		*k = KycStatusUnknown
	case "NOT_VERIFIED":
		*k = KycStatusNotVerified
	case "PENDING":
		*k = KycStatusPending
	case "VERIFIED":
		*k = KycStatusVerified
	}
	return nil
}

func (k KycStatus) StringValue() string {
	var s string
	switch k {
	default:
		s = "undefined"
	case KycStatusUnknown:
		s = "UNKNOWN"
	case KycStatusNotVerified:
		s = "NOT_VERIFIED"
	case KycStatusPending:
		s = "PENDING"
	case KycStatusVerified:
		s = "VERIFIED"
	}
	return s
}

func (k KycStatus) MarshalJSON() ([]byte, error) {
	s := k.StringValue()
	return json.Marshal(s)
}
