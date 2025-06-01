package env

import (
	"os"
	"strconv"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-base/lib/beacon"
)

const DaemonTCPPortEnvName = "DAEMON_TCP_PORT"

func daemonTCPPort() (uint16, error) {
	var value uint16 = spacebeacon.TCPPort

	{
		const envname = DaemonTCPPortEnvName

		envvalue := os.Getenv(envname)
		if "" != envvalue {
			const base = 10
			const bitSize = 16
			u64, err := strconv.ParseUint(envvalue, base, bitSize)
			if nil != err {
				return value, erorr.Errorf("problem parsing value of %q environment variable (%q) as a %d-bit base-%d unsigned-integer: %w", envname, envvalue, bitSize, base, err)
			}

			value = uint16(u64)
		}
	}

	return value, nil
}

var (
	_daemonTCPPort uint16
	_daemonTCPPortError error
)

func init() {
	 _daemonTCPPort, _daemonTCPPortError = daemonTCPPort()
}

// DaemonTCPPort returns the UDP-port that the SPACE-BASE should use.
//
// It defaults to UDP-port 21328 (0x5350).
//
// But, that can be overridden by a value set in the "DAEMON_TCP_PORT" environment-variable.
func DaemonTCPPort() (uint16, error) {
	return _daemonTCPPort, _daemonTCPPortError
}
