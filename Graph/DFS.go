package main

import (
	. "playGround/Graph/DataStructures"
)

type DepthFirstPaths struct {
	marked map[int]bool
	edgeTo map[int]int
	s int
}

func (D *DepthFirstPaths) hasPath (v int) bool {
	return D.marked[v]
}

func (D *DepthFirstPaths) pathTo (v int) []int {
	var stack Iterable
	stack = &Stack{}
	if D.hasPath(v) {
		for x:=v; x!=D.s;x=D.edgeTo[x] {
			stack.Push(x)
		}

		stack.Push(D.s)
	}

	r := make([]int, 0)
	for !stack.IsEmpty() {
		r = append(r, stack.Pull().(int))
	}

	return r
}

func (D *DepthFirstPaths) search (g *Graph, s int) {
	D.marked = make(map[int]bool)
	D.edgeTo = make(map[int]int)
	D.s = s
	D.dfs(g, s)
}

// interface之外的方法可以定义在此处；dfs不能直接使用，但可以通过标准方法调用
// 此处search被调用，用来完成递归，从而支持search
func (D *DepthFirstPaths) dfs (g *Graph, s int) {
	D.marked[s] = true
	for _, v := range g.adj[s] {
		if !D.marked[v] {
			D.edgeTo[v] = s
			D.dfs(g, v)
		}
	}
}

func (D *DepthFirstPaths) isMarked (s int) bool {
	return D.marked[s]
}

func (D *DepthFirstPaths) markCount () int {
	count := 0
	for _,v := range D.marked {
		if v {count++}
	}

	return count
}

