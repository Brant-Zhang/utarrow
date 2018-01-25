package util

import(
	"testing"
)

func TestIpint(t *testing.T) {
     ip := "192.168.34.41"
     id := Ip2int(ip)
     ip2 := Int2ip(id)
     if ip != ip2 {
         t.Fatal("ip2int failed")
     }
     t.Logf("id:%d", id)
 }
