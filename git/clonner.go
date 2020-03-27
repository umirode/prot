package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"os"
)

type Cloner struct {
}

func NewCloner() *Cloner {
	return &Cloner{}
}

func (*Cloner) Clone(outputPath string, moduleUrl string, auth transport.AuthMethod) error {
	cloneOptions := &git.CloneOptions{
		URL:      moduleUrl,
		Progress: os.Stdout,
	}

	if auth != nil {
		cloneOptions.Auth = auth
	}

	_, err := git.PlainClone(outputPath, false, cloneOptions)
	if err != nil {
		return err
	}

	return nil
}
