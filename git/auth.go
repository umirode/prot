package git

import (
	"errors"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

type AuthMethod interface {
	Create() (transport.AuthMethod, error)
}

type AuthMethodFactory struct {
	authMethods map[string]AuthMethod
}

func NewAuthMethodFactory(authMethods map[string]AuthMethod) *AuthMethodFactory {
	return &AuthMethodFactory{authMethods: authMethods}
}

func (amf *AuthMethodFactory) GetAuthMethod(name string) (transport.AuthMethod, error) {
	AuthMethod, ok := amf.authMethods[name]
	if !ok {
		return nil, errors.New("auth method " + name + " not exists")
	}

	return AuthMethod.Create()
}

//func (amf *AuthMethodFactory) authMethodGetAuthMethodMap() map[string]AuthMethod {
//	return map[string]AuthMethod{
//		"PublicKeys": newAuthPublicKeys(),
//		"Password":   newAuthPassword(),
//		"BasicAuth":  newAuthBasicAuth(),
//		"Token":      newAuthToken(),
//	}
//}
