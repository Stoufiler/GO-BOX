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

type FunboxValues struct {
	Result struct {
		Status struct {
			Base struct {
				Data struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
					} `json:"ULIntf"`
					LLIntf struct {
						Primdata struct {
							Name string `json:"Name"`
						} `json:"primdata"`
					} `json:"LLIntf"`
				} `json:"data"`
				Primdata struct {
					ULIntf struct {
						Data struct {
						} `json:"data"`
					} `json:"ULIntf"`
					LLIntf struct {
						Dhcpv6Data struct {
						} `json:"dhcpv6_data"`
						PppData struct {
						} `json:"ppp_data"`
					} `json:"LLIntf"`
				} `json:"primdata"`
				Dhcpv6Data struct {
					ULIntf struct {
						Primdata struct {
						} `json:"primdata"`
					} `json:"ULIntf"`
					LLIntf struct {
						PppData struct {
						} `json:"ppp_data"`
					} `json:"LLIntf"`
				} `json:"dhcpv6_data"`
				PppData struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
						Dslite struct {
							Name string `json:"Name"`
						} `json:"dslite"`
						Dhcpv6Data struct {
							Name string `json:"Name"`
						} `json:"dhcpv6_data"`
						Primdata struct {
							Name string `json:"Name"`
						} `json:"primdata"`
					} `json:"ULIntf"`
					LLIntf struct {
						GvlanData struct {
							Name string `json:"Name"`
						} `json:"gvlan_data"`
					} `json:"LLIntf"`
				} `json:"ppp_data"`
				GvlanData struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
						PppData struct {
							Name string `json:"Name"`
						} `json:"ppp_data"`
					} `json:"ULIntf"`
					LLIntf struct {
						Veip0 struct {
							Name string `json:"Name"`
						} `json:"veip0"`
					} `json:"LLIntf"`
				} `json:"gvlan_data"`
				Veip0 struct {
					Name   string `json:"Name"`
					Enable bool   `json:"Enable"`
					Status bool   `json:"Status"`
					Flags  string `json:"Flags"`
					ULIntf struct {
						GvlanData struct {
							Name string `json:"Name"`
						} `json:"gvlan_data"`
						GvlanIptv1 struct {
							Name string `json:"Name"`
						} `json:"gvlan_iptv1"`
						GvlanIptv2 struct {
							Name string `json:"Name"`
						} `json:"gvlan_iptv2"`
					} `json:"ULIntf"`
					LLIntf struct {
					} `json:"LLIntf"`
				} `json:"veip0"`
			} `json:"base"`
			SixRd struct {
			} `json:"6rd"`
			Alias struct {
				Data struct {
					Alias string `json:"Alias"`
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				Dhcpv6Data struct {
				} `json:"dhcpv6_data"`
				PppData struct {
					Alias string `json:"Alias"`
				} `json:"ppp_data"`
				GvlanData struct {
					Alias string `json:"Alias"`
				} `json:"gvlan_data"`
				Veip0 struct {
					Alias string `json:"Alias"`
				} `json:"veip0"`
			} `json:"alias"`
			Atm struct {
			} `json:"atm"`
			Bcmvlan struct {
				GvlanData struct {
					DscpToPbitsTable string `json:"DscpToPbitsTable"`
					QoS              struct {
						DefaultRx struct {
							Enable               bool `json:"Enable"`
							Tags                 int  `json:"Tags"`
							SetVlanID            int  `json:"SetVlanId"`
							SetPbits             int  `json:"SetPbits"`
							SetDscp              int  `json:"SetDscp"`
							SetEthertype         int  `json:"SetEthertype"`
							DscpToPbits          int  `json:"DscpToPbits"`
							PushTag              bool `json:"PushTag"`
							PopTag               bool `json:"PopTag"`
							Direction            bool `json:"Direction"`
							FilterVlanID         int  `json:"FilterVlanId"`
							FilterSkbPriority    int  `json:"FilterSkbPriority"`
							FilterSkbMarkFlowid  int  `json:"FilterSkbMarkFlowid"`
							FilterEthertype      int  `json:"FilterEthertype"`
							FilterVlanDevMacAddr int  `json:"FilterVlanDevMacAddr"`
						} `json:"default_rx"`
						ArpTx struct {
							Enable               bool `json:"Enable"`
							Tags                 int  `json:"Tags"`
							SetVlanID            int  `json:"SetVlanId"`
							SetPbits             int  `json:"SetPbits"`
							SetDscp              int  `json:"SetDscp"`
							SetEthertype         int  `json:"SetEthertype"`
							DscpToPbits          int  `json:"DscpToPbits"`
							PushTag              bool `json:"PushTag"`
							PopTag               bool `json:"PopTag"`
							Direction            bool `json:"Direction"`
							FilterVlanID         int  `json:"FilterVlanId"`
							FilterSkbPriority    int  `json:"FilterSkbPriority"`
							FilterSkbMarkFlowid  int  `json:"FilterSkbMarkFlowid"`
							FilterEthertype      int  `json:"FilterEthertype"`
							FilterVlanDevMacAddr int  `json:"FilterVlanDevMacAddr"`
						} `json:"arp_tx"`
						PppoedTx struct {
							Enable               bool `json:"Enable"`
							Tags                 int  `json:"Tags"`
							SetVlanID            int  `json:"SetVlanId"`
							SetPbits             int  `json:"SetPbits"`
							SetDscp              int  `json:"SetDscp"`
							SetEthertype         int  `json:"SetEthertype"`
							DscpToPbits          int  `json:"DscpToPbits"`
							PushTag              bool `json:"PushTag"`
							PopTag               bool `json:"PopTag"`
							Direction            bool `json:"Direction"`
							FilterVlanID         int  `json:"FilterVlanId"`
							FilterSkbPriority    int  `json:"FilterSkbPriority"`
							FilterSkbMarkFlowid  int  `json:"FilterSkbMarkFlowid"`
							FilterEthertype      int  `json:"FilterEthertype"`
							FilterVlanDevMacAddr int  `json:"FilterVlanDevMacAddr"`
						} `json:"pppoed_tx"`
						PppoesTx struct {
							Enable               bool `json:"Enable"`
							Tags                 int  `json:"Tags"`
							SetVlanID            int  `json:"SetVlanId"`
							SetPbits             int  `json:"SetPbits"`
							SetDscp              int  `json:"SetDscp"`
							SetEthertype         int  `json:"SetEthertype"`
							DscpToPbits          int  `json:"DscpToPbits"`
							PushTag              bool `json:"PushTag"`
							PopTag               bool `json:"PopTag"`
							Direction            bool `json:"Direction"`
							FilterVlanID         int  `json:"FilterVlanId"`
							FilterSkbPriority    int  `json:"FilterSkbPriority"`
							FilterSkbMarkFlowid  int  `json:"FilterSkbMarkFlowid"`
							FilterEthertype      int  `json:"FilterEthertype"`
							FilterVlanDevMacAddr int  `json:"FilterVlanDevMacAddr"`
						} `json:"pppoes_tx"`
						DefaultTx struct {
							Enable               bool `json:"Enable"`
							Tags                 int  `json:"Tags"`
							SetVlanID            int  `json:"SetVlanId"`
							SetPbits             int  `json:"SetPbits"`
							SetDscp              int  `json:"SetDscp"`
							SetEthertype         int  `json:"SetEthertype"`
							DscpToPbits          int  `json:"DscpToPbits"`
							PushTag              bool `json:"PushTag"`
							PopTag               bool `json:"PopTag"`
							Direction            bool `json:"Direction"`
							FilterVlanID         int  `json:"FilterVlanId"`
							FilterSkbPriority    int  `json:"FilterSkbPriority"`
							FilterSkbMarkFlowid  int  `json:"FilterSkbMarkFlowid"`
							FilterEthertype      int  `json:"FilterEthertype"`
							FilterVlanDevMacAddr int  `json:"FilterVlanDevMacAddr"`
						} `json:"default_tx"`
					} `json:"QoS"`
				} `json:"gvlan_data"`
			} `json:"bcmvlan"`
			Bridge struct {
			} `json:"bridge"`
			Copy struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				Dhcpv6Data struct {
				} `json:"dhcpv6_data"`
				PppData struct {
				} `json:"ppp_data"`
				GvlanData struct {
				} `json:"gvlan_data"`
				Veip0 struct {
				} `json:"veip0"`
			} `json:"copy"`
			DhcpAPI struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				Dhcpv6Data struct {
				} `json:"dhcpv6_data"`
				PppData struct {
				} `json:"ppp_data"`
				GvlanData struct {
				} `json:"gvlan_data"`
				Veip0 struct {
				} `json:"veip0"`
			} `json:"dhcp-api"`
			Dhcp struct {
			} `json:"dhcp"`
			Dhcpv6 struct {
				Dhcpv6Data struct {
					SentOption struct {
						Num15 struct {
						} `json:"15"`
						Num16 struct {
						} `json:"16"`
					} `json:"SentOption"`
					ReceivedOption struct {
					} `json:"ReceivedOption"`
				} `json:"dhcpv6_data"`
			} `json:"dhcpv6"`
			Dhcpv6Impl struct {
				Dhcpv6Data struct {
				} `json:"dhcpv6_data"`
			} `json:"dhcpv6impl"`
			DopSlave struct {
			} `json:"dop-slave"`
			Dsl struct {
			} `json:"dsl"`
			Dslbonding struct {
			} `json:"dslbonding"`
			Dslite struct {
			} `json:"dslite"`
			Dslline struct {
			} `json:"dslline"`
			Eth struct {
			} `json:"eth"`
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
				} `json:"veip0"`
			} `json:"gpon"`
			Gre struct {
			} `json:"gre"`
			Nat struct {
				PppData struct {
					NATEnabled bool `json:"NATEnabled"`
				} `json:"ppp_data"`
				GvlanData struct {
					NATEnabled bool `json:"NATEnabled"`
				} `json:"gvlan_data"`
				Veip0 struct {
					NATEnabled bool `json:"NATEnabled"`
				} `json:"veip0"`
			} `json:"nat"`
			NetdevAPI struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				Dhcpv6Data struct {
				} `json:"dhcpv6_data"`
				PppData struct {
				} `json:"ppp_data"`
				GvlanData struct {
				} `json:"gvlan_data"`
				Veip0 struct {
				} `json:"veip0"`
			} `json:"netdev-api"`
			Netdev struct {
				PppData struct {
					NetDevIndex                 int    `json:"NetDevIndex"`
					NetDevType                  string `json:"NetDevType"`
					NetDevFlags                 string `json:"NetDevFlags"`
					NetDevName                  string `json:"NetDevName"`
					LLAddress                   string `json:"LLAddress"`
					TxQueueLen                  int    `json:"TxQueueLen"`
					Mtu                         int    `json:"MTU"`
					NetDevState                 string `json:"NetDevState"`
					IPv4Forwarding              bool   `json:"IPv4Forwarding"`
					IPv4ForceIGMPVersion        int    `json:"IPv4ForceIGMPVersion"`
					IPv4AcceptSourceRoute       bool   `json:"IPv4AcceptSourceRoute"`
					IPv4AcceptLocal             bool   `json:"IPv4AcceptLocal"`
					IPv4AcceptRedirects         bool   `json:"IPv4AcceptRedirects"`
					IPv6AcceptRA                bool   `json:"IPv6AcceptRA"`
					IPv6ActAsRouter             bool   `json:"IPv6ActAsRouter"`
					IPv6AutoConf                bool   `json:"IPv6AutoConf"`
					IPv6MaxRtrSolicitations     int    `json:"IPv6MaxRtrSolicitations"`
					IPv6RtrSolicitationInterval int    `json:"IPv6RtrSolicitationInterval"`
					IPv6AcceptSourceRoute       bool   `json:"IPv6AcceptSourceRoute"`
					IPv6AcceptRedirects         bool   `json:"IPv6AcceptRedirects"`
					IPv6OptimisticDAD           bool   `json:"IPv6OptimisticDAD"`
					IPv6Disable                 bool   `json:"IPv6Disable"`
					IPv6HostPart                string `json:"IPv6HostPart"`
					RtTable                     int    `json:"RtTable"`
					IPv6AddrDelegate            string `json:"IPv6AddrDelegate"`
					IPv4Addr                    struct {
					} `json:"IPv4Addr"`
					IPv6Addr struct {
					} `json:"IPv6Addr"`
					IPv4Route struct {
					} `json:"IPv4Route"`
					IPv6Route struct {
					} `json:"IPv6Route"`
				} `json:"ppp_data"`
				GvlanData struct {
					NetDevIndex                 int    `json:"NetDevIndex"`
					NetDevType                  string `json:"NetDevType"`
					NetDevFlags                 string `json:"NetDevFlags"`
					NetDevName                  string `json:"NetDevName"`
					LLAddress                   string `json:"LLAddress"`
					TxQueueLen                  int    `json:"TxQueueLen"`
					Mtu                         int    `json:"MTU"`
					NetDevState                 string `json:"NetDevState"`
					IPv4Forwarding              bool   `json:"IPv4Forwarding"`
					IPv4ForceIGMPVersion        int    `json:"IPv4ForceIGMPVersion"`
					IPv4AcceptSourceRoute       bool   `json:"IPv4AcceptSourceRoute"`
					IPv4AcceptLocal             bool   `json:"IPv4AcceptLocal"`
					IPv4AcceptRedirects         bool   `json:"IPv4AcceptRedirects"`
					IPv6AcceptRA                bool   `json:"IPv6AcceptRA"`
					IPv6ActAsRouter             bool   `json:"IPv6ActAsRouter"`
					IPv6AutoConf                bool   `json:"IPv6AutoConf"`
					IPv6MaxRtrSolicitations     int    `json:"IPv6MaxRtrSolicitations"`
					IPv6RtrSolicitationInterval int    `json:"IPv6RtrSolicitationInterval"`
					IPv6AcceptSourceRoute       bool   `json:"IPv6AcceptSourceRoute"`
					IPv6AcceptRedirects         bool   `json:"IPv6AcceptRedirects"`
					IPv6OptimisticDAD           bool   `json:"IPv6OptimisticDAD"`
					IPv6Disable                 bool   `json:"IPv6Disable"`
					IPv6HostPart                string `json:"IPv6HostPart"`
					RtTable                     int    `json:"RtTable"`
					IPv6AddrDelegate            string `json:"IPv6AddrDelegate"`
					IPv4Addr                    struct {
					} `json:"IPv4Addr"`
					IPv6Addr struct {
					} `json:"IPv6Addr"`
					IPv4Route struct {
					} `json:"IPv4Route"`
					IPv6Route struct {
					} `json:"IPv6Route"`
				} `json:"gvlan_data"`
				Veip0 struct {
					NetDevIndex                 int    `json:"NetDevIndex"`
					NetDevType                  string `json:"NetDevType"`
					NetDevFlags                 string `json:"NetDevFlags"`
					NetDevName                  string `json:"NetDevName"`
					LLAddress                   string `json:"LLAddress"`
					TxQueueLen                  int    `json:"TxQueueLen"`
					Mtu                         int    `json:"MTU"`
					NetDevState                 string `json:"NetDevState"`
					IPv4Forwarding              bool   `json:"IPv4Forwarding"`
					IPv4ForceIGMPVersion        int    `json:"IPv4ForceIGMPVersion"`
					IPv4AcceptSourceRoute       bool   `json:"IPv4AcceptSourceRoute"`
					IPv4AcceptLocal             bool   `json:"IPv4AcceptLocal"`
					IPv4AcceptRedirects         bool   `json:"IPv4AcceptRedirects"`
					IPv6AcceptRA                bool   `json:"IPv6AcceptRA"`
					IPv6ActAsRouter             bool   `json:"IPv6ActAsRouter"`
					IPv6AutoConf                bool   `json:"IPv6AutoConf"`
					IPv6MaxRtrSolicitations     int    `json:"IPv6MaxRtrSolicitations"`
					IPv6RtrSolicitationInterval int    `json:"IPv6RtrSolicitationInterval"`
					IPv6AcceptSourceRoute       bool   `json:"IPv6AcceptSourceRoute"`
					IPv6AcceptRedirects         bool   `json:"IPv6AcceptRedirects"`
					IPv6OptimisticDAD           bool   `json:"IPv6OptimisticDAD"`
					IPv6Disable                 bool   `json:"IPv6Disable"`
					IPv6HostPart                string `json:"IPv6HostPart"`
					RtTable                     int    `json:"RtTable"`
					IPv6AddrDelegate            string `json:"IPv6AddrDelegate"`
					IPv4Addr                    struct {
					} `json:"IPv4Addr"`
					IPv6Addr struct {
					} `json:"IPv6Addr"`
					IPv4Route struct {
					} `json:"IPv4Route"`
					IPv6Route struct {
					} `json:"IPv6Route"`
				} `json:"veip0"`
			} `json:"netdev"`
			Penable struct {
			} `json:"penable"`
			Ppp struct {
				PppData struct {
					Username                        string `json:"Username"`
					ConnectionStatus                string `json:"ConnectionStatus"`
					LastConnectionError             string `json:"LastConnectionError"`
					MaxMRUSize                      int    `json:"MaxMRUSize"`
					PPPoESessionID                  int    `json:"PPPoESessionID"`
					PPPoEACName                     string `json:"PPPoEACName"`
					PPPoEServiceName                string `json:"PPPoEServiceName"`
					RemoteIPAddress                 string `json:"RemoteIPAddress"`
					LocalIPAddress                  string `json:"LocalIPAddress"`
					LastChangeTime                  int    `json:"LastChangeTime"`
					LastChange                      int    `json:"LastChange"`
					DNSServers                      string `json:"DNSServers"`
					TransportType                   string `json:"TransportType"`
					LCPEcho                         int    `json:"LCPEcho"`
					LCPEchoRetry                    int    `json:"LCPEchoRetry"`
					IPCPEnable                      bool   `json:"IPCPEnable"`
					IPv6CPEnable                    bool   `json:"IPv6CPEnable"`
					IPv6CPLocalInterfaceIdentifier  string `json:"IPv6CPLocalInterfaceIdentifier"`
					IPv6CPRemoteInterfaceIdentifier string `json:"IPv6CPRemoteInterfaceIdentifier"`
					ConnectionTrigger               string `json:"ConnectionTrigger"`
					IdleDisconnectTime              int    `json:"IdleDisconnectTime"`
				} `json:"ppp_data"`
			} `json:"ppp"`
			Ptm struct {
			} `json:"ptm"`
			RaAPI struct {
				Data struct {
				} `json:"data"`
				Primdata struct {
				} `json:"primdata"`
				Dhcpv6Data struct {
				} `json:"dhcpv6_data"`
				PppData struct {
				} `json:"ppp_data"`
				GvlanData struct {
				} `json:"gvlan_data"`
				Veip0 struct {
				} `json:"veip0"`
			} `json:"ra-api"`
			Ra struct {
				PppData struct {
					IPv6RouterDownTimeout int `json:"IPv6RouterDownTimeout"`
					IPv6Router            struct {
					} `json:"IPv6Router"`
				} `json:"ppp_data"`
			} `json:"ra"`
			Statmon struct {
			} `json:"statmon"`
			Switch struct {
			} `json:"switch"`
			Vlan struct {
				GvlanData struct {
					LastChangeTime int `json:"LastChangeTime"`
					LastChange     int `json:"LastChange"`
					Vlanid         int `json:"VLANID"`
					VLANPriority   int `json:"VLANPriority"`
				} `json:"gvlan_data"`
			} `json:"vlan"`
			Wan struct {
				Veip0 struct {
					PhysicalInterface string `json:"PhysicalInterface"`
				} `json:"veip0"`
			} `json:"wan"`
			Wlanbcmep struct {
			} `json:"wlanbcmep"`
			Wlanbcmrad struct {
			} `json:"wlanbcmrad"`
			Wlanbcmvap struct {
			} `json:"wlanbcmvap"`
			Wlanconfig struct {
			} `json:"wlanconfig"`
			Wlanendpoint struct {
			} `json:"wlanendpoint"`
			Wlanmtkrad struct {
			} `json:"wlanmtkrad"`
			Wlanquanrad struct {
			} `json:"wlanquanrad"`
			Wlanquanvap struct {
			} `json:"wlanquanvap"`
			Wlanradio struct {
			} `json:"wlanradio"`
			Wlanvap struct {
			} `json:"wlanvap"`
		} `json:"status"`
	} `json:"result"`
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
