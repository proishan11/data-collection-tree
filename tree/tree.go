package tree

type tree struct {
	rootNode *Node
}

type Tree interface {
	Insert(newNode Node)
	FindByCountry(country string) *Node
	GetRoot() Node
}

func (t *tree) GetRoot() Node{
	return *t.rootNode
}

func(t *tree) Insert(newNode Node) {

	t.rootNode.mu.Lock()
	t.rootNode.WebRequests += newNode.WebRequests
	t.rootNode.TimeSpent += newNode.TimeSpent
	c, ok := t.rootNode.Children[newNode.Country]
	t.rootNode.mu.Unlock()

	if !ok {
		cNode := NewNode()
		cNode.TimeSpent = newNode.TimeSpent
		cNode.WebRequests = newNode.WebRequests
		cNode.Country = newNode.Country

		dNode := NewNode()
		dNode.TimeSpent = newNode.TimeSpent
		dNode.WebRequests = newNode.WebRequests
		dNode.Country = newNode.Country

		t.rootNode.mu.Lock()
		t.rootNode.Children[newNode.Country] = cNode
		t.rootNode.Children[newNode.Country].Children[newNode.Device] = dNode
		t.rootNode.mu.Unlock()
	} else {

		c.mu.Lock()
		c.WebRequests += newNode.WebRequests
		c.TimeSpent += newNode.TimeSpent
		//c.mu.Unlock()
		//
		//c.mu.Lock()
		d, ok := c.Children[newNode.Device]
		c.mu.Unlock()

		if ok {
			d.mu.Lock()
			d.TimeSpent += newNode.TimeSpent
			d.WebRequests += newNode.WebRequests
			d.mu.Unlock()
		} else {
			dNode := NewNode()
			dNode.TimeSpent = newNode.TimeSpent
			dNode.WebRequests = newNode.WebRequests
			dNode.Device = newNode.Device

			t.rootNode.Children[newNode.Country].mu.Lock()
			t.rootNode.Children[newNode.Country].Children[newNode.Device] = dNode
			t.rootNode.Children[newNode.Country].mu.Unlock()
		}
	}
}

func(t *tree) FindByCountry(country string) *Node {
	return t.rootNode.Children[country]
}

func NewTree() Tree {
	newNode := NewNode()

	return &tree {
		rootNode: newNode,
	}
}