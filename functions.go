package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func postRequest(payload string) []byte {

	payloads := strings.NewReader(payload)

	req, _ := http.NewRequest("POST", URL, payloads)

	req.Header.Add("cookie", COOKIE)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", CONTEXTID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if body == nil {
		log.Fatal("Error while executing request to Livebox")
	}

	return body
}

func getOntInfos() {
	body := postRequest("{\"service\":\"NeMo.Intf.veip0\",\"method\":\"getMIBs\",\"parameters\":{\"mibs\":\"gpon\"}}")

	var ont Ont

	if err := json.Unmarshal(body, &ont); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	} else {
		GPON_SN = ont.Status.Gpon.Veip0.SerialNumber
		PON_VENDOR_ID = ont.Status.Gpon.Veip0.VendorID
		HW_HWVER = ont.Status.Gpon.Veip0.HardwareVersion
		OMCI_SW_VER1 = ont.Status.Gpon.Veip0.ONTSoftwareVersion0
		OMCI_SW_VER2 = ont.Status.Gpon.Veip0.ONTSoftwareVersion1
	}
}

// Fetch Mac Address
func getMacAddress() {
	body := postRequest("{\"service\":\"NMC\",\"method\":\"getWANStatus\",\"parameters\":{}}")

	var mac MacAddress

	if err := json.Unmarshal(body, &mac); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	} else {
		macaddress = mac.Data.MACAddress
	}
}

// Fetch Internet VLAN
func getInternetVlan() {
	body := postRequest("{\"service\":\"NeMo.Intf.data\",\"method\":\"getFirstParameter\",\"parameters\":{\"name\":\"VLANID\"}}")

	var vlaninternet VlanInternet

	if err := json.Unmarshal(body, &vlaninternet); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	} else {
		vlanid = strconv.Itoa(vlaninternet.Status)
	}
}

// Fetch DHCP Infos
func getDHCPInfos() {
	body := postRequest("{\"service\":\"NeMo.Intf.data\",\"method\":\"getMIBs\",\"parameters\":{\"mibs\":\"dhcp\"}}")

	var dhcpinfos DHCP

	if err := json.Unmarshal(body, &dhcpinfos); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}

	option70 := dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num77.Value
	option70decoded, err := hex.DecodeString(option70)
	if err != nil {
		fmt.Println("Erreur")
	}

	option70decodedstring := string(option70decoded)

	opt70 := strings.Split(option70decodedstring, "+")

	dhcpoption77 = opt70[1]

	// Add : every 2 char for DHCP Option 90
	var buffer bytes.Buffer
	var n_1 = 2 - 1
	var l_1 = len(dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num90.Value) - 1
	for i, rune := range dhcpinfos.Status.Dhcp.DhcpData.SentOption.Num90.Value {
		buffer.WriteRune(rune)
		if i%2 == n_1 && i != l_1 {
			buffer.WriteRune(':')
		}
	}
	dhcpoption90 = buffer.String()
}

func generateOMCC(oltvendorid string) {
	id := "128"
	dict := map[string]string{"HWTC": "131", "ALCL": "128"}

	if value, ok := dict[oltvendorid]; ok {
		id = value
	} else {
		fmt.Println("Unknown OLT VENDOR ID, defaulting to 128")
	}

	fmt.Println("\nExecute this command -> flash set OMCC_VER " + id + "\n")
}

func generateGponCommands() {
	var oltvendorid string

	fmt.Println("flash set GPON_PLOAM_PASSWD DEFAULT012")
	fmt.Println("flash set OMCI_TM_OPT 0")
	fmt.Println("flash set OMCI_OLT_MODE 1")
	fmt.Println("flash set GPON_SN " + GPON_SN)
	fmt.Println("flash set PON_VENDOR_ID " + PON_VENDOR_ID)
	fmt.Println("flash set HW_HWVER " + HW_HWVER)
	fmt.Println("flash set OMCI_SW_VER1 " + OMCI_SW_VER1)
	fmt.Println("flash set OMCI_SW_VER2 " + OMCI_SW_VER2)
	fmt.Println("\n## Unplug fiber from Livebox and plug it into UDM and wait a minute ##\n")
	fmt.Println("Execute this command -> omcicli mib get 131")

	// gnnYqBcYfn%5G!nY7@JDA@x

	fmt.Print("\nOLT VENDOR ID (HWTC, ALCL, ...) : ")
	fmt.Scan(&oltvendorid)
	generateOMCC(strings.ToUpper(oltvendorid))
}

func displayUDMinfos() {
	fmt.Println("NAME              : LEOX GPON")
	fmt.Println("VLAN ID           : " + vlanid)
	fmt.Println("MAC Address Clone : " + macaddress)
	fmt.Println("DHCP OPTION 60    : sagem")
	fmt.Println("DHCP OPTION 77    : " + dhcpoption77)
	fmt.Println("DHCP OPTION 90    : " + dhcpoption90)
	fmt.Println("DHCP CoS          : 6 ")
}
