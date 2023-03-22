<div align="center">
<img src="images/logo.png" width="50%">
<h1> GO-BOX</h1>
<small><i>Tool to help you to bypass your Livebox</i></small>
</div>


# Overview
GO-BOX is an application written in GO that allows you to automatically generate commands to help you to bypass your Livebox

This idea came to me while looking for how to replace my Livebox with my UDM SE.

I've spent so much time to find how to configure the SFP GPON and DHCP Options, so to learn GO and share to community I decided to create this simple app

# Setup
## Use binaries
Download latest release according to your OS

```bash
tar -xvcf gobox_Linux_x86_64.tar.gz
cd gobox_Linux_x86_64
./gobox
```

## Compile by yourself
```bash
git clone https://github.com/StephanGR/gobox.git
cd gobox
go build .
./gobox
```

# Usage
Simply fill fields asked by program

Example :
```bash
./gobox
Livebox IP : 192.168.1.1
Username : admin
Password : 
✅ Successfully connected to Livebox ! ✅

===========LEOX GPON COMMAND=============
flash set GPON_PLOAM_PASSWD DEFAULT012
flash set OMCI_TM_OPT 0
flash set OMCI_OLT_MODE 1
flash set GPON_SN XXXXXXXXXXXX
flash set PON_VENDOR_ID SMBS
flash set HW_HWVER XXXXXXXXXXXX
flash set OMCI_SW_VER1 XXXXXXXXXXXXX
flash set OMCI_SW_VER2 XXXXXXXXXXXXX

## Unplug fiber from Livebox and plug it into UDM and wait a minute ##

Execute this command -> omcicli mib get 131

OLT VENDOR ID (HWTC, ALCL, ...) : alcl

Execute this command -> flash set OMCC_VER 128

=========================================

==========UDM PRO SE SETTINGS============
NAME              : LEOX GPON
VLAN ID           : 832
MAC Address Clone : XX:XX:XX:XX:XX:XX
DHCP OPTION 60    : sagem
DHCP OPTION 77    : FSVDSL_livebox.Internet.softathome.Livebox5
DHCP OPTION 90    : 00:00:00:00:00:00:00:00:00:00:00:1a:09:00:00:05:58:01:03:41:01:0D:66:74:69:2F:67:66:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX
DHCP CoS          : 6 
=========================================
```

## Special thanks
I would never have been able to have all this information without the help of the [lafibre.info](https://lafibre.info/remplacer-livebox/mise-en-route-leox-lxt-010h-d/) forum as well as [iMordo](https://github.com/iMord0)

Thanks to them !