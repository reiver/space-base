package env

import (
	"net"
	"os"

	"github.com/reiver/go-erorr"
)

const DaemonIPAddressEnvName = "DAEMON_IP_ADDRESS"

func daemonIPAddressDefault() net.IP {
	addrs, err := net.InterfaceAddrs()
	if nil != err {
		return nil
	}

	for _, addr := range addrs {
		ipNet, casted := addr.(*net.IPNet)
		if !casted {
			continue
		}
		if ipNet.IP.IsLoopback() {
			continue
		}

		//@TODO: do we need to do this?
		if nil == ipNet.IP.To4() {
			continue
		}

		var ip net.IP = ipNet.IP
		if nil == ip {
			continue
		}

		return ip
	}

	return nil
}

func daemonIPAddress() (net.IP, error) {
	var value net.IP = daemonIPAddressDefault()

	{
		const envname = DaemonIPAddressEnvName

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
	_daemonIPAddress net.IP
	_daemonIPAddressError error
)

func init() {
	 _daemonIPAddress, _daemonIPAddressError = daemonIPAddress()
}

// DaemonIPAddress returns the multicast IP-address that the SPACE-BASE should use.
//
// It defaults to multicast IP-address 239.83.80.67 (0xEF535043).
//
// But, that can be overridden by a value set in the "DAEMON_IP_ADDRESS" environment-variable.
func DaemonIPAddress() (net.IP, error) {
	return _daemonIPAddress, _daemonIPAddressError
}
