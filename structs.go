package main

// Define struct to handle http requests : https://mholt.github.io/json-to-go/
type Context struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	} `json:"data"`
}

type FunboxContext struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
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
