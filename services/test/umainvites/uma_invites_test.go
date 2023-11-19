//go:build integration
// +build integration

package umainvites

import (
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/services"
	servicestest "github.com/lightsparkdev/go-sdk/services/test"
	"github.com/stretchr/testify/require"
	"testing"
)

func NewUmameClient() *services.LightsparkClient {
	config := servicestest.NewConfig()
	clientRequester := requester.Requester{
		ApiTokenClientId:     config.ApiClientSecret,
		ApiTokenClientSecret: config.ApiClientID,
		BaseUrl:              &config.ApiClientEndpoint,
	}
	return &services.LightsparkClient{Requester: &clientRequester}
}

func TestInvalidPhoneNumber(t *testing.T) {
	client := NewUmameClient()
	_, err := client.CreateUmaInvitationWithIncentives("bob@vasp.com", "219218", objects.RegionCodeUs)
	if err == nil {
		t.Errorf("Expected error when creating invitation with invalid phone number")
	}
}

func TestCreateInvitation(t *testing.T) {
	client := NewUmameClient()
	invitation, err := client.CreateUmaInvitation("bob@vasp.com")
	if err != nil {
		t.Fatalf("Error creating invitation: %v", err)
	}
	t.Logf("Created invitation %v", invitation)
}

func TestCreateInvitationWithIncentives(t *testing.T) {
	client := NewUmameClient()
	invitation, err := client.CreateUmaInvitationWithIncentives("bob@vasp.com", "+15555555555", objects.RegionCodeUs)
	if err != nil {
		t.Fatalf("Error creating invitation: %v", err)
	}
	t.Logf("Created invitation %v", invitation)
}

func TestClaimInvitation(t *testing.T) {
	client := NewUmameClient()
	invitation, err := client.CreateUmaInvitation("bob@vasp.com")
	if err != nil {
		t.Fatalf("Error creating invitation: %v", err)
	}
	t.Logf("Created invitation %v", invitation)
	claimedInvitation, err := client.ClaimUmaInvitation(invitation.Code, "alice@vasp2.com")
	if err != nil {
		t.Fatalf("Error claiming invitation: %v", err)
	}
	t.Logf("Claimed invitation %v", claimedInvitation)
}

func TestClaimInvitationWithIncentives(t *testing.T) {
	client := NewUmameClient()
	invitation, err := client.CreateUmaInvitationWithIncentives("bob@vasp.com", "+15555555555", objects.RegionCodeUs)
	if err != nil {
		t.Fatalf("Error creating invitation: %v", err)
	}
	t.Logf("Created invitation %v", invitation)
	claimedInvitation, err := client.ClaimUmaInvitationWithIncentives(invitation.Code, "alice@vasp2.com", "+15555555556", objects.RegionCodeUs)
	if err != nil {
		t.Fatalf("Error claiming invitation: %v", err)
	}
	t.Logf("Claimed invitation %v", claimedInvitation)
}

func TestFetchInvitation(t *testing.T) {
	client := NewUmameClient()
	invitation, err := client.CreateUmaInvitation("bob@vasp.com")
	if err != nil {
		t.Fatalf("Error creating invitation: %v", err)
	}
	t.Logf("Created invitation %v", invitation)
	fetchedInvitation, err := client.FetchUmaInvitation(invitation.Code)
	if err != nil {
		t.Fatalf("Error fetching invitation: %v", err)
	}
	require.Equal(t, invitation.Code, fetchedInvitation.Code)
}
