package doomsday

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/doomsday-project/doomsday/client/doomsday"
	"github.com/urfave/cli/v2"
)

func getCertificates(ctx context.Context, c *cli.Context) (doomsday.CacheItems, error) {
	var certificates doomsday.CacheItems

	client, err := newClient(ctx, c)
	if err != nil {
		return certificates, err
	}

	certificates, err = client.GetCache()
	if err != nil {
		return certificates, err
	}

	return certificates, nil
}

func getAllCertificates(ctx context.Context, c *cli.Context) error {
	certificates, err := getCertificates(ctx, c)
	if err != nil {
		return err
	}

	allCertificates, err := json.Marshal(certificates)

	fmt.Print(string(allCertificates))

	return nil
}

func getExpiredCertificates(ctx context.Context, c *cli.Context) error {
	certificates, err := getCertificates(ctx, c)
	if err != nil {
		return err
	}

	var within time.Duration = 0

	if err != nil {
		return err
	}

	filter := doomsday.CacheItemFilter{Within: &within}

	certificates = certificates.Filter(filter)

	expiredCertificates, err := json.Marshal(certificates)

	fmt.Print(string(expiredCertificates))

	return nil
}

func getCertificatesThatWillExpire(ctx context.Context, c *cli.Context) error {
	certificates, err := getCertificates(ctx, c)
	if err != nil {
		return err
	}

	var beyond time.Duration = 1000000000 * 0
	var within time.Duration = 1000000000 * 3600 * 24 * time.Duration(c.Int64("days"))

	if err != nil {
		return err
	}

	filter := doomsday.CacheItemFilter{Within: &within, Beyond: &beyond}

	certificates = certificates.Filter(filter)

	expiredCertificates, err := json.Marshal(certificates)

	fmt.Print(string(expiredCertificates))

	return nil
}
