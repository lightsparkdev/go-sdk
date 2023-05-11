// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type RichText struct {
	Text string `json:"rich_text_text"`
}

const (
	RichTextFragment = `
fragment RichTextFragment on RichText {
    __typename
    rich_text_text: text
}
`
)
