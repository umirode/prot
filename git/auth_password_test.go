package git

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"testing"
)

func TestAuthPassword(t *testing.T) {
	authPassword := NewAuthPassword(map[string]string{
		"User":     "user",
		"Password": "pass",
	})
	assert.NotNil(t, authPassword)

	authMethod, err := authPassword.Create()
	assert.NoError(t, err)
	assert.NotNil(t, authMethod)
	assert.Equal(t, &ssh.Password{
		User:     "user",
		Password: "pass",
	}, authMethod)
}

func TestAuthPasswordUser(t *testing.T) {
	authPassword := NewAuthPassword(map[string]string{
		"User": "user",
	})
	assert.NotNil(t, authPassword)

	authMethod, err := authPassword.Create()
	assert.Error(t, err)
	assert.Nil(t, authMethod)
}

func TestAuthPasswordPassword(t *testing.T) {
	authPassword := NewAuthPassword(map[string]string{
		"Password": "pass",
	})
	assert.NotNil(t, authPassword)

	authMethod, err := authPassword.Create()
	assert.Error(t, err)
	assert.Nil(t, authMethod)
}
