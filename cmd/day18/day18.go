package main

import (
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 18

type digInstruction struct {
	direction internal.GridDirection
	distance  int
}

type digPlan []digInstruction

func (d digPlan) Dig() internal.Grid {
	min_x, min_y, max_x, max_y := 0, 0, 0, 0
	current := internal.GridPoint{X: 0, Y: 0}
	holes := map[internal.GridPoint]rune{current: '#'}
	for _, inst := range d {
		for i := 0; i < inst.distance; i++ {
			current = current.Move(inst.direction)
			holes[current] = '#'
		}
		min_x = min(min_x, current.X)
		max_x = max(max_x, current.X)
		min_y = min(min_y, current.Y)
		max_y = max(max_y, current.Y)
	}
	minPoint := internal.GridPoint{X: min_x, Y: min_y}
	maxPoint := internal.GridPoint{X: max_x + 1, Y: max_y + 1}
	grid := internal.Grid{MinPoint: minPoint, MaxPoint: maxPoint, Points: holes}
	// fmt.Print(grid)
	for x := min_x; x < max_x; x++ {
		c := internal.GridPoint{X: x, Y: min_y}
		d := internal.GridPoint{X: x, Y: min_y + 1}
		if holes[c] == '#' && holes[d] != '#' {
			grid.Fill(d, 'X')
			return grid
		}
	}
	return grid
}

func standardParse(d, n string) (digInstruction, bool) {
	inst := digInstruction{}
	switch d {
	case "U":
		inst.direction = internal.N
	case "D":
		inst.direction = internal.S
	case "L":
		inst.direction = internal.W
	case "R":
		inst.direction = internal.E
	default:
		return inst, false
	}
	if i, e := strconv.Atoi(n); e == nil {
		inst.distance = i
		return inst, true
	}
	return inst, false
}

func colorParse(s string) (digInstruction, bool) {
	inst := digInstruction{}
	s = strings.Trim(s, "()#")
	if len(s) < 2 {
		return inst, false
	}
	h, d := s[:len(s)-1], s[len(s)-1:]
	switch d {
	case "3":
		inst.direction = internal.N
	case "1":
		inst.direction = internal.S
	case "2":
		inst.direction = internal.W
	case "0":
		inst.direction = internal.E
	default:
		return inst, false
	}
	if i, e := strconv.ParseInt(h, 16, 64); e == nil {
		inst.distance = int(i)
		return inst, true
	}
	return inst, true
}

func strToDigInstruction(s string, color_parse bool) (digInstruction, bool) {
	parts := strings.Split(s, " ")
	inst := digInstruction{}
	if len(parts) != 3 {
		return inst, false
	}
	if color_parse {
		return colorParse(parts[2])
	}
	return standardParse(parts[0], parts[1])
}

func Solve(data *[]string, color_parse bool) int {
	var plan digPlan
	for _, s := range *data {
		if di, found := strToDigInstruction(s, color_parse); found {
			plan = append(plan, di)
		}
	}
	grid := plan.Dig()
	return len(grid.Points)
}

func Problem1(data *[]string) int {
	return Solve(data, false)
}

func Problem2(data *[]string) int {
	return 0
}

func main() {
	internal.RunSolutions(Day, Problem1, Problem2, "input.txt", "input.txt", -1)
}
