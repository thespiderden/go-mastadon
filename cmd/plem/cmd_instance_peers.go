package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
	"spiderden.org/masta"
)

func cmdInstancePeers(c *cli.Context) error {
	client := c.App.Metadata["client"].(*masta.Client)
	peers, err := client.GetInstancePeers(context.Background())
	if err != nil {
		return err
	}
	for _, peer := range peers {
		fmt.Fprintln(c.App.Writer, peer)
	}
	return nil
}
