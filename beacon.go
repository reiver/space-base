package main

import (
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/reiver/space-base/lib/beacon"
	"github.com/reiver/space-base/srv/log"
)

func beacon(ipAddress net.IP, udpPort uint16) {

	log := logsrv.Prefix("beacon").Begin()
	defer log.End()

	var multicastUDPAddress = net.UDPAddr{
		IP: ipAddress,
		Port: int(udpPort),
	}
	log.Informf("SPACE-BEACON UDP address: %v", &multicastUDPAddress)

	udpConn, err := net.DialUDP("udp", nil, &multicastUDPAddress)
	if nil != err {
		log.Errorf("ERROR: could not successfully dial UDP address %v: %s", &multicastUDPAddress, err)
		os.Exit(1)
		return
	}
	defer udpConn.Close()
	log.Inform("Connected!")

	localAddr := udpConn.LocalAddr()
	if nil == localAddr {
		log.Errorf("ERROR: could not get UDP local-address: %s", err)
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
				log.Errorf("ERROR: problem writing message #%d: %s", 1+i, err)
				continue
			}
			log.Debugf("Wrote message #%d", 1+i)

			var sleepDuration time.Duration = baseSleepDuration + (time.Millisecond * time.Duration(rand.Int63n(3555)))
			log.Debugf("Will sleep for %v", sleepDuration)

			time.Sleep(sleepDuration)
		}
	}
}
