package commands

import (
	"fmt"

	"github.com/Stoufiler/go-box/api"
	"github.com/Tomy2e/livebox-api-client"
	"github.com/urfave/cli/v2"
)

func LiveboxCommand() *cli.Command {
	return createBoxCommand("livebox", "Fetch infos from Livebox", liveboxAction)
}

func liveboxAction(client *livebox.Client) error {
	ont, err := api.GetOntInfos(client)
	if err != nil {
		return err
	}

	mac, err := api.GetMacAddress(client)
	if err != nil {
		return err
	}

	vlan, err := api.GetInternetVlan(client)
	if err != nil {
		return err
	}

	dhcp, err := api.GetDHCPInfos(client)
	if err != nil {
		return err
	}

	fmt.Println("\n============LEOX GPON COMMAND============")
	api.GenerateGponCommands(
		ont.Status.Gpon.Veip0.SerialNumber,
		ont.Status.Gpon.Veip0.VendorID,
		ont.Status.Gpon.Veip0.HardwareVersion,
		ont.Status.Gpon.Veip0.ONTSoftwareVersion0,
		ont.Status.Gpon.Veip0.ONTSoftwareVersion1,
	)
	fmt.Println("=========================================")
	fmt.Println("\n===========UDM PRO SE SETTINGS===========")
	api.DisplayUDMinfos(vlan, mac, dhcp)
	fmt.Println("=========================================")

	return nil
}
