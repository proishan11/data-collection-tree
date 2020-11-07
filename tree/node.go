package tree

import "sync"

type Node struct {
	Children sync.Map
	Country string
	Device string

	mu sync.Mutex
	WebRequests int32
	TimeSpent int32
}

func NewNode() *Node {
	return &Node{
		WebRequests: 0,
		TimeSpent: 0,
		mu: sync.Mutex{},
	}
}