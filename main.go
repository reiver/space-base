package main

import (
	"github.com/reiver/space-base/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-base ⚡")
	blur()

	log.Inform("A daemon will be spawned…")
	daemon()

	log.Inform("I will let SPACE-COMMAND know we are here…")
	beacon()
}
