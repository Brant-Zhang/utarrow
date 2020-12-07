package nfence

import (
	"testing"
)

var etcds = []string{"http://92.168.29.100:12379", "http://92.168.29.100:22379", "http://92.168.29.100:32379"}

func TestNodeRun(t *testing.T) {
	var id = "1"
	var host = "192.168.30.100"
	n := newNode(id, host, etcds)
	n.Start()
}
