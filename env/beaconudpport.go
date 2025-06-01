package env

import (
	"os"
	"strconv"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-base/lib/beacon"
)

const BeaconUDPPortEnvName = "BEACON_UDP_PORT"

func beaconUDPPort() (uint16, error) {
	var value uint16 = spacebeacon.UDPPort

	{
		const envname = BeaconUDPPortEnvName

		envvalue := os.Getenv(envname)
		if "" != envvalue {
			const base = 10
			const bitSize = 16
			u64, err := strconv.ParseUint(envvalue, base, bitSize)
			if nil != err {
				return value, erorr.Errorf("problem parsing value of %q environment variable: %w", envname, err)
			}

			value = uint16(u64)
		}
	}

	return value, nil
}

var (
	_beaconUDPPort uint16
	_beaconUDPPortError error
)

func init() {
	 _beaconUDPPort, _beaconUDPPortError = beaconUDPPort()
}

// BeaconUDPPort returns the UDP-port that the HTTP server should use.
//
// It defaults to UDP-port 21328 (0x5350).
//
// But, that can be overridden by a value set in the "BEACON_UDP_PORT" environment-variable.
func BeaconUDPPort() (uint16, error) {
	return _beaconUDPPort, _beaconUDPPortError
}
