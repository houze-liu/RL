package main

import ."playGround/Graph/DataStructures"

type BreadthFirstPaths struct {
	marked map[int]bool
	edgeTo map[int]int
	s int
}

func (B *BreadthFirstPaths) hasPath (v int) bool {
	return B.marked[v]
}

func (B *BreadthFirstPaths) pathTo (v int) []int {
	var stack Iterable
	stack = &Stack{}
	if B.hasPath(v) {
		for x:=v; x!=B.s;x=B.edgeTo[x] {
			stack.Push(x)
		}

		stack.Push(B.s)
	}

	r := make([]int, 0)
	for !stack.IsEmpty() {
		r = append(r, stack.Pull().(int))
	}

	return r
}

func (B *BreadthFirstPaths) search (g *Graph, s int) {
	B.marked = make(map[int]bool)
	B.edgeTo = make(map[int]int)
	B.s = s
	B.bfs(g, s)
}

func (B *BreadthFirstPaths) bfs (g *Graph, s int) {
	var q Iterable
	q = &Queue{}
	q.Push(s)
	B.marked[s] = true
	for !q.IsEmpty() {
		v:= q.Pull().(int)
		for _,v_:=range g.adj[v] {
			if !B.marked[v_] {
				B.edgeTo[v_] = v
				B.marked[v_] = true
				q.Push(v_)
			}
		}
	}
}