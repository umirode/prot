package cmd

import (
	"github.com/umirode/prot/config"
	"github.com/umirode/prot/tools"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		outputDir := path.Clean(context.String("output"))
		if !filepath.IsAbs(outputDir) {
			outputDir = tools.JoinPathAndFileName("", currentDir, outputDir)
		}
		err = tools.CreateDirRecursive(outputDir)
		if err != nil {
			return err
		}

		template, err := config.GetConfigTemplate()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(tools.JoinPathAndFileName("prot.yaml", outputDir), []byte(template), 0644)
		if err != nil {
			return err
		}

		return nil
	},
}
