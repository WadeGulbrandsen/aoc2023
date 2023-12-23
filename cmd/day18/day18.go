package main

import (
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 18

type digInstruction struct {
	direction grid.Direction
	distance  int
}

type digPlan []digInstruction

func (d digPlan) DigTrench() []grid.Point {
	current := grid.Point{X: 0, Y: 0}
	trench := []grid.Point{current}
	for _, inst := range d {
		current = current.Move(inst.direction, inst.distance)
		trench = append(trench, current)
	}
	return trench
}

func standardParse(d, n string) (digInstruction, bool) {
	inst := digInstruction{}
	switch d {
	case "U":
		inst.direction = grid.N
	case "D":
		inst.direction = grid.S
	case "L":
		inst.direction = grid.W
	case "R":
		inst.direction = grid.E
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
		inst.direction = grid.N
	case "1":
		inst.direction = grid.S
	case "2":
		inst.direction = grid.W
	case "0":
		inst.direction = grid.E
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
	trench := plan.DigTrench()
	area := grid.ShoelaceArea(trench)
	return area
}

func Problem1(data *[]string) int {
	return Solve(data, false)
}

func Problem2(data *[]string) int {
	return Solve(data, true)
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
