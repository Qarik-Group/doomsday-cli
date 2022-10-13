package doomsday

import (
	"context"
	"net/url"

	"github.com/doomsday-project/doomsday/client/doomsday"
	"github.com/urfave/cli/v2"
)

func newClient(ctx context.Context, c *cli.Context) (doomsday.Client, error) {
	client := doomsday.Client{}

	url, err := url.Parse(c.String("server"))
	if err != nil {
		return client, err
	}
	client.URL = *url

	client.UserpassAuth(c.String("username"), c.String("password"))

	return client, nil
}
