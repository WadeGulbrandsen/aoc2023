package graph

type Graph[K comparable, V any] struct {
	Vertices map[K]*Vertex[K, V]
}

type Vertex[K comparable, V any] struct {
	Item  V
	Edges map[K]*Edge[K, V]
}

type Edge[K comparable, V any] struct {
	Weight int
	Vertex *Vertex[K, V]
}

func (g *Graph[K, V]) AddVertex(key K, item V) {
	if _, ok := g.Vertices[key]; ok {
		return
	}
	g.Vertices[key] = &Vertex[K, V]{Item: item, Edges: map[K]*Edge[K, V]{}}
}

func (g *Graph[K, V]) AddEdge(source, dest K, weight int) {
	if _, ok := g.Vertices[source]; !ok {
		return
	}
	if _, ok := g.Vertices[dest]; !ok {
		return
	}
	g.Vertices[source].Edges[dest] = &Edge[K, V]{Weight: weight, Vertex: g.Vertices[dest]}
}

func (g *Graph[K, V]) Neighbors(key K) []V {
	result := []V{}
	for _, edge := range g.Vertices[key].Edges {
		result = append(result, edge.Vertex.Item)
	}
	return result
}

func New[K comparable, V any]() *Graph[K, V] {
	return &Graph[K, V]{Vertices: map[K]*Vertex[K, V]{}}
}
