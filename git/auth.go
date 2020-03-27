package git

import (
	"errors"
	"github.com/sirupsen/logrus"
	ssh2 "golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"strconv"
)

type authMethod interface {
	Create(map[string]string) (transport.AuthMethod, error)
}

func GetAuthMethod(name string, config map[string]string) (transport.AuthMethod, error) {
	authMethodMap := authMethodGetAuthMethodMap()

	authMethod, ok := authMethodMap[name]
	if !ok {
		return nil, errors.New("auth method " + name + " not exists")
	}

	return authMethod.Create(config)
}

func authMethodGetAuthMethodMap() map[string]authMethod {
	return map[string]authMethod{
		"PublicKeys": newAuthPublicKeys(),
		"Password":   newAuthPassword(),
		"BasicAuth":  newAuthBasicAuth(),
	}
}

type authPublicKeys struct {
}

func newAuthPublicKeys() *authPublicKeys {
	return &authPublicKeys{}
}

func (a *authPublicKeys) Create(config map[string]string) (transport.AuthMethod, error) {
	authConfig := struct {
		PemFile         string
		PemFilePassword string
		User            string
		IgnoreHostKey   bool
	}{}

	var ok bool
	var err error

	authConfig.PemFile, ok = config["PemFile"]
	if !ok {
		return nil, errors.New("parameter PemFile required for auth type PublicKeys")
	}

	authConfig.User, ok = config["User"]
	if !ok || authConfig.User == "" {
		authConfig.User = "git"
	}

	authConfig.PemFilePassword, _ = config["PemFilePassword"]

	ignoreHostKeyString, ok := config["IgnoreHostKey"]
	if ok {
		authConfig.IgnoreHostKey, err = strconv.ParseBool(ignoreHostKeyString)
		if err != nil {
			return nil, errors.New("parameter IgnoreHostKey must be true or false for auth type PublicKey")
		}
	}

	auth, err := ssh.NewPublicKeysFromFile(authConfig.User, authConfig.PemFile, authConfig.PemFilePassword)
	if err != nil {
		logrus.Fatal(err)
	}

	if authConfig.IgnoreHostKey {
		auth.HostKeyCallback = ssh2.InsecureIgnoreHostKey()
	}

	return auth, nil
}

type authPassword struct {
}

func newAuthPassword() *authPassword {
	return &authPassword{}
}

func (a *authPassword) Create(config map[string]string) (transport.AuthMethod, error) {
	authConfig := struct {
		User     string
		Password string
	}{}

	var ok bool

	authConfig.User, ok = config["User"]
	if !ok {
		return nil, errors.New("parameter User required for auth type Password")
	}

	authConfig.Password, ok = config["Password"]
	if !ok {
		return nil, errors.New("parameter Password required for auth type Password")
	}

	return &ssh.Password{
		User:     authConfig.User,
		Password: authConfig.Password,
	}, nil
}

type authBasicAuth struct {
}

func newAuthBasicAuth() *authBasicAuth {
	return &authBasicAuth{}
}

func (a *authBasicAuth) Create(config map[string]string) (transport.AuthMethod, error) {
	authConfig := struct {
		Username string
		Password string
	}{}

	var ok bool

	authConfig.Username, ok = config["Username"]
	if !ok {
		return nil, errors.New("parameter Username required for auth type Password")
	}

	authConfig.Password, ok = config["Password"]
	if !ok {
		return nil, errors.New("parameter Password required for auth type Password")
	}

	return &http.BasicAuth{
		Username: authConfig.Username,
		Password: authConfig.Password,
	}, nil
}
