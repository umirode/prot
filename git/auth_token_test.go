package git

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"testing"
)

func TestAuthToken(t *testing.T) {
	authToken := NewAuthToken(map[string]string{
		"Token": "token",
	})
	assert.NotNil(t, authToken)

	authMethod, err := authToken.Create()
	assert.NoError(t, err)
	assert.NotNil(t, authMethod)
	assert.Implements(t, (*transport.AuthMethod)(nil), authMethod)
	assert.Equal(t, "token", authMethod.(*http.TokenAuth).Token)
}

func TestAuthTokenError(t *testing.T) {
	authToken := NewAuthToken(map[string]string{})
	assert.NotNil(t, authToken)

	authMethod, err := authToken.Create()
	assert.Error(t, err)
	assert.Nil(t, authMethod)
}
