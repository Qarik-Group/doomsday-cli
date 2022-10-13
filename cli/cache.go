package doomsday

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
)

func refreshCache(ctx context.Context, c *cli.Context) error {
	client, err := newClient(ctx, c)
	if err != nil {
		return err
	}

	err = client.RefreshCache()
	if err != nil {
		return err
	}

	fmt.Print("Certificates cache has been refreshed")

	return nil
}
