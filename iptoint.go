package iptoint

import (
	"net"
)

func IpToInt(ip net.IP) int {
	if ipv4 := ip.To4(); ipv4 != nil {
		return int(ipv4[0])<<24 | int(ipv4[1])<<16 | int(ipv4[2])<<8 | int(ipv4[3])
	}

	return 0
}
