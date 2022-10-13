package doomsday

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

func getInfo(ctx context.Context, c *cli.Context) error {
	client, err := newClient(ctx, c)
	if err != nil {
		return err
	}

	resp, err := client.Info()
	if err != nil {
		return err
	}

	info, err := json.Marshal(resp)

	fmt.Print(string(info))

	return nil
}
