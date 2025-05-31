package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("space-base ⚡")

	// 239.83.80.67 (0xEF535043)
	var multicastIPAddress net.IP = net.IPv4(239, 'S', 'P', 'C')
	fmt.Printf("multicast ip-address: %v\n", multicastIPAddress)

	// 21328 (0x5350)
	var udpPort uint16 = (uint16('S') << 8) | uint16('P')
	fmt.Printf("UDP port: %v (0x%X)\n", udpPort, udpPort)

	var multicastUDPAddress = net.UDPAddr{
		IP: multicastIPAddress,
		Port: int(udpPort),
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
		var buffer [508]byte

		const limit = 10
		const sleepDuration = 5 * time.Second


		for i:=0; i<limit; i++ {

			var msg []byte = buffer[0:0]

			msg = append(msg, "SPACE/1.0\nDOROOD\n\n"...)

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
