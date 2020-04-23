package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthPublicKeys(t *testing.T) {
	authPublicKeys := NewAuthPublicKeys(map[string]string{
		"PemFile":         "pem_file",
		"PemFilePassword": "pem_file_password",
		"User":            "user",
		"Password":        "pass",
	})
	assert.NotNil(t, authPublicKeys)

	//authMethod, err := authPublicKeys.Create()
	//assert.NoError(t, err)
	//assert.NotNil(t, authMethod)
	//assert.Equal(t, &ssh.Password{
	//	User: "user",
	//	Password: "pass",
	//}, authMethod)
}
