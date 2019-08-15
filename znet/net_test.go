package znet

import (
	"testing"
)

func TestIprange(t *testing.T) {
	ips, err := Iprange("192.168.29.154/24")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ips)
}
