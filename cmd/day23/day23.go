package main

import (
	"cmp"
	"container/heap"
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal/graph"
	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/heaps"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 23

type Path struct {
	Position grid.Point
	Steps    int
	Prev     *Path
}

type PathList []Path

func (p *Path) Contains(pos grid.Point) bool {
	if p.Position == pos {
		return true
	}
	if p.Prev != nil {
		return p.Prev.Contains(pos)
	}
	return false
}

func (p *Path) SlipperySuccessors(g *grid.Grid) PathList {
	var next PathList
	var to_check []grid.Point
	switch g.At(p.Position) {
	case '^':
		to_check = append(to_check, p.Position.N())
	case '>':
		to_check = append(to_check, p.Position.E())
	case 'v':
		to_check = append(to_check, p.Position.S())
	case '<':
		to_check = append(to_check, p.Position.W())
	default:
		to_check = []grid.Point{p.Position.N(), p.Position.E(), p.Position.S(), p.Position.W()}
	}
	for _, n := range to_check {
		dest := g.At(n)
		if g.InBounds(n) && dest != '#' && !p.Contains(n) {
			switch {
			case dest == '^' && n.N() == p.Position:
				continue
			case dest == '>' && n.E() == p.Position:
				continue
			case dest == '<' && n.S() == p.Position:
				continue
			case dest == '<' && n.W() == p.Position:
				continue
			}
			path := Path{Position: n, Steps: p.Steps + 1, Prev: p}
			next = append(next, path)
		}
	}
	return next
}

func (p *Path) Successors(g *grid.Grid) []grid.Point {
	var next []grid.Point
	for _, n := range [4]grid.Point{p.Position.N(), p.Position.E(), p.Position.S(), p.Position.W()} {
		if g.InBounds(n) && g.At(n) != '#' && !p.Contains(n) {
			next = append(next, n)
		}
	}
	return next
}

func (p *Path) Slice() []grid.Point {
	positions := []grid.Point{p.Position}
	prev := p.Prev
	for prev != nil {
		positions = append(positions, prev.Position)
		prev = prev.Prev
	}
	slices.Reverse(positions)
	return positions
}

func (p *Path) Print(g *grid.Grid) string {
	pathmap := make(map[grid.Point]rune)
	for _, pos := range p.Slice() {
		pathmap[pos] = 'o'
	}
	s := ""
	for y := 0; y < g.MaxPoint.Y; y++ {
		for x := 0; x < g.MaxPoint.X; x++ {
			pos := grid.Point{X: x, Y: y}
			if r := pathmap[pos]; r != 0 {
				s += string(r)
			} else {
				s += string(g.At(pos))
			}
		}
		s += "\n"
	}
	return s
}

func findPath(data *[]string) int {
	q := &heaps.PriorityQueue[Path]{}
	g := grid.MakeGridFromLines(data)
	s, e := grid.Point{X: 1, Y: 0}, grid.Point{X: g.MaxPoint.X - 2, Y: g.MaxPoint.Y - 1}
	heap.Init(q)
	heap.Push(q, &heaps.Item[Path]{Value: Path{Position: s}})
	var results PathList
	for q.Len() > 0 {
		path := heap.Pop(q).(*heaps.Item[Path]).Value
		if path.Position == e {
			results = append(results, path)
			continue
		}
		for _, n := range path.SlipperySuccessors(&g) {
			heap.Push(q, &heaps.Item[Path]{Value: n, Priority: n.Steps})
		}
	}
	if len(results) == 0 {
		return 0
	}
	longest := slices.MaxFunc(results, cmpPaths)
	return longest.Steps
}

func findPathThroughCrossroads(cr *graph.Graph[grid.Point, grid.Point], s, e grid.Point) Path {
	q := &heaps.PriorityQueue[Path]{}
	heap.Init(q)
	heap.Push(q, &heaps.Item[Path]{Value: Path{Position: s}})
	var results PathList
	for q.Len() > 0 {
		path := heap.Pop(q).(*heaps.Item[Path]).Value
		if path.Position == e {
			results = append(results, path)
			continue
		}
		next := cr.Vertices[path.Position].Edges
		for _, edge := range next {
			n := edge.Vertex.Item
			if !path.Contains(n) {
				new_path := Path{Position: n, Prev: &path, Steps: path.Steps + edge.Weight}
				heap.Push(q, &heaps.Item[Path]{Value: new_path, Priority: new_path.Steps})
			}
		}
	}
	if len(results) == 0 {
		return Path{}
	}
	longest := slices.MaxFunc(results, cmpPaths)
	return longest
}

func cmpPaths(a, b Path) int {
	return cmp.Compare(a.Steps, b.Steps)
}

func addCrossroad(cr *graph.Graph[grid.Point, grid.Point], p *Path) {
	path := p.Slice()
	l := len(path)
	if l < 2 {
		return
	}
	start, end := path[0], path[l-1]
	cr.AddVertex(start, start)
	cr.AddVertex(end, end)
	cr.AddEdge(start, end, l-1)
	cr.AddEdge(end, start, l-1)
}

type pointPair struct {
	a, b grid.Point
}

func findCrossroads(g *grid.Grid) *graph.Graph[grid.Point, grid.Point] {
	s, e := grid.Point{X: 1, Y: 0}, grid.Point{X: g.MaxPoint.X - 2, Y: g.MaxPoint.Y - 1}
	crossroads := graph.New[grid.Point, grid.Point]()
	paths := PathList{Path{Position: s}}
	seen := map[pointPair]bool{}
	for len(paths) > 0 {
		head := paths[0]
		paths = paths[1:]
		next := head.Successors(g)
		if head.Position == e {
			addCrossroad(crossroads, &head)
			continue
		}
		if len(next) == 1 {
			paths = append(paths, Path{Position: next[0], Prev: &head, Steps: head.Steps + 1})
			continue
		}
		if len(next) > 1 {
			addCrossroad(crossroads, &head)
			for _, n := range head.SlipperySuccessors(g) {
				pair := pointPair{head.Position, n.Position}
				if !seen[pair] {
					paths = append(paths, Path{Position: n.Position, Prev: &Path{Position: head.Position}})
					seen[pair] = true
				} else {
					continue
				}
			}
		}
	}
	return crossroads
}

func Problem1(data *[]string) int {
	return findPath(data)
}

func Problem2(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	s, e := grid.Point{X: 1, Y: 0}, grid.Point{X: g.MaxPoint.X - 2, Y: g.MaxPoint.Y - 1}
	c := findCrossroads(&g)
	longest := findPathThroughCrossroads(c, s, e)
	return longest.Steps
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
