package main

import (
	"container/heap"
	"math"
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal"
	priorityqueue "github.com/WadeGulbrandsen/aoc2023/internal/priority_queue"
)

const Day = 17

func GetRuneValue(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return math.MaxInt
}

type Crucible struct {
	Point     internal.GridPoint
	Direction internal.GridDirection
	Traveled  int
}

func (c Crucible) Successors(g *CityGraph, min_steps, max_steps int) []Crucible {
	var to_check []Crucible
	if c.Traveled+1 < max_steps {
		forward := Crucible{Point: c.Point.Move(c.Direction), Direction: c.Direction, Traveled: c.Traveled + 1}
		if g.grid.InBounds(forward.Point) {
			to_check = append(to_check, forward)
		}
	}
	if c.Traveled+1 < min_steps {
		return to_check
	}

	left := Crucible{Point: c.Point.Move(c.Direction.TurnL()), Direction: c.Direction.TurnL()}
	if g.grid.InBounds(left.Point) {
		to_check = append(to_check, left)
	}
	right := Crucible{Point: c.Point.Move(c.Direction.TurnR()), Direction: c.Direction.TurnR()}
	if g.grid.InBounds(right.Point) {
		to_check = append(to_check, right)
	}

	return to_check
}

type CityGraph struct {
	grid   internal.Grid
	blocks map[internal.GridPoint]int
}

func (g CityGraph) PrintSeen(searched map[Crucible]bool) string {
	visited := make(map[internal.GridPoint]bool)
	for k := range searched {
		visited[k.Point] = true
	}
	s := ""
	for y := 0; y < g.grid.Size.Y; y++ {
		for x := 0; x < g.grid.Size.X; x++ {
			if visited[internal.GridPoint{X: x, Y: y}] {
				s += "X"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (g CityGraph) PrintPath(path []Crucible) string {
	pathmap := make(map[internal.GridPoint]rune)
	prev := path[0].Direction.Reverse()
	run := 0
	for _, c := range path {
		if c.Direction == prev {
			run++
			pathmap[c.Point] = '0' + rune(run%10)
		} else {
			run = 1
			prev = c.Direction
			switch c.Direction {
			case internal.N:
				pathmap[c.Point] = '^'
			case internal.E:
				pathmap[c.Point] = '>'
			case internal.S:
				pathmap[c.Point] = 'v'
			case internal.W:
				pathmap[c.Point] = '<'
			}
		}
	}
	s := ""
	for y := 0; y < g.grid.Size.Y; y++ {
		for x := 0; x < g.grid.Size.X; x++ {
			p := internal.GridPoint{X: x, Y: y}
			if r := pathmap[p]; r != 0 {
				s += string(r)
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (g CityGraph) Cost(p []Crucible) int {
	sum := 0
	for _, c := range p {
		sum += g.blocks[c.Point]
	}
	return sum
}

func findPath(g *CityGraph, min_steps, max_steps int) []Crucible {
	pq := &priorityqueue.PriorityQueue[[]Crucible]{}
	seen := make(map[Crucible]bool)
	s, e := internal.GridPoint{X: 0, Y: 0}, internal.GridPoint{X: g.grid.Size.X - 1, Y: g.grid.Size.Y - 1}
	heap.Init(pq)
	heap.Push(pq, &priorityqueue.Item[[]Crucible]{Value: []Crucible{{Point: s.Move(internal.E), Direction: internal.E}}})
	heap.Push(pq, &priorityqueue.Item[[]Crucible]{Value: []Crucible{{Point: s.Move(internal.S), Direction: internal.S}}})
	for i := 0; pq.Len() > 0; i++ {
		path := heap.Pop(pq).(*priorityqueue.Item[[]Crucible]).Value
		c := path[len(path)-1]
		if seen[c] {
			continue
		}
		seen[c] = true
		if c.Point == e && c.Traveled+1 >= min_steps {
			return path
		}
		for _, n := range c.Successors(g, min_steps, max_steps) {
			np := slices.Clone(path)
			np = append(np, n)
			heap.Push(pq, &priorityqueue.Item[[]Crucible]{Value: np, Priority: -g.Cost(np)})
		}
	}
	return nil
}

func solve(data *[]string, min_steps, max_steps int) int {
	g := internal.MakeGridFromLines(data)
	heat_map := make(map[internal.GridPoint]int)
	for p, r := range g.Points {
		heat_map[p] = GetRuneValue(r)
	}
	graph := CityGraph{g, heat_map}
	p := findPath(&graph, min_steps, max_steps)
	// fmt.Println(graph.PrintPath(p))
	return graph.Cost(p)
}

func Problem1(data *[]string) int {
	return solve(data, 1, 3)
}

func Problem2(data *[]string) int {
	return solve(data, 4, 10)
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
