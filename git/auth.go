package git

import (
	"errors"
	"github.com/sirupsen/logrus"
	ssh2 "golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"strconv"
)

type authMethod interface {
	Create(map[string]string) (ssh.AuthMethod, error)
}

func GetAuthMethod(name string, config map[string]string) (ssh.AuthMethod, error) {
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
	}
}

type authPublicKeys struct {
}

func newAuthPublicKeys() *authPublicKeys {
	return &authPublicKeys{}
}

func (a *authPublicKeys) Create(config map[string]string) (ssh.AuthMethod, error) {
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
