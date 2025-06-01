package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/reiver/space-base/lib/beacon"
)

func beacon(ipAddress net.IP, udpPort uint16) {

	var multicastUDPAddress = net.UDPAddr{
		IP: ipAddress,
		Port: int(udpPort),
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
		const baseSleepDuration = 2 * time.Second

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

			var sleepDuration time.Duration = baseSleepDuration + (time.Millisecond * time.Duration(rand.Int63n(3555)))
			fmt.Printf("Will sleep for %v\n", sleepDuration)

			time.Sleep(sleepDuration)
		}
	}
}
