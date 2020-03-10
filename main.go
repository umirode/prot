package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/umirode/prot/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	fmt.Println("test")
	app := &cli.App{
		Commands: []*cli.Command{
			cmd.InstallCmd,
			cmd.InitCmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}

	os.Exit(0)
}
