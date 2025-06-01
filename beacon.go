package main

import (
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/reiver/space-base/env"
	"github.com/reiver/space-base/lib/beacon"
	"github.com/reiver/space-base/srv/log"
)

func beacon(wwwDaemonTCPAddress string) {

	log := logsrv.Prefix("beacon").Begin()
	defer log.End()

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
	log.Informf("SPACE-BEACON multicast IP-address: %v", spaceBeaconMulticastIPAddress)

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

	_beacon(spaceBeaconMulticastIPAddress, spaceBeaconUDPPort, wwwDaemonTCPAddress)
}

func _beacon(mutlicastIPAddress net.IP, udpPort uint16, wwwDaemonTCPAddress string) {

	log := logsrv.Prefix("_beacon").Begin()
	defer log.End()

	var multicastUDPAddress = net.UDPAddr{
		IP: mutlicastIPAddress,
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
			msg = append(msg, "DOROOD\n"...)
			msg = append(msg, wwwDaemonTCPAddress...)
			msg = append(msg, '\n')
			msg = append(msg, '\n')

			log.Debug("Sending message to SPACE-COMMAND…")
			_, err := udpConn.Write(msg)
			if nil != err {
				log.Errorf("ERROR: problem writing message #%d: %s", 1+i, err)
				continue
			}
			log.Debugf("Message #%d sent!", 1+i)

			var sleepDuration time.Duration = baseSleepDuration + (time.Millisecond * time.Duration(rand.Int63n(3555)))
			log.Debugf("Will sleep for %v", sleepDuration)

			time.Sleep(sleepDuration)
		}
	}
}
