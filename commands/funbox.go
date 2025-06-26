package commands

import (
	"github.com/Stoufiler/go-box/api"
	"github.com/Tomy2e/livebox-api-client"
	"github.com/urfave/cli/v2"
)

func FunboxCommand() *cli.Command {
	return createBoxCommand("funbox", "Fetch infos from Funbox", funboxAction)
}

func funboxAction(client *livebox.Client) error {
	return api.DisplayFunboxValues(client)
}
