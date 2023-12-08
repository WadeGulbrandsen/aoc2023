package main

import (
	"slices"
	"strconv"
	"sync"
	"unicode"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 3

type Point struct {
	X, Y int
}

type NumOnGrid struct {
	StartX, EndX, Y, Value int
	grid                   *Grid
}

type Grid struct {
	lock    sync.RWMutex
	symbols map[Point]rune
	digits  map[Point]rune
	size    Point
	numbers []NumOnGrid
}

func (n *NumOnGrid) AdjacentPoints() []Point {
	var p []Point
	size := n.grid.Size()
	for y := max(0, n.Y-1); y < min(size.Y, n.Y+2); y++ {
		for x := max(0, n.StartX-1); x < min(size.X, n.EndX+2); x++ {
			if y != n.Y || x < n.StartX || x > n.EndX {
				p = append(p, Point{x, y})
			}
		}
	}
	return p
}

func (n *NumOnGrid) String() string {
	size := n.grid.Size()
	s := ""
	for y := max(0, n.Y-1); y < min(size.Y, n.Y+2); y++ {
		for x := max(0, n.StartX-1); x < min(size.X, n.EndX+2); x++ {
			s += string(n.grid.At(Point{x, y}))
		}
		s += "\n"
	}
	return s
}

func (g *Grid) Size() Point {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.size
}

func (g *Grid) At(p Point) rune {
	g.lock.RLock()
	defer g.lock.RUnlock()
	if s := g.symbols[p]; s != 0 {
		return s
	}
	if n := g.digits[p]; n != 0 {
		return n
	}
	return '.'
}

func (g *Grid) HasSymbolAt(p Point) bool {
	g.lock.RLock()
	defer g.lock.RUnlock()
	s := g.symbols[p]
	return s != 0
}

func (g *Grid) String() string {
	g.lock.RLock()
	defer g.lock.RUnlock()
	grid := ""
	for y := 0; y < g.size.Y; y++ {
		for x := 0; x < g.size.X; x++ {
			grid += string(g.At(Point{x, y}))
		}
		grid += "\n"
	}
	return grid
}

func parseLineToGrid(line string, y int, g *Grid, wg *sync.WaitGroup) {
	defer wg.Done()
	s := make(map[Point]rune)
	d := make(map[Point]rune)
	l, num_start := 0, 0
	var nums []NumOnGrid
	num_str := ""
	for x, c := range line {
		l++
		if unicode.IsDigit(c) {
			if num_str == "" {
				num_start = x
			}
			num_str += string(c)
			d[Point{x, y}] = c
		} else {
			if c != '.' {
				s[Point{x, y}] = c
			}
			if num_str != "" {
				if v, err := strconv.Atoi(num_str); err == nil {
					nums = append(nums, NumOnGrid{num_start, x - 1, y, v, g})
				}
				num_str = ""
				num_start = 0
			}
		}
	}
	if num_str != "" {
		if v, err := strconv.Atoi(num_str); err == nil {
			nums = append(nums, NumOnGrid{num_start, l, y, v, g})
		}
		num_str = ""
		num_start = 0
	}
	g.lock.Lock()
	for k, v := range s {
		g.symbols[k] = v
	}
	for k, v := range d {
		g.digits[k] = v
	}
	if y == 0 {
		g.size.X = l
	}
	g.numbers = append(g.numbers, nums...)
	g.lock.Unlock()
}

func sliceToGrid(data *[]string) *Grid {
	var wg sync.WaitGroup
	g := Grid{symbols: make(map[Point]rune), digits: make(map[Point]rune)}

	for i, s := range *data {
		wg.Add(1)
		go parseLineToGrid(s, i, &g, &wg)
	}
	wg.Wait()

	g.lock.Lock()
	g.size.Y = len(*data)
	g.lock.Unlock()

	return &g
}

func numBySymbol(num NumOnGrid, ch chan int) {
	for _, p := range num.AdjacentPoints() {
		if num.grid.HasSymbolAt(p) {
			ch <- num.Value
			return
		}
	}
	ch <- 0
}

func Problem1(data *[]string) int {
	sum := 0
	g := sliceToGrid(data)
	g.lock.RLock()
	ch := make(chan int)
	l := len(g.numbers)
	for _, n := range g.numbers {
		go numBySymbol(n, ch)
	}
	g.lock.RUnlock()

	for i := 0; i < l; i++ {
		sum += <-ch
	}

	return sum
}

func getGearRatio(p Point, g *Grid, ch chan int) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	var adjacent_numbers []NumOnGrid
	for _, n := range g.numbers {
		if slices.Contains(n.AdjacentPoints(), p) {
			adjacent_numbers = append(adjacent_numbers, n)
		}
	}
	if len(adjacent_numbers) == 2 {
		ch <- adjacent_numbers[0].Value * adjacent_numbers[1].Value
		return
	}
	ch <- 0
}

func Problem2(data *[]string) int {
	sum := 0
	g := sliceToGrid(data)
	g.lock.RLock()
	ch := make(chan int)
	l := 0
	for p, c := range g.symbols {
		if c == '*' {
			l++
			go getGearRatio(p, g, ch)
		}
	}
	g.lock.RUnlock()

	for i := 0; i < l; i++ {
		sum += <-ch
	}

	return sum
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
