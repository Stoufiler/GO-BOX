package api

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Tomy2e/livebox-api-client"
	"github.com/Tomy2e/livebox-api-client/api/request"
)

// NewClient creates a new Livebox API client.
func NewClient(ip, password string) (*livebox.Client, error) {
	client, err := livebox.NewClient(password, livebox.WithAddress("http://"+ip))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	return client, nil
}

// GetFunboxValues fetches values from a Funbox.
func GetFunboxValues(client *livebox.Client) (*FunboxValues, error) {
	var fbvalues FunboxValues
	if err := client.Request(context.Background(), request.New("NeMo.Intf.data", "getMIBs", nil), &fbvalues); err != nil {
		return nil, fmt.Errorf("could not get funbox values: %w", err)
	}
	return &fbvalues, nil
}

// GetOntInfos fetches ONT information from the Livebox.
func GetOntInfos(client *livebox.Client) (*Ont, error) {
	var ont Ont
	if err := client.Request(context.Background(), request.New("NeMo.Intf.veip0", "getMIBs", map[string]interface{}{"mibs": "gpon"}), &ont); err != nil {
		return nil, fmt.Errorf("could not get ONT infos: %w", err)
	}
	return &ont, nil
}

// GetMacAddress fetches the MAC address of the Livebox.
func GetMacAddress(client *livebox.Client) (string, error) {
	var mac MacAddress
	if err := client.Request(context.Background(), request.New("NMC", "getWANStatus", nil), &mac); err != nil {
		return "", fmt.Errorf("could not get MAC address: %w", err)
	}
	return mac.Data.MACAddress, nil
}

// GetInternetVlan fetches the VLAN ID for the internet connection.
func GetInternetVlan(client *livebox.Client) (int, error) {
	var vlaninternet VlanInternet
	if err := client.Request(context.Background(), request.New("NeMo.Intf.data", "getFirstParameter", map[string]interface{}{"name": "VLANID"}), &vlaninternet); err != nil {
		return 0, fmt.Errorf("could not get internet vlan: %w", err)
	}
	return vlaninternet.Status, nil
}

// GetDHCPInfos fetches DHCP information from the Livebox.
func GetDHCPInfos(client *livebox.Client) (*DHCPInfo, error) {
	var dhcpinfos DHCP
	if err := client.Request(context.Background(), request.New("NeMo.Intf.data", "getMIBs", map[string]interface{}{"mibs": "dhcp"}), &dhcpinfos); err != nil {
		return nil, fmt.Errorf("could not get DHCP infos: %w", err)
	}

	option60 := dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num60.Value
	option60decoded, err := hex.DecodeString(option60)
	if err != nil {
		return nil, fmt.Errorf("could not decode option 60: %w", err)
	}

	option70 := dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num77.Value
	option70decoded, err := hex.DecodeString(option70)
	if err != nil {
		return nil, fmt.Errorf("could not decode option 70: %w", err)
	}

	option70decodedstring := string(option70decoded)

	opt70 := strings.Split(option70decodedstring, "+")

	var dhcpoption77 string
	if len(opt70) > 1 {
		dhcpoption77 = opt70[1]
	} else {
		dhcpoption77 = option70decodedstring
	}

	// Add : every 2 char for DHCP Option 90
	var buffer bytes.Buffer
	var n1 = 2 - 1
	var l1 = len(dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num90.Value) - 1
	for i, runner := range dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num90.Value {
		buffer.WriteRune(runner)
		if i%2 == n1 && i != l1 {
			buffer.WriteRune(':')
		}
	}
	dhcpoption90 := buffer.String()

	return &DHCPInfo{
		Option60: string(option60decoded),
		Option77: dhcpoption77,
		Option90: dhcpoption90,
	}, nil
}

var oltVendorIDMap = map[string]string{
	"HWTC": "136",
	"ALCL": "128",
	"130":  "130", // Placeholder for your OLT Vendor ID
}

func generateOMCC(oltvendorid, hwHwver, omciSwVer1, omciSwVer2 string) {
	id, ok := oltVendorIDMap[oltvendorid]
	if !ok {
		id = "128" // default value
		fmt.Println("Unknown OLT VENDOR ID, defaulting to 128")
	}

	fmt.Println("Execute this command -> flash set OMCC_VER " + id)

	if oltvendorid == "ALCL" {
		if len(hwHwver) != 0 {
			fmt.Println("flash set HW_HWVER " + hwHwver)
		}
		fmt.Println("flash set OMCI_SW_VER1 " + omciSwVer1)
		fmt.Println("flash set OMCI_SW_VER2 " + omciSwVer2)
		fmt.Println("flash set OMCI_TM_OPT 0")
		fmt.Println("flash set OMCI_OLT_MODE 1")
	}
}

// GenerateGponCommands prints the GPON commands.
func GenerateGponCommands(gponSn, ponVendorId, hwHwver, omciSwVer1, omciSwVer2 string) {
	//fmt.Println("flash set GPON_PLOAM_PASSWD DEFAULT012")
	fmt.Println("flash set GPON_SN " + gponSn)
	fmt.Println("flash set PON_VENDOR_ID " + ponVendorId)
	fmt.Println("## Unplug fiber from Livebox and plug it into UDM and wait a minute ##")
	fmt.Println("Execute this command -> omcicli mib get 131")

	fmt.Print("OLT VENDOR ID (HWTC, ALCL, ...) : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		oltvendorid := scanner.Text()
		generateOMCC(strings.ToUpper(strings.TrimSpace(oltvendorid)), hwHwver, omciSwVer1, omciSwVer2)
	}
}

// DisplayUDMinfos prints the UDM information.
func DisplayUDMinfos(vlanid int, macaddress string, dhcpinfos *DHCPInfo) {
	fmt.Println("NAME              : LEOX GPON")
	fmt.Println("VLAN ID           : " + strconv.Itoa(vlanid))
	fmt.Println("MAC Address Clone : " + macaddress)
	fmt.Println("DHCP OPTION 60    : " + dhcpinfos.Option60)
	fmt.Println("DHCP OPTION 77    : " + dhcpinfos.Option77)
	fmt.Println("DHCP OPTION 90    : " + dhcpinfos.Option90)
	fmt.Println("DHCP CoS          : 6 (Requires Network App 7.4.X or later)")
}

// DisplayFunboxValues prints the Funbox values.
func DisplayFunboxValues(client *livebox.Client) error {
	fbvalues, err := GetFunboxValues(client)
	if err != nil {
		return err
	}

	fmt.Println("===========UDM PRO SE SETTINGS===========")
	fmt.Println("PPPoE Username : " + fbvalues.Result.Status.Ppp.PppData.Username)
	fmt.Println("PPPoE Password : https://www.orange.pl/moj-orange -> Ustawienia -> Zabezpieczenia -> Hasło do Neostrady -> Zmień hasło")
	fmt.Println("VLAN ID: " + strconv.Itoa(fbvalues.Result.Status.Vlan.GvlanData.Vlanid))
	fmt.Println("=========================================")

	gponSn := fbvalues.Result.Status.Gpon.Veip0.SerialNumber
	ponVendorId := fbvalues.Result.Status.Gpon.Veip0.SerialNumber[0:4]
	hwHwver := fbvalues.Result.Status.Gpon.Veip0.HardwareVersion
	omciSwVer1 := fbvalues.Result.Status.Gpon.Veip0.ONTSoftwareVersion0
	omciSwVer2 := fbvalues.Result.Status.Gpon.Veip0.ONTSoftwareVersion1

	fmt.Println("")
	fmt.Println("============LEOX GPON COMMAND============")
	GenerateGponCommands(gponSn, ponVendorId, hwHwver, omciSwVer1, omciSwVer2)
	fmt.Println("=========================================")
	return nil
}
