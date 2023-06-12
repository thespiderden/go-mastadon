package main

import (
	"context"
	"errors"
	"fmt"

	"spiderden.org/masta"
	"github.com/urfave/cli/v2"
)

func cmdToot(c *cli.Context) error {
	var toot string
	ff := c.String("ff")
	if ff != "" {
		text, err := readFile(ff)
		if err != nil {
			return err
		}
		toot = string(text)
	} else {
		if !c.Args().Present() {
			return errors.New("arguments required")
		}
		toot = argstr(c)
	}
	client := c.App.Metadata["client"].(*masta.Client)
	_, err := client.PostStatus(context.Background(), &masta.Toot{
		Status:      toot,
		InReplyToID: masta.ID(fmt.Sprint(c.String("i"))),
	})
	return err
}
