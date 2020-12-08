package nfence

type Node struct {
	ID     string
	Host   string
	etcds  []string
	status stateType
	zcli   *Zclient
}

func NewNode(id string, host string, etcds []string) *Node {
	n := &Node{
		ID:     id,
		Host:   host,
		etcds:  etcds,
		status: StateFollower,
		zcli:   NewZclient(etcds, id, host),
	}
	return n
}

func (n *Node) Start() {
	Infof("Node:%s start....", n.ID)
	n.zcli.run()
}

func (n *Node) Stop() {
	Infof("Node:%s stop....", n.ID)
	n.zcli.Stop()
}

func (n *Node) becomeLeader() {

}

func (n *Node) becomeFollower() {

}

func init() {
	Setup("", "DEBUG")
}
