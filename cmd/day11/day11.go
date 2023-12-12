package main

import (
	"github.com/WadeGulbrandsen/aoc2023/internal"
	"github.com/mowshon/iterium"
)

const Day = 11

func hasGalaxy(g *internal.Grid) (map[int]bool, map[int]bool) {
	rows, cols := make(map[int]bool), make(map[int]bool)
	for p := range g.Points {
		rows[p.Y] = true
		cols[p.X] = true
	}
	return rows, cols
}

func expandUniverse(g *internal.Grid, scale int) {
	if scale < 2 {
		return
	}
	rows, cols := hasGalaxy(g)
	new_points := make(map[internal.GridPoint]rune)
	for p, v := range g.Points {
		new_x, new_y := p.X, p.Y
		for x := 0; x < p.X; x++ {
			if !cols[x] {
				new_x += scale - 1
			}
		}
		for y := 0; y < p.Y; y++ {
			if !rows[y] {
				new_y += scale - 1
			}
		}
		g.Size.X = max(g.Size.X, new_x+1)
		g.Size.Y = max(g.Size.Y, new_y+1)
		new_points[internal.GridPoint{X: new_x, Y: new_y}] = v
	}
	g.Points = new_points
}

func expandAndFindDistances(data *[]string, scale int) int {
	g := internal.MakeGridFromLines(data)
	expandUniverse(&g, scale)
	combos := iterium.Combinations(internal.GetMapKeys(g.Points), 2)
	sum := 0
	for combo := range combos.Chan() {
		sum += combo[0].Distance(&combo[1])
	}
	return sum
}

func Problem1(data *[]string) int {
	return expandAndFindDistances(data, 2)
}

func Problem2(data *[]string) int {
	return expandAndFindDistances(data, 1000000)
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
