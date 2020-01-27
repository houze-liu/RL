package main

import "fmt"

type CC struct {
marked map[int]bool
id map[int]int
count int
}

func (c *CC) consume (g *Graph) {
	c.id = make(map[int]int)
	c.marked = make(map[int]bool)
	for s:=0; s<g.V; s++ {
		if !c.marked[s] {
			c.dfs(g, s)
			c.count ++
		}
	}
}

func (c *CC) connected (v,w int) bool {
	if c.id[v] == c.id[w] {return true} else {return false}
}
// CC 所用的dfs和DFS所用dfs不是一个，区别在于接受者，前者是*CC，后者是*DFS
func (c *CC) dfs (g *Graph, s int) {
	c.marked[s] = true
	c.id[s] = c.count
	for _,v := range g.adj[s] {
		if !c.marked[v] {
			c.dfs(g, v)
			//当adj中第一个dfs结束时，v adj全被marked：
			flag := true
			for _, v_:= range g.adj[s] {
				if !c.marked[v_] {flag=false}
			}
			if !flag {fmt.Printf("Disconnected vertex: %v \n",s)}
		}
	}
}

func (c *CC) idGet (v int) int {
	return c.id[v]
}

func (c *CC) countGet () int {
	return c.count
}