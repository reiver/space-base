package main

import (
	"fmt"
	"net"
	"os"

	"github.com/reiver/space-base/env"
)

func main() {
	fmt.Println("space-base ⚡")

	var spaceBeaconMulticastIPAddress net.IP
	{
		const envname = env.BeaconMulticastIPAddressEnvName
		fmt.Printf("Note that, the SPACE-BEACON multicast IP-address can be set by using the %q environment-variable.\n", envname)

		value, err := env.BeaconMulticastIPAddress()
		if nil != err {
			fmt.Printf("ERROR: value of %q environment-variable invalid: %s\n", envname, err)
			os.Exit(1)
			return
		}

		spaceBeaconMulticastIPAddress = value
	}
	fmt.Printf("SPACE-BEACON multicast ip-address: %v\n", spaceBeaconMulticastIPAddress)

	var spaceBeaconUDPPort uint16
	{
		const envname = env.BeaconUDPPortEnvName
		fmt.Printf("Note that, the SPACE-BEACON UDP-port number can be set by using the %q environment-variable.\n", envname)

		value, err := env.BeaconUDPPort()
		if nil != err {
			fmt.Printf("ERROR: value of %q environment-variable invalid: %s\n", envname, err)
			os.Exit(1)
			return
		}

		spaceBeaconUDPPort = value
	}
	fmt.Printf("SPACE-BEACON UDP port: %v (0x%X)\n", spaceBeaconUDPPort, spaceBeaconUDPPort)

	beacon(spaceBeaconMulticastIPAddress, spaceBeaconUDPPort)
}
