package env

import (
	"net"
	"os"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-base/lib/beacon"
)

const BeaconMulticastIPAddressEnvName = "BEACON_IP_ADDRESS"

func beaconMulticastIPAddress() (net.IP, error) {
	var value net.IP = spacebeacon.MulticastIPAddress()

	{
		const envname = BeaconMulticastIPAddressEnvName

		envvalue := os.Getenv(envname)
		if "" != envvalue {
			parsedvalue := net.ParseIP(envvalue)
			if nil == parsedvalue {
				return value, erorr.Errorf("problem parsing value of %q environment variable (%q) as an IP-address", envname, envvalue)
			}

			value = parsedvalue
		}
	}

	if nil == value {
		return value, erorr.Error("nil IP-address")
	}
	return value, nil
}

var (
	_beaconMulticastIPAddress net.IP
	_beaconMulticastIPAddressError error
)

func init() {
	 _beaconMulticastIPAddress, _beaconMulticastIPAddressError = beaconMulticastIPAddress()
}

// BeaconMulticastIPAddress returns the multicast IP-address that the SPACE-BASE should use.
//
// It defaults to multicast IP-address 239.83.80.67 (0xEF535043).
//
// But, that can be overridden by a value set in the "BEACON_IP_ADDRESS" environment-variable.
func BeaconMulticastIPAddress() (net.IP, error) {
	return _beaconMulticastIPAddress, _beaconMulticastIPAddressError
}
