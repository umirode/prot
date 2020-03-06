package cmd

import (
	"github.com/umirode/prot/config"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Generate prot.yml config in current directory",
	Action: func(context *cli.Context) error {
		template, err := config.GetConfigTemplate()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("prot.yaml", []byte(template), 0644)
		if err != nil {
			return err
		}

		return nil
	},
}
