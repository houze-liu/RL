package main

import (
	"fmt"
	"strconv"
)

type Graph struct {
	V int
	E int
	adj [][]int
}

func (g *Graph) removeV (v int) {
	g.adj = append(g.adj[:v], g.adj[v+1:]...)
	for j,l := range g.adj {
		for i:=len(g.adj[j])-1;i>=0;i-- {
			if g.adj[j][len(g.adj[j])-i-1] == v{
				t := len(g.adj[j])-i-1
				g.adj[j] = append(l[:t],l[t+1:]...)
			} else if g.adj[j][len(g.adj[j])-i-1] > v {
				g.adj[j][len(g.adj[j])-i-1] = g.adj[j][len(g.adj[j])-i-1] - 1
			}
		}
	}
	g.V--
}

func (g *Graph)addEdge(v,w int){
	if v == w {fmt.Println("Self loop added!")}
	for _, v_:=range g.adj[w] {
		if v_ == v {
			fmt.Printf("Parallel edge: %v, %v added! \n", v,w)
		}
	}

	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.E ++
}

func copy(g Graph) Graph{
	g_ := Graph{g.V, g.E, make([][]int, len(g.adj))} // 即使结构是copy，但内层数据是一个
	for i,v := range g.adj{
		for _,w := range v{
			fmt.Println(w)
			g_.adj[i] = append(g_.adj[i],w)
		}
	}
	return g_
}

func (g *Graph)hasEdge(v, w int) bool {
	for _,p := range g.adj {
		for j,v_ := range p {
			if v == v_ {
				for _,w_ := range g.adj[j] {
					if w_ == w {return true}
				}
			}
		}
	}

	return false
}

func (g *Graph)toString() string{
	var s string
	str := strconv.Itoa
	s = str(g.V) + " vertices, " + str(g.E) + "edges\n"
	for v:=0; v<g.V; v++ {
		s += str(v) + " : "
		for _,e := range g.adj[v] {
			s +=  str(e) + " "
		}
		s += "\n"
	}

	return s
}

func main() {
	V := 12
	E := 16
	g := Graph{V,E,
		make([][]int, V)}
	g.addEdge(8,4)
	g.addEdge(2,3)
	g.addEdge(1, 11)
	g.addEdge(0,6)
	g.addEdge(3,6)
	g.addEdge(10,3)
	g.addEdge(7,11)
	g.addEdge(7,8)
	g.addEdge(11,8)
	g.addEdge(2,0)
	g.addEdge(6,2)
	g.addEdge(5,2)
	g.addEdge(5,10)
	g.addEdge(3,10)
	g.addEdge(8,1)
	g.addEdge(4,1)

	root := 0
	target := 10
	var BFS BreadthFirstPathsClass
	BFS = &BreadthFirstPaths{}
	BFS.search(&g, root)
	if BFS.hasPath(target){fmt.Printf("BFS finds path from %v to %v: ", root, target);
	fmt.Println(BFS.pathTo(target))} else {fmt.Println("No path found from %v to %v: ", root, target)}

	var DFS DepthFirstPathsClass
	DFS = &DepthFirstPaths{}
	DFS.search(&g, root)
	if DFS.hasPath(target){fmt.Printf("DFS finds path from %v to %v: ", root, target);
	fmt.Println(DFS.pathTo(target))} else {fmt.Println("No path found from %v to %v: ", root, target)}

	var Cc CcClass
	Cc = &CC{}
	Cc.consume(&g)
	fmt.Printf("Total Components: %v \n", Cc.countGet())
	fmt.Println(Cc.connected(0,10))
	fmt.Println(Cc.idGet(10))

}

