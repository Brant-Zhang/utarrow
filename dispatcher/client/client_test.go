package client

import (
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	etcd := []string{"192.168.29.100:22379", "192.168.29.100:32379", "192.168.29.100:2379"}
	n := NewNode(2, "controller/part2/seed", "11.11.11.22", etcd)
	go n.Start()
	nb := NewNode(2, "controller/part2/back", "11.11.11.33", etcd)
	go nb.Start()
	time.Sleep(1e9)
	n.gettest("controller/part2/")
	select {}
}
