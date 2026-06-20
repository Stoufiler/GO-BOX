package api

// PPP holds the PPPoE credentials returned by "NMC get". Unlike the "ppp" MIB
// (which strips the password), NMC exposes both Username and Password.
type PPP struct {
	Status struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	} `json:"status"`
}

type Ont struct {
	Status struct {
		Gpon struct {
			Veip0 struct {
				SerialNumber        string `json:"SerialNumber"`
				VendorID            string `json:"VendorId"`
				HardwareVersion     string `json:"HardwareVersion"`
				ONTSoftwareVersion0 string `json:"ONTSoftwareVersion0"`
				ONTSoftwareVersion1 string `json:"ONTSoftwareVersion1"`
			} `json:"veip0"`
		} `json:"gpon"`
	} `json:"status"`
}

type MacAddress struct {
	Data struct {
		MACAddress string `json:"MACAddress"`
	} `json:"data"`
}

type VlanInternet struct {
	Status int `json:"status"`
}

type DHCP struct {
	Status struct {
		Dhcp struct {
			DhcpData struct {
				SentOption struct {
					Num60 struct {
						Value string `json:"Value"`
					} `json:"60"`
					Num77 struct {
						Value string `json:"Value"`
					} `json:"77"`
					Num90 struct {
						Value string `json:"Value"`
					} `json:"90"`
				} `json:"SentOption"`
			} `json:"dhcp_data"`
		} `json:"dhcp"`
	} `json:"status"`
}

type OltVendor struct {
	Status struct {
		Classp []struct {
			Meid       int    `json:"MEId"`
			Attributes string `json:"Attributes"`
		} `json:"classp"`
	} `json:"status"`
}

// DHCPInfo holds the DHCP options.
type DHCPInfo struct {
	Option60 string
	Option77 string
	Option90 string
}
