package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// Define some global variables used on further functions
var leoxCommands []string
var fti, ftipass, GPON_SN, PON_VENDOR_ID, HW_HWVER, OMCI_SW_VER1, OMCI_SW_VER2, dhcpoption90, dhcpoption77, vlanid, macaddress string

// Define struct to handle http requests : https://mholt.github.io/json-to-go/
type Context struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	} `json:"data"`
}

type Ont struct {
	Status struct {
		Gpon struct {
			Veip0 struct {
				RegistrationID           string `json:"RegistrationID"`
				VeipPptpUni              bool   `json:"VeipPptpUni"`
				OmciIsTmOwner            bool   `json:"OmciIsTmOwner"`
				MaxBitRateSupported      int    `json:"MaxBitRateSupported"`
				SignalRxPower            int    `json:"SignalRxPower"`
				SignalTxPower            int    `json:"SignalTxPower"`
				Temperature              int    `json:"Temperature"`
				Voltage                  int    `json:"Voltage"`
				Bias                     int    `json:"Bias"`
				SerialNumber             string `json:"SerialNumber"`
				HardwareVersion          string `json:"HardwareVersion"`
				EquipmentID              string `json:"EquipmentId"`
				VendorID                 string `json:"VendorId"`
				VendorProductCode        int    `json:"VendorProductCode"`
				PonID                    string `json:"PonId"`
				ONTSoftwareVersion0      string `json:"ONTSoftwareVersion0"`
				ONTSoftwareVersion1      string `json:"ONTSoftwareVersion1"`
				ONTSoftwareVersionActive int    `json:"ONTSoftwareVersionActive"`
				ONUState                 string `json:"ONUState"`
				DownstreamMaxRate        int    `json:"DownstreamMaxRate"`
				UpstreamMaxRate          int    `json:"UpstreamMaxRate"`
				DownstreamCurrRate       int    `json:"DownstreamCurrRate"`
				UpstreamCurrRate         int    `json:"UpstreamCurrRate"`
			} `json:"veip0"`
		} `json:"gpon"`
	} `json:"status"`
}

type Mtu struct {
	Status int `json:"status"`
}

type MacAddress struct {
	Status bool `json:"status"`
	Data   struct {
		WanState            string `json:"WanState"`
		LinkType            string `json:"LinkType"`
		LinkState           string `json:"LinkState"`
		GponState           string `json:"GponState"`
		MACAddress          string `json:"MACAddress"`
		Protocol            string `json:"Protocol"`
		ConnectionState     string `json:"ConnectionState"`
		LastConnectionError string `json:"LastConnectionError"`
		IPAddress           string `json:"IPAddress"`
		RemoteGateway       string `json:"RemoteGateway"`
		DNSServers          string `json:"DNSServers"`
		IPv6Address         string `json:"IPv6Address"`
	} `json:"data"`
}

type VlanInternet struct {
	Status int `json:"status"`
}

type DHCP struct {
	Status struct {
		Dhcp struct {
			DhcpData struct {
				DHCPStatus                 string `json:"DHCPStatus"`
				LastConnectionError        string `json:"LastConnectionError"`
				Renew                      bool   `json:"Renew"`
				IPAddress                  string `json:"IPAddress"`
				SubnetMask                 string `json:"SubnetMask"`
				IPRouters                  string `json:"IPRouters"`
				DNSServers                 string `json:"DNSServers"`
				DHCPServer                 string `json:"DHCPServer"`
				LeaseTime                  int    `json:"LeaseTime"`
				LeaseTimeRemaining         int    `json:"LeaseTimeRemaining"`
				Uptime                     int    `json:"Uptime"`
				DSCPMark                   int    `json:"DSCPMark"`
				PriorityMark               int    `json:"PriorityMark"`
				Formal                     bool   `json:"Formal"`
				BroadcastFlag              int    `json:"BroadcastFlag"`
				CheckAuthentication        bool   `json:"CheckAuthentication"`
				AuthenticationInformation  string `json:"AuthenticationInformation"`
				ResetOnPhysDownTimeout     int    `json:"ResetOnPhysDownTimeout"`
				RetransmissionStrategy     string `json:"RetransmissionStrategy"`
				RetransmissionRenewTimeout int    `json:"RetransmissionRenewTimeout"`
				SendMaxMsgSize             bool   `json:"SendMaxMsgSize"`
				SentOption                 struct {
					Num60 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"60"`
					Num61 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"61"`
					Num77 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"77"`
					Num90 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"90"`
				} `json:"SentOption"`
				ReqOption struct {
					Num1 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"1"`
					Num3 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"3"`
					Num6 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"6"`
					Num15 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"15"`
					Num28 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"28"`
					Num51 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"51"`
					Num58 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"58"`
					Num59 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"59"`
					Num90 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"90"`
					Num119 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"119"`
					Num120 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"120"`
					Num125 struct {
						Enable bool   `json:"Enable"`
						Alias  string `json:"Alias"`
						Tag    int    `json:"Tag"`
						Value  string `json:"Value"`
					} `json:"125"`
				} `json:"ReqOption"`
			} `json:"dhcp_data"`
		} `json:"dhcp"`
	} `json:"status"`
}

type NMC struct {
	Status struct {
		WanModeList           string `json:"WanModeList"`
		WanMode               string `json:"WanMode"`
		Username              string `json:"Username"`
		FactoryResetScheduled bool   `json:"FactoryResetScheduled"`
		ConnectionError       bool   `json:"ConnectionError"`
		DefaultsLoaded        bool   `json:"DefaultsLoaded"`
		ProvisioningState     string `json:"ProvisioningState"`
		OfferType             string `json:"OfferType"`
		OfferName             string `json:"OfferName"`
		IPTVMode              string `json:"IPTVMode"`
	} `json:"status"`
}

func main() {
	var ip, username string

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
	fmt.Print("Username: ")
	fmt.Scan(&username)

	// Ask for Password without displaying it
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln("Please provide a password")
	} else if len(password) == 0 {
		log.Fatalln("Password cannot be empty")
	}

	// Concatenate the URL
	url := "http://" + ip + "/ws"

	instantiateConnection(url, username, string(password))
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
		getOntInfos(instanciation.Data.ContextID, cookie, url)
		getMacAddress(instanciation.Data.ContextID, cookie, url)
		getInternetVlan(instanciation.Data.ContextID, cookie, url)
		getDHCPInfos(instanciation.Data.ContextID, cookie, url)
		// createOption90(instanciation.Data.ContextID, cookie, url)

		displayNecessaryInformations()
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

// Fetch ONT Informations
func getOntInfos(ContextID string, Cookie string, Url string) {

	payload := strings.NewReader("{\"service\":\"NeMo.Intf.veip0\",\"method\":\"getMIBs\",\"parameters\":{\"mibs\":\"gpon\"}}")

	req, _ := http.NewRequest("POST", Url, payload)

	req.Header.Add("cookie", Cookie)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", ContextID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var ont Ont

	if err := json.Unmarshal(body, &ont); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}

	GPON_SN = ont.Status.Gpon.Veip0.SerialNumber
	PON_VENDOR_ID = ont.Status.Gpon.Veip0.VendorID
	HW_HWVER = ont.Status.Gpon.Veip0.HardwareVersion
	OMCI_SW_VER1 = ont.Status.Gpon.Veip0.ONTSoftwareVersion0
	OMCI_SW_VER2 = ont.Status.Gpon.Veip0.ONTSoftwareVersion1
}

// Fetch Mac Address
func getMacAddress(ContextID string, Cookie string, Url string) {

	payload := strings.NewReader("{\"service\":\"NMC\",\"method\":\"getWANStatus\",\"parameters\":{}}")

	req, _ := http.NewRequest("POST", Url, payload)

	req.Header.Add("cookie", Cookie)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", ContextID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var mac MacAddress

	if err := json.Unmarshal(body, &mac); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}

	macaddress = mac.Data.MACAddress
}

// Fetch Internet VLAN
func getInternetVlan(ContextID string, Cookie string, Url string) {

	payload := strings.NewReader("{\"service\":\"NeMo.Intf.data\",\"method\":\"getFirstParameter\",\"parameters\":{\"name\":\"VLANID\"}}")

	req, _ := http.NewRequest("POST", Url, payload)

	req.Header.Add("cookie", Cookie)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", ContextID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var vlaninternet VlanInternet

	if err := json.Unmarshal(body, &vlaninternet); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}
	vlanid = strconv.Itoa(vlaninternet.Status)
}

// Fetch DHCP Infos
func getDHCPInfos(ContextID string, Cookie string, Url string) {

	payload := strings.NewReader("{\"service\":\"NeMo.Intf.data\",\"method\":\"getMIBs\",\"parameters\":{\"mibs\":\"dhcp\"}}")

	req, _ := http.NewRequest("POST", Url, payload)

	req.Header.Add("cookie", Cookie)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", ContextID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

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

// Automatically create the option 90 : https://jsfiddle.net/kgersen/3mnsc6wy/
func createOption90(ContextID string, Cookie string, Url string) {

	payload := strings.NewReader("{\"service\":\"NMC\",\"method\":\"get\",\"parameters\":{}}")

	req, _ := http.NewRequest("POST", Url, payload)

	req.Header.Add("cookie", Cookie)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", ContextID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var nmc NMC

	if err := json.Unmarshal(body, &nmc); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}

	const Salt = "1234567890123456"
	const Byte = "A"

	fmt.Print("Account FTI Password: ")
	ftipass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Please provide a password")
		os.Exit(1)
	}

	fmt.Println("")

	var h = md5.New()
	h.Write([]byte(Byte + string(ftipass) + Salt))

	md5 := hex.EncodeToString(h.Sum(nil))

	var md5s string

	for i := 0; i < len(md5); i += 2 {
		md5s += string(md5[i]) + string(md5[i+1])
		if i < len(md5)-2 {
			md5s = md5s + ":"
		}
	}

	var st11zero = "00:00:00:00:00:00:00:00:00:00:00"
	var idorange = "01" // variable
	var idsalt = "3c"   // 16
	var idhash = "03"   //1+16
	var fixed = "1a:09:00:00:05:58:01:03:41"

	output := st11zero + ":" + fixed + ":" + TLofTLS(idorange, len(nmc.Status.Username)+2) + ":" + SofTLS(nmc.Status.Username) + ":" + TLofTLS(idsalt, 2+16) + ":" + SofTLS(Salt) + ":" + TLofTLS(idhash, 2+1+16) + ":" + SofTLS(Byte) + ":" + md5s

	dhcpoption90 = output
}

func TLofTLS(id string, l int) string {
	var toAdd = strings.ToUpper((hex.EncodeToString([]byte(string(l)))))
	if len(toAdd) < 2 {
		toAdd = "0" + toAdd
	}
	return string(id + ":" + toAdd)
}

func SofTLS(s string) string {
	var toAdd string
	var res string = ""

	for i := 0; i < len(s); i++ {
		charcode := []rune(s)[i]
		toAdd = strings.ToUpper((hex.EncodeToString([]byte(string(charcode)))))
		if len(toAdd) < 2 {
			toAdd = "0" + toAdd
		}
		res += toAdd
		if i < len(s)-1 {
			res += ":"
		}
	}

	return res
}

func generateGponCommands() {
	fmt.Println("flash set GPON_PLOAM_PASSWD DEFAULT012")
	fmt.Println("flash set OMCI_TM_OPT 0")
	fmt.Println("flash set OMCI_OLT_MODE 1")
	fmt.Println("flash set GPON_SN " + GPON_SN)
	fmt.Println("flash set PON_VENDOR_ID " + PON_VENDOR_ID)
	fmt.Println("flash set HW_HWVER " + HW_HWVER)
	fmt.Println("flash set OMCI_SW_VER1 " + OMCI_SW_VER1)
	fmt.Println("flash set OMCI_SW_VER2 " + OMCI_SW_VER2)
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
