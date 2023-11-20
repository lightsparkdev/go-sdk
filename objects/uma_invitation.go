// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// UmaInvitation This is an object representing an UMA.ME invitation.
type UmaInvitation struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"uma_invitation_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"uma_invitation_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"uma_invitation_updated_at"`

	// Code The code that uniquely identifies this invitation.
	Code string `json:"uma_invitation_code"`

	// Url The URL where this invitation can be claimed.
	Url string `json:"uma_invitation_url"`

	// InviterUma The UMA of the user who created the invitation.
	InviterUma string `json:"uma_invitation_inviter_uma"`

	// InviteeUma The UMA of the user who claimed the invitation.
	InviteeUma *string `json:"uma_invitation_invitee_uma"`

	// IncentivesStatus The current status of the incentives that may be tied to this invitation.
	IncentivesStatus IncentivesStatus `json:"uma_invitation_incentives_status"`

	// IncentivesIneligibilityReason The reason why the invitation is not eligible for incentives, if applicable.
	IncentivesIneligibilityReason *IncentivesIneligibilityReason `json:"uma_invitation_incentives_ineligibility_reason"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	UmaInvitationFragment = `
fragment UmaInvitationFragment on UmaInvitation {
    __typename
    uma_invitation_id: id
    uma_invitation_created_at: created_at
    uma_invitation_updated_at: updated_at
    uma_invitation_code: code
    uma_invitation_url: url
    uma_invitation_inviter_uma: inviter_uma
    uma_invitation_invitee_uma: invitee_uma
    uma_invitation_incentives_status: incentives_status
    uma_invitation_incentives_ineligibility_reason: incentives_ineligibility_reason
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj UmaInvitation) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj UmaInvitation) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj UmaInvitation) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj UmaInvitation) GetTypename() string {
	return obj.Typename
}
