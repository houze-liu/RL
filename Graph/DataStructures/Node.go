package DataStructures

import "fmt"

type Node struct {
	value interface{}
	next *Node
}

func (root *Node) add (i interface{}) {
	// 'root' is a copy pointer of the passed-in pointer pointing to a Node
	for root.next != nil {
		root = root.next
	}
	// 当while 结束时，root.next 一定为nil
	root.next = &Node{i, nil}
}

func (root *Node) printAll () {
	for root.next != nil {
		fmt.Print(root.value)
		fmt.Print("\n")
		root = root.next
	}
	fmt.Print(root.value)
	fmt.Print("\n")
}
