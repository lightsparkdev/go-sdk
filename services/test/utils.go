package services_test

import (
	"encoding/hex"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateInvoiceForNode(client *services.LightsparkClient, nodeID string) (*objects.Invoice, error) {
	invoice, err := client.CreateInvoice(nodeID, 10_000_000, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func LoadSeedAsSigningKey(
	t *testing.T, nodeId string, seedHex string, network objects.BitcoinNetwork, client *services.LightsparkClient) {
	seedBytes, err := hex.DecodeString(seedHex)
	require.NoError(t, err)
	client.LoadNodeSigningKey(
		nodeId,
		*services.NewSigningKeyLoaderFromSignerMasterSeed(seedBytes, network))
}
