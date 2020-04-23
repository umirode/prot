package git

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"testing"
)

type emptyAuthMethod struct {
	config map[string]string
}

func (e emptyAuthMethod) Create() (transport.AuthMethod, error) {
	return nil, nil
}

func TestAuthMethodFactory(t *testing.T) {
	authMethodFactory := NewAuthMethodFactory(map[string]AuthMethod{
		"empty": &emptyAuthMethod{
			config: map[string]string{},
		},
	})

	_, err := authMethodFactory.GetAuthMethod("empty")
	assert.NoError(t, err)
}

func TestAuthMethodFactoryError(t *testing.T) {
	authMethodFactory := NewAuthMethodFactory(map[string]AuthMethod{})

	_, err := authMethodFactory.GetAuthMethod("test")
	assert.Error(t, err)
}
