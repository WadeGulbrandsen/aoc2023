package main

import (
	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 16

type direction rune

const (
	N direction = 'N'
	E direction = 'E'
	S direction = 'S'
	W direction = 'W'
)

type plane rune

const (
	H plane = 'H'
	V plane = 'V'
)

type lazer struct {
	position  grid.GridPoint
	direction direction
}

func (l *lazer) plane() plane {
	switch l.direction {
	case N, S:
		return V
	default:
		return H
	}
}

func immaFirinMahLazer(g *grid.Grid, start lazer) int {
	lazer_map := make(map[grid.GridPoint]rune)
	seen := make(map[lazer]bool)
	lazers := []lazer{start}
	for len(lazers) != 0 {
		var next []lazer
		for _, l := range lazers {
			if seen[l] {
				continue
			}
			seen[l] = true
			if !g.InBounds(l.position) {
				continue
			}
			map_rune, plane, grid_rune := lazer_map[l.position], l.plane(), g.At(l.position)
			switch {
			case (map_rune == '|' && plane == H) || (map_rune == '-' && plane == V):
				lazer_map[l.position] = '+'
			case l.plane() == H:
				lazer_map[l.position] = '-'
			default:
				lazer_map[l.position] = '|'
			}
			switch {
			case grid_rune == '|' && plane == H:
				next = append(next, lazer{l.position.N(), N})
				next = append(next, lazer{l.position.S(), S})
			case grid_rune == '-' && plane == V:
				next = append(next, lazer{l.position.E(), E})
				next = append(next, lazer{l.position.W(), W})
			case (grid_rune == '\\' && l.direction == N) || (grid_rune == '/' && l.direction == S):
				next = append(next, lazer{l.position.W(), W})
			case (grid_rune == '\\' && l.direction == S) || (grid_rune == '/' && l.direction == N):
				next = append(next, lazer{l.position.E(), E})
			case (grid_rune == '\\' && l.direction == E) || (grid_rune == '/' && l.direction == W):
				next = append(next, lazer{l.position.S(), S})
			case (grid_rune == '\\' && l.direction == W) || (grid_rune == '/' && l.direction == E):
				next = append(next, lazer{l.position.N(), N})
			default:
				switch l.direction {
				case N:
					next = append(next, lazer{l.position.N(), N})
				case E:
					next = append(next, lazer{l.position.E(), E})
				case S:
					next = append(next, lazer{l.position.S(), S})
				case W:
					next = append(next, lazer{l.position.W(), W})
				}
			}
		}
		lazers = next
	}
	return len(lazer_map)
}

func Problem1(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	return immaFirinMahLazer(&g, lazer{grid.GridPoint{X: 0, Y: 0}, E})
}

func Problem2(data *[]string) int {
	g := grid.MakeGridFromLines(data)
	m := 0
	for x := 0; x < g.MaxPoint.X; x++ {
		l1, l2 := lazer{grid.GridPoint{X: x, Y: 0}, S}, lazer{grid.GridPoint{X: x, Y: g.MaxPoint.Y - 1}, N}
		m = max(m, immaFirinMahLazer(&g, l1), immaFirinMahLazer(&g, l2))
	}
	for y := 0; y < g.MaxPoint.Y; y++ {
		l1, l2 := lazer{grid.GridPoint{X: 0, Y: y}, E}, lazer{grid.GridPoint{X: g.MaxPoint.X - 1, Y: y}, W}
		m = max(m, immaFirinMahLazer(&g, l1), immaFirinMahLazer(&g, l2))
	}
	return m
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
