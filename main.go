package main

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/prot/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Usage:   "protobuf files manager",
		Version: "1.0.0",
		Commands: []*cli.Command{
			cmd.InstallCmd,
			cmd.InitCmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
