package main

// 数据只能interface中定义的方法而不能有单独方法；interface标准化的
// 所以可以在interface中定义方法，struct定义数据
// 可以在其中简单解释各个方法的作用，但最好是名字就能解释！
type BreadthFirstPathsClass interface {
	hasPath (v int) bool
	pathTo (v int) []int
	search (g *Graph, s int)
}

type DepthFirstPathsClass interface {
	hasPath (v int) bool
	pathTo (v int) []int
	search (g *Graph, s int)
	isMarked (s int) bool
	markCount () int // return how many vertices are marked
}

type CcClass interface {
	consume (g *Graph)
	connected (v,w int) bool
	idGet (v int) int
	countGet () int
}