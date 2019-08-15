package znet

import (
	"bytes"
	"errors"
	"net"
)

func MacToHex(bb []byte) ([6]byte, error) {
	var tt [6]byte
	bb = bytes.ToLower(bb)
	if len(bb) != 12 {
		return tt, errors.New("source paramter's len wrong")
	}
	for k, v := range bb {
		if k%2 == 0 {
			if v >= 'a' {
				v = (v - 'a' + 10)
			} else {
				v = v - '0'
			}
			tt[5-k/2] = v * 16
		} else {
			if v >= 'a' {
				v = (v - 'a' + 10)
			} else {
				v = v - '0'
			}
			tt[5-k/2] += v
		}
	}
	return tt, nil
}
func Iprange(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	if len(ips) < 3 {
		return ips[0:len(ips)], nil
	}
	// remove broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
