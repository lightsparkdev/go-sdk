package umautils

import (
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/services"
	"github.com/stretchr/testify/require"
)

func TestHashUmaIdentifier(t *testing.T) {
	privKeyBytes := []byte("xyz")

	mockTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	hashedUma := services.HashUmaIdentifier("user@domain.com", privKeyBytes, mockTime)
	hashedUmaSameMonth := services.HashUmaIdentifier("user@domain.com", privKeyBytes, mockTime)
	require.Equal(t, hashedUma, hashedUmaSameMonth)
	t.Log(hashedUma)

	mockTime = time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)
	hashedUmaNewMonth := services.HashUmaIdentifier("user@domain.com", privKeyBytes, mockTime)
	require.NotEqual(t, hashedUma, hashedUmaNewMonth)
	t.Log(hashedUmaNewMonth)
}
