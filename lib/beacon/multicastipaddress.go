package spacebeacon

import (
	"net"
)

// MulticastIPAddress returns the default multicast IP-address used by the SPACE-BEACON protocol.
// I.e., 239.83.80.67 (0xEF535043)
func MulticastIPAddress() net.IP {
	return net.IPv4(239, 'S', 'P', 'C')
}
