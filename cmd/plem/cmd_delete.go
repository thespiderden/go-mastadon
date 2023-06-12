package main

import (
	"context"
	"errors"

	"spiderden.org/masta"
	"github.com/urfave/cli/v2"
)

func cmdDelete(c *cli.Context) error {
	client := c.App.Metadata["client"].(*masta.Client)
	if !c.Args().Present() {
		return errors.New("arguments required")
	}
	for i := 0; i < c.NArg(); i++ {
		err := client.DeleteStatus(context.Background(), masta.ID(c.Args().Get(i)))
		if err != nil {
			return err
		}
	}
	return nil
}
