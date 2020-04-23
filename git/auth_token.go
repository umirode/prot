package git

import (
	"errors"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type AuthToken struct {
	config map[string]string
}

func NewAuthToken(config map[string]string) *AuthToken {
	return &AuthToken{config: config}
}

func (a *AuthToken) Create() (transport.AuthMethod, error) {
	authConfig := struct {
		Token string
	}{}

	var ok bool

	authConfig.Token, ok = a.config["Token"]
	if !ok {
		return nil, errors.New("parameter Token required for auth type Token")
	}

	return &http.TokenAuth{
		Token: authConfig.Token,
	}, nil
}
