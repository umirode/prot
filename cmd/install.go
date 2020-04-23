package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/prot/config"
	"github.com/umirode/prot/git"
	"github.com/umirode/prot/helpers"
	"github.com/umirode/prot/protobuf"
	"github.com/urfave/cli/v2"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var InstallCmd = &cli.Command{
	Name:  "install",
	Usage: "Install dependencies from config file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Value:       "prot.yaml",
			Aliases:     []string{"c"},
			Usage:       "Path to configuration file",
			DefaultText: "prot.yaml",
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Output path",
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
			outputDir = helpers.JoinPathAndFileName("", currentDir, outputDir, "prot_vendor")
		}
		err = helpers.CreateDirRecursive(outputDir)
		if err != nil {
			return err
		}

		configPath := path.Clean(context.String("config"))
		if !filepath.IsAbs(configPath) {
			configPath = helpers.JoinPathAndFileName(configPath, currentDir)
		}

		filledConfig, err := config.NewConfig(configPath)
		if err != nil {
			return err
		}

		generator := protobuf.NewGenerator()
		cloner := git.NewCloner()

		wg := sync.WaitGroup{}

		for name, module := range filledConfig.Modules {
			wg.Add(1)
			go func() {
				defer wg.Done()

				var authMethod transport.AuthMethod
				if module.Auth != nil {
					authConfig := (*module.Auth).Config

					var err error
					authMethod, err = git.NewAuthMethodFactory(map[string]git.AuthMethod{
						"PublicKeys": git.NewAuthPublicKeys(authConfig),
						"Password":   git.NewAuthPassword(authConfig),
						"BasicAuth":  git.NewAuthBasicAuth(authConfig),
						"Token":      git.NewAuthToken(authConfig),
					}).GetAuthMethod((*module.Auth).Type)
					if err != nil {
						logrus.Error("creating auth method error: ", err)
						return
					}
				}

				moduleDir := helpers.JoinPathAndFileName("", outputDir, strings.ToLower(name))
				err := helpers.CleanDir(moduleDir)
				if err != nil {
					logrus.Error("module directory cleaning error: ", err)
					return
				}

				err = cloner.Clone(moduleDir, module.Repository, authMethod)
				if err != nil {
					logrus.Error("creating module directory error: ", err)
					return
				}

				files, err := generator.GenerateProto(moduleDir, filledConfig.Lang)
				if err != nil {
					_ = helpers.CleanDir(moduleDir)
					logrus.Error("generating proto files error: ", err)
					return
				}

				_ = files
			}()
		}

		wg.Wait()

		return nil
	},
}
