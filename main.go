package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/reiver/space-base/env"
	"github.com/reiver/space-base/lib/beacon"
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

	var multicastUDPAddress = net.UDPAddr{
		IP: spaceBeaconMulticastIPAddress,
		Port: int(spaceBeaconUDPPort),
	}
	fmt.Printf("SPACE-BEACON UDP address: %v\n", &multicastUDPAddress)

	udpConn, err := net.DialUDP("udp", nil, &multicastUDPAddress)
	if nil != err {
		fmt.Printf("ERROR: could not successfully dial UDP address %v: %s", &multicastUDPAddress, err)
		os.Exit(1)
		return
	}
	defer udpConn.Close()
	fmt.Println("Connected!")

	localAddr := udpConn.LocalAddr()
	if nil == localAddr {
		fmt.Printf("ERROR: could not get UDP local-address: %s\n", err)
		os.Exit(1)
		return
	}

	{
		var buffer [spacebeacon.MaxLength]byte

		const limit = 10
		const sleepDuration = 5 * time.Second


		for i:=0; i<limit; i++ {

			var msg []byte = buffer[0:0]

			msg = append(msg, spacebeacon.Magic...)
			msg = append(msg, "DOROOD\n\n"...)

			_, err := udpConn.Write(msg)
			if nil != err {
				fmt.Printf("ERROR: problem writing message #%d: %s\n", 1+i, err)
				continue
			}
			fmt.Printf("Wrote message #%d\n", 1+i)

			time.Sleep(sleepDuration)
		}
	}
}
