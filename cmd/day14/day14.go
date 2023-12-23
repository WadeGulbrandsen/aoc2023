package main

import (
	"maps"

	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 14

func tiltNorth(g *grid.Grid) {
	for y := 1; y < g.MaxPoint.Y; y++ {
		for x := 0; x < g.MaxPoint.X; x++ {
			pos := grid.Point{X: x, Y: y}
			if r := g.At(pos); r == 'O' {
				new_pos := pos
				for i := 0; i < y; i++ {
					if to_check := new_pos.N(); g.At(to_check) == 0 {
						new_pos = to_check
					} else {
						break
					}
				}
				if pos != new_pos {
					g.Points[new_pos] = r
					delete(g.Points, pos)
				}
			}
		}
	}
}

func tiltEast(g *grid.Grid) {
	for x := 1; x < g.MaxPoint.X; x++ {
		for y := 0; y < g.MaxPoint.Y; y++ {
			pos := grid.Point{X: g.MaxPoint.X - x - 1, Y: y}
			if r := g.At(pos); r == 'O' {
				new_pos := pos
				for i := 0; i < x; i++ {
					if to_check := new_pos.E(); g.At(to_check) == 0 {
						new_pos = to_check
					} else {
						break
					}
				}
				if pos != new_pos {
					g.Points[new_pos] = r
					delete(g.Points, pos)
				}
			}
		}
	}
}

func tiltSouth(g *grid.Grid) {
	for y := 1; y < g.MaxPoint.Y; y++ {
		for x := 0; x < g.MaxPoint.X; x++ {
			pos := grid.Point{X: x, Y: g.MaxPoint.Y - y - 1}
			if r := g.At(pos); r == 'O' {
				new_pos := pos
				for i := 0; i < y; i++ {
					if to_check := new_pos.S(); g.At(to_check) == 0 {
						new_pos = to_check
					} else {
						break
					}
				}
				if pos != new_pos {
					g.Points[new_pos] = r
					delete(g.Points, pos)
				}
			}
		}
	}
}

func tiltWest(g *grid.Grid) {
	for x := 1; x < g.MaxPoint.X; x++ {
		for y := 0; y < g.MaxPoint.Y; y++ {
			pos := grid.Point{X: x, Y: y}
			if r := g.At(pos); r == 'O' {
				new_pos := pos
				for i := 0; i < x; i++ {
					if to_check := new_pos.W(); g.At(to_check) == 0 {
						new_pos = to_check
					} else {
						break
					}
				}
				if pos != new_pos {
					g.Points[new_pos] = r
					delete(g.Points, pos)
				}
			}
		}
	}
}

func spin(g *grid.Grid) {
	tiltNorth(g)
	tiltWest(g)
	tiltSouth(g)
	tiltEast(g)
}

func calculateLoad(g *grid.Grid) int {
	load := 0
	for y := 0; y < g.MaxPoint.Y; y++ {
		for x := 0; x < g.MaxPoint.X; x++ {
			if g.At(grid.Point{X: x, Y: y}) == 'O' {
				load += g.MaxPoint.Y - y
			}
		}
	}
	return load
}

func Problem1(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	tiltNorth(&g)
	return calculateLoad(&g)
}

func getHistoryIndex(history *[]map[grid.Point]rune, m map[grid.Point]rune) int {
	for i, p := range *history {
		if maps.Equal(p, m) {
			return i
		}
	}
	return -1
}

func findSpinCycle(g *grid.Grid) (int, int, []map[grid.Point]rune) {
	var history []map[grid.Point]rune
	for {
		spin(g)
		if idx := getHistoryIndex(&history, g.Points); idx != -1 {
			cycle := history[idx:]
			cycle_length := len(cycle)
			return idx, cycle_length, cycle
		}
		history = append(history, maps.Clone(g.Points))
	}
}

func Problem2(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	start, length, cycle := findSpinCycle(&g)
	if length == 0 {
		return 0
	}
	idx := (1000000000 - start - 1) % length
	end_state := grid.Grid{MaxPoint: g.MaxPoint, Points: cycle[idx]}
	return calculateLoad(&end_state)
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
