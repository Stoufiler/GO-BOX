package commands

import (
	"fmt"

	"github.com/Stoufiler/go-box/api"
	"github.com/Tomy2e/livebox-api-client"
	"github.com/urfave/cli/v2"
)

type boxAction func(client *livebox.Client) error

func createBoxCommand(name, usage string, action boxAction) *cli.Command {
	return &cli.Command{
		Name:  name,
		Usage: usage,
		Action: func(ctx *cli.Context) error {
			client, err := api.NewClient(ctx.String("ip"), ctx.String("password"))
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			fmt.Printf("\n✅ Successfully connected to %s ! ✅\n", name)

			if err := action(client); err != nil {
				return fmt.Errorf("error executing %s command: %w", name, err)
			}

			return nil
		},
	}
}