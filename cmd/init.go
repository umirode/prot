package cmd

import (
	"github.com/umirode/prot/config"
	"github.com/umirode/prot/tools"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Generate prot.yml config in current directory",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Output path `PATH`",
			DefaultText: "current directory",
		},
	},
	Action: func(context *cli.Context) error {
		outputPath, err := tools.CreateDirAndFormatPath(context.String("output"), false)
		if err != nil {
			return err
		}

		protConfigPath := "prot.yaml"
		if outputPath != "" {
			protConfigPath = outputPath + protConfigPath
		}

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
