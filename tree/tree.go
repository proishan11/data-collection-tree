package tree

import "sync"

type tree struct {
	rootNode *Node
}

type Tree interface {
	Insert(newNode Node)
	FindByCountry(country string) *Node
	GetRoot() Node
}

var (
	mu = sync.Mutex{}
)

func (t *tree) GetRoot() Node{
	return *t.rootNode
}

func(t *tree) Insert(newNode Node) {

	t.rootNode.mu.Lock()
	defer t.rootNode.mu.Unlock()
	t.rootNode.WebRequests += newNode.WebRequests
	t.rootNode.TimeSpent += newNode.TimeSpent

	_, ok := t.rootNode.Children.Load(newNode.Country)


	if !ok {
		cNode := NewNode()
		cNode.TimeSpent = newNode.TimeSpent
		cNode.WebRequests = newNode.WebRequests
		cNode.Country = newNode.Country

		dNode := NewNode()
		dNode.TimeSpent = newNode.TimeSpent
		dNode.WebRequests = newNode.WebRequests
		dNode.Country = newNode.Country

		cNode.Children.Store(newNode.Device, dNode)
		t.rootNode.Children.Store(newNode.Country, cNode)

	} else {

		temp, _ := t.rootNode.Children.Load(newNode.Country)
		c := temp.(*Node)
		c.mu.Lock()
		defer c.mu.Unlock()
		c.WebRequests += newNode.WebRequests
		c.TimeSpent += newNode.TimeSpent
		d, ok := c.Children.Load(newNode.Device)

		if ok {
			dNode := d.(*Node)
			dNode.mu.Lock()
			defer dNode.mu.Unlock()
			dNode.TimeSpent += newNode.TimeSpent
			dNode.WebRequests += newNode.WebRequests
		} else {
			dNode := NewNode()
			dNode.TimeSpent = newNode.TimeSpent
			dNode.WebRequests = newNode.WebRequests
			dNode.Device = newNode.Device
			c.Children.Store(newNode.Device, dNode)
		}
	}
}

func(t *tree) FindByCountry(country string) *Node {
	res, ok := t.rootNode.Children.Load(country)
	if ok {
		return res.(*Node)
	} else {
		return nil
	}
}

func NewTree() Tree {
	newNode := NewNode()

	return &tree {
		rootNode: newNode,
	}
}