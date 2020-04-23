package git

import (
	"errors"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

type AuthPassword struct {
	config map[string]string
}

func NewAuthPassword(config map[string]string) *AuthPassword {
	return &AuthPassword{config: config}
}

func (a *AuthPassword) Create() (transport.AuthMethod, error) {
	authConfig := struct {
		User     string
		Password string
	}{}

	var ok bool

	authConfig.User, ok = a.config["User"]
	if !ok {
		return nil, errors.New("parameter User required for auth type Password")
	}

	authConfig.Password, ok = a.config["Password"]
	if !ok {
		return nil, errors.New("parameter Password required for auth type Password")
	}

	return &ssh.Password{
		User:     authConfig.User,
		Password: authConfig.Password,
	}, nil
}
