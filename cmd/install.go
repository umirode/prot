package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/prot/config"
	"github.com/umirode/prot/git"
	"github.com/umirode/prot/tools"
	"github.com/urfave/cli/v2"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var InstallCmd = &cli.Command{
	Name:  "install",
	Usage: "Install proto dependencies from config file ",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Value:       "prot.yaml",
			Aliases:     []string{"c"},
			Usage:       "Load configuration from `FILE`",
			DefaultText: "prot.yaml",
		},
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
			outputDir = tools.JoinPathAndFileName("", currentDir, outputDir, "prot_vendor")
		}
		err = tools.CreateDirRecursive(outputDir)
		if err != nil {
			return err
		}

		configPath := path.Clean(context.String("config"))
		if !filepath.IsAbs(configPath) {
			configPath = tools.JoinPathAndFileName(configPath, currentDir)
		}

		filledConfig, err := config.NewConfig(configPath)
		if err != nil {
			return err
		}

		generator := tools.NewGenerator()
		cloner := git.NewCloner()

		wg := sync.WaitGroup{}

		for name, module := range filledConfig.Modules {
			wg.Add(1)
			go func() {
				defer wg.Done()

				var authMethod ssh.AuthMethod
				if module.Auth != nil {
					var err error
					authMethod, err = git.GetAuthMethod((*module.Auth).Type, (*module.Auth).Config)
					if err != nil {
						logrus.Error("Creating auth method error: ", err)
						return
					}
				}

				moduleDir := tools.JoinPathAndFileName("", outputDir, strings.ToLower(name))
				err = cloner.Clone(moduleDir, module.Repository, authMethod)
				if err != nil {
					logrus.Error("Creating module directory error: ", err)
					return
				}

				files, err := generator.GenerateProto(moduleDir, filledConfig.Lang)
				if err != nil {
					logrus.Error("Generating proto files error: ", err)
					return
				}

				_ = files
			}()
		}

		wg.Wait()

		return nil
	},
}
