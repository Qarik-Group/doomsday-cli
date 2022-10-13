package doomsday

import (
	"context"

	"github.com/urfave/cli/v2"
)

func GetCommands(ctx context.Context) []*cli.Command {
	var commands []*cli.Command

	// Certificates defines the certificates command
	var Certificates = &cli.Command{
		Name:        "certificates",
		Usage:       "Get certificates",
		Description: "Use 'doomsday certificates ...' to get certifictes from Doomsday.",
		Action: func(c *cli.Context) error {
			err := getAllCertificates(ctx, c)
			if err != nil {
				return err
			}
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:        "expired",
				Aliases:     []string{"e"},
				Usage:       "Get expired certificates",
				Description: "Get expired certificates",
				Action: func(c *cli.Context) error {
					err := getExpiredCertificates(ctx, c)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:        "willexpire",
				Aliases:     []string{"we"},
				Usage:       "Get certificates that will expire",
				Description: "Get certificates that will expire in given number of days",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:    "days",
						Aliases: []string{"d"},
						Value:   30,
						Usage:   "Number of days for which certificates will expire. Defaults to 30 days if not provided.",
					},
				},
				Action: func(c *cli.Context) error {
					err := getCertificatesThatWillExpire(ctx, c)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	// Refresh defines the refresh command
	var Refresh = &cli.Command{
		Name:        "refresh",
		Usage:       "Refresh certificates cache",
		Description: "Use 'doomsday refresh' to refresh certifictes cache.",
		Action: func(c *cli.Context) error {
			err := refreshCache(ctx, c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	// Info defines the info command
	var Info = &cli.Command{
		Name:        "info",
		Usage:       "Doomsday information",
		Description: "Use 'doomsday info' to get information about Doomsday version and authentication method.",
		Action: func(c *cli.Context) error {
			err := getInfo(ctx, c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	commands = append(commands, Certificates)
	commands = append(commands, Refresh)
	commands = append(commands, Info)

	return commands
}
