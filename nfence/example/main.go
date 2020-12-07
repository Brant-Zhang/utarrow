package main

import (
	"github.com/branthz/utarrow/nfence"
)

var etcds = []string{"http://92.168.29.100:12379", "http://92.168.29.100:22379", "http://92.168.29.100:32379"}

var nodeID = "1"
var ipaddr = "192.168.29.100"

func main() {
	n := nfence.NewNode(nodeID, ipaddr, etcds)
	n.Start()
}
