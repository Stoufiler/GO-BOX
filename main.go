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

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

// Define essentially global variable to make requests once connected to Livebox
var CONTEXTID, URL, COOKIE string

// Define some global variables used on further functions
var GPON_SN, PON_VENDOR_ID, HW_HWVER, OMCI_SW_VER1, OMCI_SW_VER2, dhcpoption90, dhcpoption77, vlanid, macaddress string

func main() {
	app := &cli.App{
		Name:      "GO-BOX",
		Usage:     "CLI tool to fetch infos from Livebox/Funbox",
		UsageText: "GO-BOX --ip <box ip> --box <livebox/funbox>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "box",
				Usage:    "Type of orange box (Livebox, Funbox)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "ip",
				Usage:    "IP address of box",
				Required: true,
				Action: func(ctx *cli.Context, ip string) error {
					addr := net.ParseIP(ip)
					if addr == nil {
						return cli.Exit("Please provide a real IP Address", 86)
					} else if !addr.IsPrivate() {
						return cli.Exit("Please provide a private IP Address", 86)
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:   "username",
				Hidden: true,
				Value:  "admin",
				Usage:  "Username of admin management box",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Print("Password : ")
			password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatalln("Please provide a password")
			} else if len(password) == 0 {
				log.Fatalln("Password cannot be empty")
			}
			boxtype := strings.ToLower(ctx.String("box"))
			switch boxtype {
			case "livebox":
				// Concatenate the URL
				url := "http://" + ctx.String("ip") + "/ws"
				instantiateConnection(url, ctx.String("username"), string(password))
			case "funbox":
				instantiateFunboxConnection(ctx.String("ip"), ctx.String("username"), string(password))
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
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
	url := "http://" + ip + "/authenticate?username=" + username + "&password=" + string(password)

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

	var instanciation FunboxContext

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

	displayFunboxValues(ip)
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
