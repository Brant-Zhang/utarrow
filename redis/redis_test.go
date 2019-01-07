package redis

import (
	"testing"

	rd "github.com/gomodule/redigo/redis"
)

func init() {
	Init("127.0.0.1:6379")
}

func TestBasic(t *testing.T) {
	var keystr = "hangzhou"
	var valuestr = "zhejiang"

	_, err := Client.Set(keystr, valuestr)
	if err != nil {
		t.Fatal(err)
	}
	prov, err := rd.String(Client.Get(keystr))
	if err != nil {
		t.Fatal(err)
	}
	if prov != valuestr {
		t.Fatal("get diffrent from set")
	}

	var cid = "11"
	var sip = "192.168.1.1"
	var mac = "11:22:33:44:55:66"
	_, err = Client.Hset(cid, sip, mac)
	if err != nil {
		t.Fatal(err)
	}
	var sipB = "192.168.1.2"
	var macB = "aa:bb:cc"
	_, err = Client.Hset(cid, sipB, macB)
	if err != nil {
		t.Fatal(err)
	}
	val, err := rd.String(Client.Hget(cid, sip))
	if err != nil {
		t.Fatal(err)
	}
	if val != mac {
		t.Fatal("hget:get diffrent from set")
	}
	var sipC = "192.168.1.3"
	ret, err := Client.Hmget(cid, []string{sip, sipB, sipC})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hmget:%v", ret)
}
