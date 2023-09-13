// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

type Signable struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"signable_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"signable_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"signable_updated_at"`
}

const (
	SignableFragment = `
fragment SignableFragment on Signable {
    __typename
    signable_id: id
    signable_created_at: created_at
    signable_updated_at: updated_at
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Signable) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Signable) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Signable) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
