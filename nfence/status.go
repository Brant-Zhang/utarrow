package nfence

const (
	StateLeader stateType = iota
	StateFollower
)

type stateType int
