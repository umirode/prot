package git

import (
	"errors"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type AuthBasicAuth struct {
	config map[string]string
}

func NewAuthBasicAuth(config map[string]string) *AuthBasicAuth {
	return &AuthBasicAuth{config: config}
}

func (a *AuthBasicAuth) Create() (transport.AuthMethod, error) {
	authConfig := struct {
		Username string
		Password string
	}{}

	var ok bool

	authConfig.Username, ok = a.config["Username"]
	if !ok {
		return nil, errors.New("parameter Username required for auth type Password")
	}

	authConfig.Password, ok = a.config["Password"]
	if !ok {
		return nil, errors.New("parameter Password required for auth type Password")
	}

	return &http.BasicAuth{
		Username: authConfig.Username,
		Password: authConfig.Password,
	}, nil
}
