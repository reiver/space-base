package main

import (
	"net"
	"os"

	"github.com/reiver/space-base/env"
	"github.com/reiver/space-base/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-base ⚡")

	var spaceBeaconMulticastIPAddress net.IP
	{
		const envname = env.BeaconMulticastIPAddressEnvName
		log.Informf("Note that, the SPACE-BEACON multicast IP-address can be set by using the %q environment-variable.", envname)

		value, err := env.BeaconMulticastIPAddress()
		if nil != err {
			log.Errorf("ERROR: value of %q environment-variable invalid: %s", envname, err)
			os.Exit(1)
			return
		}

		spaceBeaconMulticastIPAddress = value
	}
	log.Informf("SPACE-BEACON multicast ip-address: %v", spaceBeaconMulticastIPAddress)

	var spaceBeaconUDPPort uint16
	{
		const envname = env.BeaconUDPPortEnvName
		log.Informf("Note that, the SPACE-BEACON UDP-port number can be set by using the %q environment-variable.", envname)

		value, err := env.BeaconUDPPort()
		if nil != err {
			log.Errorf("ERROR: value of %q environment-variable invalid: %s", envname, err)
			os.Exit(1)
			return
		}

		spaceBeaconUDPPort = value
	}
	log.Informf("SPACE-BEACON UDP port: %v (0x%X)", spaceBeaconUDPPort, spaceBeaconUDPPort)

	beacon(spaceBeaconMulticastIPAddress, spaceBeaconUDPPort)
}
