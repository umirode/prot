package cmd

import (
	"github.com/umirode/prot/config"
	"github.com/umirode/prot/helpers"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Generate config for Prot application",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Output path for config",
			DefaultText: "current directory",
		},
	},
	Action: func(context *cli.Context) error {
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		outputDir := path.Clean(context.String("output"))
		if !filepath.IsAbs(outputDir) {
			outputDir = helpers.JoinPathAndFileName("", currentDir, outputDir)
		}
		err = helpers.CreateDirRecursive(outputDir)
		if err != nil {
			return err
		}

		template, err := config.GetConfigTemplate()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(helpers.JoinPathAndFileName("prot.yaml", outputDir), []byte(template), 0644)
		if err != nil {
			return err
		}

		return nil
	},
}
