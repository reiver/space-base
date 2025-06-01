package spacebeacon

// TCPPort is the default TCP port used by the SPACE-BEACON protocol's HTTP server.
// I.e., 21328 (0x5350)
//
// Numerically, it is the same value as [UDPPort].
const TCPPort uint16 = UDPPort
