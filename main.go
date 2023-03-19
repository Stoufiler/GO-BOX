package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

var leoxCommands []string
var fti string
var ftipass string

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

	fmt.Print("Livebox IP : ")
	fmt.Scan(&ip)

	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Please provide a real IP address")
		os.Exit(1)
	} else if !addr.IsPrivate() {
		fmt.Println("Please provide a local IP address")
		os.Exit(1)
	}

	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Please provide a password")
		os.Exit(1)
	}

	fmt.Print("Account FTI Password: ")
	// ftipass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	// if err != nil {
	// 	fmt.Println("Please provide a password")
	// 	os.Exit(1)
	// }

	url := "http://" + ip + "/ws"

	instantiateConnection(url, username, string(password))
}

func instantiateConnection(url string, username string, password string) {
	payload := strings.NewReader("{\"service\":\"sah.Device.Information\",\"method\":\"createContext\",\"parameters\":{\"applicationName\":\"webui\",\"username\":\"" + username + "\",\"password\":\"" + password + "\"}}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "X-Sah-Login")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

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
		fmt.Println("There was an error when connecting to Livebox")
		os.Exit(1)
	} else {
		fmt.Println("Successfully connected to Livebox !")
		displayNecessaryInformations(instanciation.Data.ContextID, cookie, url)
	}

}

func displayNecessaryInformations(ContextID string, Cookie string, Url string) {
	fmt.Println("=========================================")
	fmt.Println(getOntInfos(ContextID, Cookie, Url))
	fmt.Println(getMacAddress(ContextID, Cookie, Url))
	fmt.Println(getInternetVlan(ContextID, Cookie, Url))
	fmt.Println(getDHCPInfos(ContextID, Cookie, Url))
	fmt.Println(createOption90(ContextID, Cookie, Url))
}

func getOntInfos(ContextID string, Cookie string, Url string) string {

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

	return string("Serial Number       : " + ont.Status.Gpon.Veip0.SerialNumber + "\n" +
		"Hardware Version    : " + ont.Status.Gpon.Veip0.HardwareVersion + "\n" +
		"Vendor ID           : " + ont.Status.Gpon.Veip0.VendorID + "\n" +
		"ONTSoftwareVersion0 : " + ont.Status.Gpon.Veip0.ONTSoftwareVersion0 + "\n" +
		"ONTSoftwareVersion1 : " + ont.Status.Gpon.Veip0.ONTSoftwareVersion1,
	)
}

func getMacAddress(ContextID string, Cookie string, Url string) string {

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

	return string("Mac Address         : " + mac.Data.MACAddress)
}

func getInternetVlan(ContextID string, Cookie string, Url string) string {

	payload := strings.NewReader("{\"service\":\"NeMo.Intf.data\",\"method\":\"getFirstParameter\",\"parameters\":{\"name\":\"VLANID\"}}")

	req, _ := http.NewRequest("POST", Url, payload)

	req.Header.Add("cookie", Cookie)
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", ContextID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var vlanid VlanInternet

	if err := json.Unmarshal(body, &vlanid); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}

	return string("VLAN ID             : " + strconv.Itoa(vlanid.Status))
}

func getDHCPInfos(ContextID string, Cookie string, Url string) string {

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

	return string("DHCP Option 77      : " + opt70[1])
}

func createOption90(ContextID string, Cookie string, Url string) string {

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
	fmt.Println(ftipass)
	const Salt = "123456667890123456"
	const Byte = "A"

	var h = md5.New()
	io.WriteString(h, Byte+ftipass+Salt)
	md5 := h.Sum(nil)

	fmt.Println(len(md5))

	var md5s byte

	for i := 0; i < len(md5); i += 2 {
		md5s += md5[i] + md5[i+1]
		// if i < len(md5)-2 {
		// 	md5s += ":"
		// }
	}

	fmt.Println(md5s)

	return string("DHCP Option 90      : " + nmc.Status.Username)
}
