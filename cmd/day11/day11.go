package main

import (
	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/mowshon/iterium"
)

const Day = 11

func hasGalaxy(g *grid.Grid) (map[int]bool, map[int]bool) {
	rows, cols := make(map[int]bool), make(map[int]bool)
	for p := range g.Points {
		rows[p.Y] = true
		cols[p.X] = true
	}
	return rows, cols
}

func expandUniverse(g *grid.Grid, scale int) {
	if scale < 2 {
		return
	}
	rows, cols := hasGalaxy(g)
	new_points := make(map[grid.GridPoint]rune)
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
		g.MaxPoint.X = max(g.MaxPoint.X, new_x+1)
		g.MaxPoint.Y = max(g.MaxPoint.Y, new_y+1)
		new_points[grid.GridPoint{X: new_x, Y: new_y}] = v
	}
	g.Points = new_points
}

func expandAndFindDistances(data *[]string, scale int) int {
	g := grid.MakeGridFromLines(data)
	expandUniverse(&g, scale)
	combos := iterium.Combinations(utils.GetMapKeys(g.Points), 2)
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
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
