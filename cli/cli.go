package cli

import (
	"log"
	"os"

	"github.com/Stoufiler/go-box/commands"
	"github.com/Stoufiler/go-box/utils"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:  "GO-BOX",
		Usage: "CLI tool to fetch infos from Livebox/Funbox",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "ip",
				Usage:    "IP address of box",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Password of box",
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.String("password") == "" {
				password, err := utils.GetPassword()
				if err != nil {
					return err
				}
				return ctx.Set("password", password)
			}
			return nil
		},
		Commands: []*cli.Command{
			commands.LiveboxCommand(),
			commands.FunboxCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
