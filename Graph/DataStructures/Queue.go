package DataStructures

import "fmt"

type Queue struct {
	root *Node
}

func (q *Queue) PrintAll () {
	var cur *Node
	cur = q.root
	if cur == nil {
		fmt.Println("Empty Queue")
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

func (q *Queue) IsEmpty () bool {
	if q.root == nil{return true} else {return false}
}

func (q *Queue) Pull () interface{} {
	var r interface{}
	if q.root == nil {
		fmt.Println("Empty Queue")
		return nil
	} else
	if q.root.next == nil {
		r = q.root.value
		q.root = nil
	} else {
		r = q.root.value
		q.root = q.root.next
	}

	return r
}

func (q *Queue) Push (e interface{}) {
	if q.root == nil {q.root = &Node{e, nil}} else
	{
		q.root.add(e)
	}
}
