package nfence

type node struct {
	ID     string
	Host   string
	etcds  []string
	status stateType
	zcli   *Zclient
}

func NewNode(id string, host string, etcds []string) *node {
	n := &node{
		ID:     id,
		Host:   host,
		etcds:  etcds,
		status: StateFollower,
		zcli:   NewZclient(etcds, id, host),
	}
	return n
}

func (n *node) Start() {
	n.zcli.run()
}

func (n *node) becomeLeader() {

}

func (n *node) becomeFollower() {

}

func init() {
	Setup("", "DEBUG")
}
