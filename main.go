package main

import (
	"fmt"
	"net"
	"time"

	"github.com/reiver/space-base/lib/beacon"
)

func main() {
	fmt.Println("space-base ⚡")

	var multicastIPAddress net.IP = spacebeacon.MulticastIPAddress()
	fmt.Printf("multicast ip-address: %v\n", multicastIPAddress)

	fmt.Printf("UDP port: %v (0x%X)\n", spacebeacon.UDPPort, spacebeacon.UDPPort)

	var multicastUDPAddress = net.UDPAddr{
		IP: multicastIPAddress,
		Port: int(spacebeacon.UDPPort),
	}
	fmt.Printf("UDP address: %v\n", &multicastUDPAddress)

	udpConn, err := net.DialUDP("udp", nil, &multicastUDPAddress)
	if nil != err {
		fmt.Printf("ERROR: could not successfully dial UDP address %v: %s", &multicastUDPAddress, err)
		return
	}
	defer udpConn.Close()
	fmt.Println("Connected!")

	localAddr := udpConn.LocalAddr()
	if nil == localAddr {
		fmt.Printf("ERROR: could not get UDP local-address: %s\n", err)
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
