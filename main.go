package main

import (
	"fmt"
	"github.com/proishan11/data-collection-tree/tree"
	"unsafe"
)

func main() {
	fmt.Println("Data collection tree")

	t := tree.NewTree()
	node := tree.NewNode()

	node.Country = "US"
	node.Device = "Mobile"
	node.WebRequests = 10
	node.TimeSpent = 100

	n2 := tree.NewNode()
	n2.Country = "US"
	n2.WebRequests = 20
	n2.TimeSpent = 90
	n2.Device = "Laptop"

	n3 := tree.NewNode()
	n3.Country = "US"
	n3.WebRequests = 20
	n3.TimeSpent = 90
	n3.Device = "Laptop"

	t.Insert(node)
	t.Insert(n2)
	t.Insert(n3)

	fmt.Println("Insertion Complete")
	fmt.Println(unsafe.Sizeof(t))

}
