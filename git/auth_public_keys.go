package git

import (
	"errors"
	"github.com/sirupsen/logrus"
	ssh2 "golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"strconv"
)

type AuthPublicKeys struct {
	config map[string]string
}

func NewAuthPublicKeys(config map[string]string) *AuthPublicKeys {
	return &AuthPublicKeys{config: config}
}

func (a *AuthPublicKeys) Create() (transport.AuthMethod, error) {
	authConfig := struct {
		PemFile         string
		PemFilePassword string
		User            string
		IgnoreHostKey   bool
	}{}

	var ok bool
	var err error

	authConfig.PemFile, ok = a.config["PemFile"]
	if !ok {
		return nil, errors.New("parameter PemFile required for auth type PublicKeys")
	}

	authConfig.User, ok = a.config["User"]
	if !ok || authConfig.User == "" {
		authConfig.User = "git"
	}

	authConfig.PemFilePassword, _ = a.config["PemFilePassword"]

	ignoreHostKeyString, ok := a.config["IgnoreHostKey"]
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
