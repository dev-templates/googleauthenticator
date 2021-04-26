package googleauthenticator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticator(t *testing.T) {
	formattedKey := GenerateKey()
	authenticator := NewAuthenticator("issuer", "xxx@gmail.com", formattedKey)
	passcode := authenticator.GenerateToken()
	assert.True(t, authenticator.VerifyToken(passcode))
}
