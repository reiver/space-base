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
	blur()

	beacon()

	var daemonIPAddress net.IP
	{
		const envname = env.DaemonIPAddressEnvName
		log.Informf("Note that, the WWW-DAEMON IP-address can be set by using the %q environment-variable.", envname)

		value, err := env.DaemonIPAddress()
		if nil != err {
			log.Errorf("ERROR: value of %q environment-variable invalid: %s", envname, err)
			os.Exit(1)
			return
		}

		daemonIPAddress = value
	}
	log.Informf("WWW-DAEMON ip-address: %v", daemonIPAddress)

}
