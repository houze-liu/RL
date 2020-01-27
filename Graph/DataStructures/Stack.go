package DataStructures

import "fmt"

type Stack struct {
	root *Node
}

func (s *Stack) PrintAll () {
	var cur *Node
	cur = s.root
	if cur == nil {
		fmt.Println("Empty Stack")
		return
	}
	for cur.next != nil {
		fmt.Print(cur.value)
		fmt.Print("\n")
		cur = cur.next
	}
	fmt.Print(cur.value)
	fmt.Print("\n")
}

func (s *Stack) IsEmpty() bool {
	if s.root == nil{return true} else {return false}
}

func (s *Stack) Push(e interface{}) {
	if s.root == nil {s.root = &Node{e, nil}} else
	{
		s.root.add(e)
	}
}

func (s *Stack) Pull() interface{} {
	if s.IsEmpty() {fmt.Println("pull from empty stack!"); return nil}
	var prev *Node
	var cur *Node
	cur = s.root
	// 找最后一个元素推出来，并解除链接
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}

	if prev != nil {prev.next=nil} else {s.root = nil}

	return cur.value
}
