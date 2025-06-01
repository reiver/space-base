package main

import (
	"net"
	"net/http"
	"os"

	"github.com/reiver/space-base/env"
	"github.com/reiver/space-base/srv/log"
)

func _daemon(tcpAddress string) {
	go func() {
		log := logsrv.Prefix("_daemon").Begin()
		defer log.End()

		err := http.ListenAndServe(tcpAddress, nil)
		if nil != err {
			log.Errorf("problem HTTP listening-and-serving: %s", err)
		}
	}()
}

func daemon() string {
	log := logsrv.Prefix("daemon").Begin()
	defer log.End()

	var daemonIPAddress net.IP
	{
		const envname = env.DaemonIPAddressEnvName
		log.Informf("Note that, the WWW-DAEMON IP-address can be set by using the %q environment-variable.", envname)

		value, err := env.DaemonIPAddress()
		if nil != err {
			log.Errorf("ERROR: value of %q environment-variable invalid: %s", envname, err)
			os.Exit(1)
			return ""
		}

		daemonIPAddress = value
	}
	log.Informf("WWW-DAEMON IP-address: %v", daemonIPAddress)

	var daemonTCPPort uint16
	{
		const envname = env.DaemonTCPPortEnvName
		log.Informf("Note that, the WWW-DAEMON TCP-port can be set by using the %q environment-variable.", envname)

		value, err := env.DaemonTCPPort()
		if nil != err {
			log.Errorf("ERROR: value of %q environment-variable invalid: %s", envname, err)
			os.Exit(1)
			return ""
		}

		daemonTCPPort = value
	}
	log.Informf("WWW-DAEMON TCP-PORT: %v", daemonTCPPort)


	var daemonTCPAddress string = (&net.TCPAddr{IP:daemonIPAddress,Port:int(daemonTCPPort)}).String()
	log.Informf("WWW-DAEMON TCP-address: %v", daemonTCPAddress)

	_daemon(daemonTCPAddress)

	return daemonTCPAddress
}

