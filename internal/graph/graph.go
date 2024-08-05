package graph

import "sync"

type key string

type Node struct {
	Parent   key
	Children map[key]int
	Weight   float32
}

type Graph struct {
	nodes map[key]*Node
	lock  *sync.RWMutex
}

func NewGraph() *Graph {
	g := new(Graph)
	g.nodes = make(map[key]*Node, 0)
	g.lock = new(sync.RWMutex)

	return g
}

func (g *Graph) Node(k string) *Node {
	g.lock.RLock()
	node := g.nodes[key(k)]
	g.lock.RUnlock()

	return node
}

func (g *Graph) UpsertNode(k, parent string, weight ...float32) {
	// Upsert node.
	g.lock.Lock()
	if g.nodes[key(k)] == nil {
		g.nodes[key(k)] = new(Node)
	}
	g.nodes[key(k)].Parent = key(parent)
	if len(weight) > 0 {
		g.nodes[key(k)].Weight += weight[0]
	}
	g.lock.Unlock()

	// Update parent's children.
	if g.Node(parent) == nil {
		g.lock.Lock()
		g.nodes[key(parent)] = new(Node)
		g.lock.Unlock()
	}
	if g.Node(parent).Children == nil {
		g.lock.Lock()
		g.nodes[key(parent)].Children = make(map[key]int, 0)
		g.lock.Unlock()
	}

	g.lock.Lock()
	g.nodes[key(parent)].Children[key(k)]++
	g.lock.Unlock()
}
