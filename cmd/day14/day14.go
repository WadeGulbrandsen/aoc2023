package main

import (
	"maps"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 14

func tiltNorth(g *internal.Grid) {
	for y := 1; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
			pos := internal.GridPoint{X: x, Y: y}
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

func tiltEast(g *internal.Grid) {
	for x := 1; x < g.Size.X; x++ {
		for y := 0; y < g.Size.Y; y++ {
			pos := internal.GridPoint{X: g.Size.X - x - 1, Y: y}
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

func tiltSouth(g *internal.Grid) {
	for y := 1; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
			pos := internal.GridPoint{X: x, Y: g.Size.Y - y - 1}
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

func tiltWest(g *internal.Grid) {
	for x := 1; x < g.Size.X; x++ {
		for y := 0; y < g.Size.Y; y++ {
			pos := internal.GridPoint{X: x, Y: y}
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

func spin(g *internal.Grid) {
	tiltNorth(g)
	tiltWest(g)
	tiltSouth(g)
	tiltEast(g)
}

func calculateLoad(g *internal.Grid) int {
	load := 0
	for y := 0; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
			if g.At(internal.GridPoint{X: x, Y: y}) == 'O' {
				load += g.Size.Y - y
			}
		}
	}
	return load
}

func Problem1(data *[]string) int {
	g := internal.MakeGridFromLines(data)
	tiltNorth(&g)
	return calculateLoad(&g)
}

func getHistoryIndex(history *[]map[internal.GridPoint]rune, m map[internal.GridPoint]rune) int {
	for i, p := range *history {
		if maps.Equal(p, m) {
			return i
		}
	}
	return -1
}

func findSpinCycle(g *internal.Grid) (int, int, []map[internal.GridPoint]rune) {
	var history []map[internal.GridPoint]rune
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
	g := internal.MakeGridFromLines(data)
	start, length, cycle := findSpinCycle(&g)
	if length == 0 {
		return 0
	}
	idx := (1000000000 - start - 1) % length
	end_state := internal.Grid{Size: g.Size, Points: cycle[idx]}
	return calculateLoad(&end_state)
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
