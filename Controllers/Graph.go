package Controllers

import "fmt"

type Graph struct {
	vertices []*vertex
}

type vertex struct {
	key      int
	adjacent []*vertex
}

func (g *Graph) addEdge(from, to int) {
	for _, v := range g.vertices {
		if v.key == from {
			toVertex := &vertex{key: to}
			if contains(v.adjacent, toVertex.key) {
				fmt.Println("already exists")
				continue
			}
			v.adjacent = append(v.adjacent, toVertex)
		}
	}
}

func (g *Graph) AddVertex(k int) {
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &vertex{key: k})
	}
}

func contains(s []*vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}

	return false
}

func (g *Graph) Print() {
	for _, val := range g.vertices {
		fmt.Printf("\nVertex %v :", val.key)
		for _, val = range val.adjacent {
			fmt.Printf("%v, ", val.key)
		}
	}
}
func GraphMain() {
	test := &Graph{}
	for i := 1; i <= 5; i++ {
		test.AddVertex(i)
	}
	test.addEdge(1, 2)
	test.addEdge(1, 3)
	test.addEdge(1, 3)
	test.addEdge(2, 5)
	test.Print()
}
