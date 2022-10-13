package main

import (
	"context"
	"log"
	"os"

	doomsday "github.com/starkandwayne/doomsday-cli/cli"

	"github.com/urfave/cli/v2"
)

var (
	appVersion = "v0.0.1"
)

func main() {
	ctx := context.Background()

	app := &cli.App{
		Name:     "Doomsday API CLI",
		Usage:    "A tool that helps interact with Doomsday API",
		Commands: doomsday.GetCommands(ctx),
		Version:  appVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server",
				Aliases: []string{"s"},
				Value:   "https://localhost:443",
				Usage:   "Doomsday API server URL",
				EnvVars: []string{"DOOMSDAY_SERVER"},
			},
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Value:   "doomsday",
				Usage:   "API server username",
				EnvVars: []string{"DOOMSDAY_USER"},
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "API server password",
				EnvVars: []string{"DOOMSDAY_PASSWORD"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
