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
	},
	Action: func(context *cli.Context) error {
		filledConfig, err := config.NewConfig(context.String("config"))
		if err != nil {
			return err
		}

		generator := tools.NewGenerator()
		cloner := git.NewCloner()

		vendorDir := tools.CreateDirAndFormatPath("prot_vendor")

		wg := sync.WaitGroup{}

		for name, module := range filledConfig.Modules {
			wg.Add(1)
			go func() {
				defer wg.Done()

				var authMethod ssh.AuthMethod
				if (module.AuthType != nil) && (module.AuthConfig != nil) {
					var err error
					authMethod, err = git.GetAuthMethod(*module.AuthType, *module.AuthConfig)
					if err != nil {
						logrus.Error(err)
					}
				}

				outputModuleDirectory := vendorDir + "proto_" + strings.ToLower(name) + "/"
				err = cloner.Clone(outputModuleDirectory, module.Url, authMethod)
				if err != nil {
					logrus.Error(err)
				}

				files, err := generator.GenerateProto(outputModuleDirectory, filledConfig.Lang)
				if err != nil {
					logrus.Error(err)
				}

				_ = files
			}()
		}

		wg.Wait()

		return nil
	},
}
