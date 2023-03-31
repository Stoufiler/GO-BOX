package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// Define essentially global variable to make requests once connected to Livebox
var CONTEXTID, URL, COOKIE string

// Define some global variables used on further functions
var GPON_SN, PON_VENDOR_ID, HW_HWVER, OMCI_SW_VER1, OMCI_SW_VER2, dhcpoption90, dhcpoption77, vlanid, macaddress string

func main() {
	var ip, username string
	var boxid uint8

	fmt.Println("Which box do you have ?")
	fmt.Println("1 - Livebox (DHCP Mode) ?")
	fmt.Println("2 - Funbox (PPPoE Mode) ?")
	fmt.Scan(&boxid)

	// Ask for IP
	fmt.Print("Livebox IP : ")
	fmt.Scan(&ip)

	// Check if IP is good and private
	addr := net.ParseIP(ip)
	if addr == nil {
		log.Fatalln("Please provide a real IP Address")
	} else if !addr.IsPrivate() {
		log.Fatalln("Please provide a private IP Address")
	}

	// Ask for Username
	fmt.Print("Username : ")
	fmt.Scan(&username)

	// Ask for Password without displaying it
	fmt.Print("Password : ")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln("Please provide a password")
	} else if len(password) == 0 {
		log.Fatalln("Password cannot be empty")
	}

	switch boxid {
	case 1:
		// Concatenate the URL
		url := "http://" + ip + "/ws"
		instantiateConnection(url, username, string(password))
	case 2:
		instantiateFunboxConnection(ip, username, string(password))
	default:
		log.Fatalln("Please provide a good option")
	}

}

// This function will instanciate a connection to Livebox to catch ContextID and Cookie
func instantiateConnection(url string, username string, password string) {
	payload := strings.NewReader("{\"service\":\"sah.Device.Information\",\"method\":\"createContext\",\"parameters\":{\"applicationName\":\"webui\",\"username\":\"" + username + "\",\"password\":\"" + password + "\"}}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "X-Sah-Login")
	req.Header.Add("Content-Type", "application/json")

	var netClient = &http.Client{
		Timeout: time.Second * 5,
	}

	res, err := netClient.Do(req)

	if err != nil {
		fmt.Println()
		log.Fatalln("Timeout (5s) exceeded while connecting to Livebox")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	cookie := (res.Header.Get("Set-Cookie"))
	str := strings.Split(cookie, ";")
	cookie = str[0]

	var instanciation Context

	if err := json.Unmarshal(body, &instanciation); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	if instanciation.Status != 0 {
		errors.New("!!! Unable to connect to Livebox !!! Please check your login/pass")
	} else {
		fmt.Println("")
		fmt.Println("✅ Successfully connected to Livebox ! ✅")
		CONTEXTID = instanciation.Data.ContextID
		COOKIE = cookie
		URL = url

		getOntInfos()
		getMacAddress()
		getInternetVlan()
		getDHCPInfos()

		displayNecessaryInformations()
	}
}

// This function will instanciate a connection to Livebox to catch ContextID and Cookie
func instantiateFunboxConnection(ip string, username string, password string) {
	url := "http://" + ip + "/authentication?username=" + username + "&password=" + string(password)

	req, _ := http.NewRequest("POST", url, nil)

	var netClient = &http.Client{
		Timeout: time.Second * 5,
	}

	res, err := netClient.Do(req)

	if err != nil {
		fmt.Println()
		log.Fatalln("Timeout (5s) exceeded while connecting to Livebox")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	cookie := (res.Header.Get("Set-Cookie"))
	str := strings.Split(cookie, ";")
	cookie = str[0]

	var instanciation Context

	if err := json.Unmarshal(body, &instanciation); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	if instanciation.Status != 0 {
		errors.New("!!! Unable to connect to Livebox !!! Please check your login/pass")
	} else {
		fmt.Println("")
		fmt.Println("✅ Successfully connected to Livebox ! ✅")
		CONTEXTID = instanciation.Data.ContextID
		COOKIE = cookie
		URL = url
	}
}

func displayNecessaryInformations() {
	fmt.Println("")
	fmt.Println("===========LEOX GPON COMMAND=============")
	generateGponCommands()
	fmt.Println("=========================================")
	fmt.Println("")
	fmt.Println("==========UDM PRO SE SETTINGS============")
	displayUDMinfos()
	fmt.Println("=========================================")
}
