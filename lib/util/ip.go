package util

import (
	"encoding/binary"
	"net"
)

func Ip2int(ips string) int {
	ip := net.ParseIP(ips)
	if len(ip) == 16 {
		return int(binary.BigEndian.Uint32(ip[12:16]))
	}
	return int(binary.BigEndian.Uint32(ip))
}

func Int2ip(nn int) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, uint32(nn))
	return ip.String()
}
