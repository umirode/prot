package git

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"testing"
)

func TestAuthBasicAuth(t *testing.T) {
	authBasicAuth := NewAuthBasicAuth(map[string]string{
		"Username": "user",
		"Password": "pass",
	})
	assert.NotNil(t, authBasicAuth)

	authMethod, err := authBasicAuth.Create()
	assert.NoError(t, err)
	assert.NotNil(t, authMethod)
	assert.Equal(t, &http.BasicAuth{
		Username: "user",
		Password: "pass",
	}, authMethod)
}

func TestAuthBasicAuthErrorUsername(t *testing.T) {
	authBasicAuth := NewAuthBasicAuth(map[string]string{
		"Username": "user",
	})
	assert.NotNil(t, authBasicAuth)

	authMethod, err := authBasicAuth.Create()
	assert.Error(t, err)
	assert.Nil(t, authMethod)
}

func TestAuthBasicAuthErrorPassword(t *testing.T) {
	authBasicAuth := NewAuthBasicAuth(map[string]string{
		"Password": "pass",
	})
	assert.NotNil(t, authBasicAuth)

	authMethod, err := authBasicAuth.Create()
	assert.Error(t, err)
	assert.Nil(t, authMethod)
}
