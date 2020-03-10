package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/prot/config"
	"github.com/umirode/prot/git"
	"github.com/umirode/prot/tools"
	"github.com/urfave/cli/v2"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
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
		filledConfig, err := config.NewConfig(context.String("config"))
		if err != nil {
			return err
		}

		generator := tools.NewGenerator()
		cloner := git.NewCloner()

		outputPath, err := tools.CreateDirAndFormatPath(context.String("output"), false)
		if err != nil {
			return err
		}

		vendorDirPath := "prot_vendor"
		if outputPath != "" {
			vendorDirPath = outputPath + vendorDirPath
		}

		vendorDir, err := tools.CreateDirAndFormatPath(vendorDirPath, true)
		if err != nil {
			return err
		}

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
					}
				}

				outputModuleDirectory := vendorDir + strings.ToLower(name) + "/"
				err = cloner.Clone(outputModuleDirectory, module.Repository, authMethod)
				if err != nil {
					logrus.Error("Creating module directory error: ", err)
				}

				files, err := generator.GenerateProto(outputModuleDirectory, filledConfig.Lang)
				if err != nil {
					logrus.Error("Generating proto files error: ", err)
				}

				_ = files
			}()
		}

		wg.Wait()

		return nil
	},
}
