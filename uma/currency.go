package uma

type Currency struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	MillisatoshiPerUnit int64  `json:"multiplier"`
	MinSendable         int64  `json:"minSendable"`
	MaxSendable         int64  `json:"maxSendable"`
}
