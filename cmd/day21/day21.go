package main

import (
	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 21

func Problem1(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	steps := 64
	if g.MaxPoint.X < 12 {
		steps = 6
	}
	current := utils.KeysByVal(g.Points, 'S')
	if len(current) != 1 {
		return 0
	}
	seen := make(map[grid.Point]bool)
	for i := 0; i <= steps; i++ {
		var next []grid.Point
		for _, p := range current {
			if seen[p] {
				continue
			}
			seen[p] = true
			if i%2 == 0 {
				g.Points[p] = 'O'
			}
			for _, n := range [4]grid.Point{p.N(), p.E(), p.S(), p.W()} {
				if !seen[n] && g.At(n) == 0 {
					next = append(next, n)
				}
			}
		}
		current = next
	}
	possible := utils.KeysByVal(g.Points, 'O')
	return len(possible)
}

func remap(p, g *grid.Point) grid.Point {
	x, y := ((p.X%g.X)+g.X)%g.X, ((p.Y%g.Y)+g.Y)%g.Y
	return grid.Point{X: x, Y: y}
}

func Problem2(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	current := utils.KeysByVal(g.Points, 'S')
	if len(current) != 1 || g.MaxPoint.X != g.MaxPoint.Y {
		return 0
	}
	steps := 26501365
	seen := make(map[grid.Point]bool)
	even := 0
	odd := 0
	var xs []int
	for i := 0; i < steps; i++ {
		var next []grid.Point
		for _, p := range current {
			if seen[p] {
				continue
			}
			seen[p] = true
			if i%2 == 0 {
				even++
			} else {
				odd++
			}
			for _, n := range [4]grid.Point{p.N(), p.E(), p.S(), p.W()} {
				remapped := remap(&n, &g.MaxPoint)
				if !seen[n] && g.At(remapped) != '#' {
					next = append(next, n)
				}
			}
		}
		current = next
		if i%g.MaxPoint.X == steps%g.MaxPoint.X {
			if i%2 == 0 {
				xs = append(xs, (even))
			} else {
				xs = append(xs, (odd))
			}
			if len(xs) == 3 {
				b0, b1, b2 := xs[0], xs[1]-xs[0], xs[2]-xs[1]
				x := steps / g.MaxPoint.X
				answer := b0 + b1*x + (x*(x-1)/2)*(b2-b1)
				return answer
			}
		}
	}
	return 0
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
