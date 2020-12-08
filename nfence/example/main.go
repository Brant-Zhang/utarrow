package main

import (
	"time"

	"github.com/branthz/utarrow/nfence"
)

var etcds = []string{"192.168.30.100:2379"}

var nodeID = "1"
var ipaddr = "192.168.29.100"

func ticker(n *nfence.Node) {
	tk := time.NewTicker(1e10)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			n.Stop()
			break
		}
	}
}

func main() {
	n := nfence.NewNode(nodeID, ipaddr, etcds)
	n.Start()
}
