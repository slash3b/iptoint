package iptoint

import (
	"net"
)

func Ipv4ToInt(ipString string) int {
	ip := net.ParseIP(ipString)

	res := ip.To4()
	var output int

	output = int(res[0]) << 24
	output += int(res[1]) << 16
	output += int(res[2]) << 8
	output += int(res[3])

	return output
}
