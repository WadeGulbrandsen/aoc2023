package main

import (
	"container/heap"
	"math"
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	priorityqueue "github.com/WadeGulbrandsen/aoc2023/internal/priority_queue"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog/log"
)

const Day = 17

func GetRuneValue(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return math.MaxInt
}

type Crucible struct {
	Point     grid.GridPoint
	Direction grid.GridDirection
	Traveled  int
}

type CrucibleList struct {
	Crucible Crucible
	Cost     int
	Prev     *CrucibleList
}

func (cl CrucibleList) Slice() []Crucible {
	crucibles := []Crucible{cl.Crucible}
	prev := cl.Prev
	for prev != nil {
		crucibles = append(crucibles, prev.Crucible)
		prev = prev.Prev
	}
	slices.Reverse(crucibles)
	return crucibles
}

func (prev CrucibleList) Successors(g *CityGraph, min_steps, max_steps int) []CrucibleList {
	var to_check []Crucible
	c := prev.Crucible
	if c.Traveled+1 < max_steps {
		forward := Crucible{Point: c.Point.Move(c.Direction, 1), Direction: c.Direction, Traveled: c.Traveled + 1}
		if g.grid.InBounds(forward.Point) {
			to_check = append(to_check, forward)
		}
	}
	if c.Traveled+1 >= min_steps {
		left := Crucible{Point: c.Point.Move(c.Direction.TurnL(), 1), Direction: c.Direction.TurnL()}
		right := Crucible{Point: c.Point.Move(c.Direction.TurnR(), 1), Direction: c.Direction.TurnR()}
		to_check = append(to_check, left, right)
	}
	var successors []CrucibleList
	for _, next := range to_check {
		if g.grid.InBounds(next.Point) {
			successors = append(successors, CrucibleList{Crucible: next, Cost: prev.Cost + g.blocks[next.Point], Prev: &prev})
		}
	}
	return successors
}

type CityGraph struct {
	grid   grid.Grid
	blocks map[grid.GridPoint]int
}

func (g CityGraph) PrintSeen(searched map[Crucible]bool) string {
	visited := make(map[grid.GridPoint]bool)
	for k := range searched {
		visited[k.Point] = true
	}
	s := ""
	for y := 0; y < g.grid.MaxPoint.Y; y++ {
		for x := 0; x < g.grid.MaxPoint.X; x++ {
			if visited[grid.GridPoint{X: x, Y: y}] {
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
	pathmap := make(map[grid.GridPoint]rune)
	for _, c := range path {
		switch c.Direction {
		case grid.N:
			pathmap[c.Point] = '^'
		case grid.E:
			pathmap[c.Point] = '>'
		case grid.S:
			pathmap[c.Point] = 'v'
		case grid.W:
			pathmap[c.Point] = '<'
		}
	}
	s := ""
	for y := 0; y < g.grid.MaxPoint.Y; y++ {
		for x := 0; x < g.grid.MaxPoint.X; x++ {
			p := grid.GridPoint{X: x, Y: y}
			if r := pathmap[p]; r != 0 {
				s += string(r)
			} else {
				s += string(g.grid.At(p))
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

func findPath(g *CityGraph, min_steps, max_steps int) CrucibleList {
	pq := &priorityqueue.PriorityQueue[CrucibleList]{}
	seen := make(map[Crucible]bool)
	s, e := grid.GridPoint{X: 0, Y: 0}, grid.GridPoint{X: g.grid.MaxPoint.X - 1, Y: g.grid.MaxPoint.Y - 1}
	heap.Init(pq)
	east, south := s.Move(grid.E, 1), s.Move(grid.S, 1)
	heap.Push(pq, &priorityqueue.Item[CrucibleList]{Value: CrucibleList{Crucible: Crucible{Point: east, Direction: grid.E}, Cost: g.blocks[east]}})
	heap.Push(pq, &priorityqueue.Item[CrucibleList]{Value: CrucibleList{Crucible: Crucible{Point: south, Direction: grid.S}, Cost: g.blocks[south]}})
	for i := 0; pq.Len() > 0; i++ {
		cl := heap.Pop(pq).(*priorityqueue.Item[CrucibleList]).Value
		if seen[cl.Crucible] {
			continue
		}
		seen[cl.Crucible] = true
		if cl.Crucible.Point == e && cl.Crucible.Traveled+1 >= min_steps {
			return cl
		}
		for _, n := range cl.Successors(g, min_steps, max_steps) {
			heap.Push(pq, &priorityqueue.Item[CrucibleList]{Value: n, Priority: -(n.Cost + e.Distance(&n.Crucible.Point))})
		}
	}
	return CrucibleList{}
}

func solve(data *[]string, min_steps, max_steps int) int {
	g := grid.MakeGridFromLines(data)
	heat_map := make(map[grid.GridPoint]int)
	for p, r := range g.Points {
		heat_map[p] = GetRuneValue(r)
	}
	graph := CityGraph{g, heat_map}
	p := findPath(&graph, min_steps, max_steps)
	log.Debug().Msgf("Least heat loss is %v following:\n%v", p.Cost, graph.PrintPath(p.Slice()))
	return p.Cost
}

func Problem1(data *[]string) int {
	return solve(data, 1, 3)
}

func Problem2(data *[]string) int {
	return solve(data, 4, 10)
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
