package main

import (
	"container/heap"
	"math"

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
	Cost      int
}

func (c *Crucible) Costless() Crucible {
	return Crucible{Point: c.Point, Direction: c.Direction, Traveled: c.Traveled}
}

func (c Crucible) Successors(g *CityGraph, min_steps, max_steps int) []Crucible {
	var to_check []Crucible
	if c.Traveled+1 < max_steps {
		forward := Crucible{Point: c.Point.Move(c.Direction, 1), Direction: c.Direction, Traveled: c.Traveled + 1}
		forward.Cost = c.Cost + g.blocks[forward.Point]
		if g.grid.InBounds(forward.Point) {
			to_check = append(to_check, forward)
		}
	}
	if c.Traveled+1 < min_steps {
		return to_check
	}

	left := Crucible{Point: c.Point.Move(c.Direction.TurnL(), 1), Direction: c.Direction.TurnL()}
	left.Cost = c.Cost + g.blocks[left.Point]
	if g.grid.InBounds(left.Point) {
		to_check = append(to_check, left)
	}
	right := Crucible{Point: c.Point.Move(c.Direction.TurnR(), 1), Direction: c.Direction.TurnR()}
	right.Cost = c.Cost + g.blocks[right.Point]
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
	for y := 0; y < g.grid.MaxPoint.Y; y++ {
		for x := 0; x < g.grid.MaxPoint.X; x++ {
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
	for y := 0; y < g.grid.MaxPoint.Y; y++ {
		for x := 0; x < g.grid.MaxPoint.X; x++ {
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

func findPath(g *CityGraph, min_steps, max_steps int) Crucible {
	pq := &priorityqueue.PriorityQueue[Crucible]{}
	seen := make(map[Crucible]bool)
	s, e := internal.GridPoint{X: 0, Y: 0}, internal.GridPoint{X: g.grid.MaxPoint.X - 1, Y: g.grid.MaxPoint.Y - 1}
	heap.Init(pq)
	east, south := s.Move(internal.E, 1), s.Move(internal.S, 1)
	heap.Push(pq, &priorityqueue.Item[Crucible]{Value: Crucible{Point: east, Direction: internal.E, Cost: g.blocks[east]}})
	heap.Push(pq, &priorityqueue.Item[Crucible]{Value: Crucible{Point: south, Direction: internal.S, Cost: g.blocks[south]}})
	for i := 0; pq.Len() > 0; i++ {
		c := heap.Pop(pq).(*priorityqueue.Item[Crucible]).Value
		cl := c.Costless()
		if seen[cl] {
			continue
		}
		seen[cl] = true
		if c.Point == e && c.Traveled+1 >= min_steps {
			return c
		}
		for _, n := range c.Successors(g, min_steps, max_steps) {
			heap.Push(pq, &priorityqueue.Item[Crucible]{Value: n, Priority: -n.Cost})
		}
	}
	return Crucible{}
}

func solve(data *[]string, min_steps, max_steps int) int {
	g := internal.MakeGridFromLines(data)
	heat_map := make(map[internal.GridPoint]int)
	for p, r := range g.Points {
		heat_map[p] = GetRuneValue(r)
	}
	graph := CityGraph{g, heat_map}
	p := findPath(&graph, min_steps, max_steps)
	return p.Cost
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
