package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 10
const render = false
const showNthFrame = 100

type XY struct {
	x, y int
}

func (xy *XY) N() XY {
	return XY{xy.x, xy.y - 1}
}

func (xy *XY) S() XY {
	return XY{xy.x, xy.y + 1}
}

func (xy *XY) W() XY {
	return XY{xy.x - 1, xy.y}
}

func (xy *XY) E() XY {
	return XY{xy.x + 1, xy.y}
}

type Pipe rune

type PipeMaze struct {
	start    XY
	size     XY
	tiles    map[XY]Pipe
	expanded bool
}

func (pm PipeMaze) String() string {
	str := "  "
	for x := 0; x < pm.size.x; x++ {
		str += fmt.Sprint(x % 10)
	}
	str += "\n"
	for y := 0; y < pm.size.y; y++ {
		str += fmt.Sprintf("%v ", y%10)
		for x := 0; x < pm.size.x; x++ {
			xy := XY{x, y}
			r := pm.At(xy)
			if r == 0 {
				r = '.'
			}
			str += string(r)
		}
		str += "\n"
	}
	return str
}

func (pm *PipeMaze) Exits(xy XY) []XY {
	var exits []XY
	r := pm.At(xy)
	switch r {
	case '|':
		exits = append(exits, xy.N(), xy.S())
	case '-':
		exits = append(exits, xy.W(), xy.E())
	case 'L':
		exits = append(exits, xy.N(), xy.E())
	case 'J':
		exits = append(exits, xy.N(), xy.W())
	case '7':
		exits = append(exits, xy.S(), xy.W())
	case 'F':
		exits = append(exits, xy.S(), xy.E())
	case 'S':
		exits = append(exits, xy.N(), xy.S(), xy.E(), xy.W())
	}
	return exits
}

func (pm *PipeMaze) NextPipe(xy XY, prev XY) (XY, error) {
	exits := pm.Exits(xy)
	if !slices.Contains(exits, prev) {
		return XY{}, fmt.Errorf("cannot come from %v to %v '%v'", prev, xy, pm.At(xy))
	}
	for _, next := range exits {
		if next != prev && slices.Contains(pm.Exits(next), xy) {
			return next, nil
		}
	}
	return XY{}, fmt.Errorf("could not find next pipe from %v through '%v' at %v", prev, pm.At(xy), xy)
}

func (pm *PipeMaze) At(xy XY) Pipe {
	return pm.tiles[xy]
}

func (pm *PipeMaze) AddRow(s string, y int) {
	for x, r := range s {
		xy := XY{x, y}
		if r != '.' {
			pm.tiles[xy] = Pipe(r)
		}
		if r == 'S' {
			pm.start = xy
		}
	}
}

func (pm *PipeMaze) FindLoop() []XY {
	for _, xy := range pm.Exits(pm.start) {
		if pm.At(xy) == 0 {
			continue
		}
		var path []XY
		prev := pm.start
		current := xy
		for {
			path = append(path, current)
			next, err := pm.NextPipe(current, prev)
			if err != nil {
				break
			}
			if next == pm.start {
				return path
			}
			prev, current = current, next
		}
	}
	return nil
}

func (pm *PipeMaze) loopOnly() {
	loop := pm.FindLoop()
	for xy := range pm.tiles {
		if !slices.Contains(loop, xy) && xy != pm.start {
			delete(pm.tiles, xy)
		}
	}
}

func (pm *PipeMaze) expand() {
	if pm.expanded {
		return
	}
	tiles := make(map[XY]Pipe)
	for xy, r := range pm.tiles {
		doubled := XY{xy.x * 2, xy.y * 2}
		tiles[doubled] = r
		exits := pm.Exits(xy)
		n, w := xy.N(), xy.W()
		if slices.Contains(exits, n) {
			tiles[doubled.N()] = '|'
		}
		if slices.Contains(exits, w) {
			tiles[doubled.W()] = '-'
		}
	}
	pm.size = XY{pm.size.x * 2, pm.size.y * 2}
	pm.tiles = tiles
	pm.expanded = true
}

func (pm *PipeMaze) contract() {
	if !pm.expanded {
		return
	}
	tiles := make(map[XY]Pipe)
	for xy, r := range pm.tiles {
		x, xr := utils.DivMod(xy.x, 2)
		y, yr := utils.DivMod(xy.y, 2)
		if xr == 0 && yr == 0 {
			tiles[XY{x, y}] = r
		}
	}
	pm.size = XY{pm.size.x / 2, pm.size.y / 2}
	pm.tiles = tiles
	pm.expanded = false
}

func (pm *PipeMaze) InBounds(xy XY) bool {
	return xy.x >= 0 && xy.x < pm.size.x && xy.y >= 0 && xy.y < pm.size.y
}

func (pm *PipeMaze) fillOutside(images *[]*image.Paletted) {
	var to_visit []XY
	for y := 0; y < pm.size.y; y++ {
		for x := 0; x < pm.size.x; x++ {
			if x == 0 || y == 0 || x == pm.size.x-1 || y == pm.size.y-1 {
				xy := XY{x, y}
				if pm.At(xy) == 0 {
					to_visit = append(to_visit, xy)
				}
			}
		}
	}
	frame := 0
	for len(to_visit) != 0 {
		xy := to_visit[0]
		to_visit = to_visit[1:]
		pm.tiles[xy] = 'O'
		for _, next := range [4]XY{xy.N(), xy.S(), xy.E(), xy.W()} {
			if pm.At(next) == 0 && pm.InBounds(next) && !slices.Contains(to_visit, next) {
				to_visit = append(to_visit, next)
			}
		}
		if render && xy.x%2 == 0 && xy.y%2 == 0 {
			frame++
			if frame%showNthFrame == 0 {
				*images = append(*images, drawMaze(pm))
			}
		}
	}
}

func (pm *PipeMaze) FindInclosed(images *[]*image.Paletted) []XY {
	pm.loopOnly()
	pm.expand()
	if render {
		*images = append(*images, drawMaze(pm))
	}
	pm.fillOutside(images)
	var inclosed []XY
	pm.contract()
	frame := 0
	for y := 0; y < pm.size.y; y++ {
		for x := 0; x < pm.size.x; x++ {
			xy := XY{x, y}
			if pm.At(xy) == 0 {
				inclosed = append(inclosed, xy)
				pm.tiles[xy] = 'I'
				if render {
					frame++
					if frame%showNthFrame == 0 {
						*images = append(*images, drawMaze(pm))
					}
				}
			}
		}
	}
	return inclosed
}

func linesToMaze(data *[]string) PipeMaze {
	maze := PipeMaze{tiles: make(map[XY]Pipe), size: XY{len((*data)[0]), len(*data)}}
	for i, s := range *data {
		maze.AddRow(s, i)
	}
	return maze
}

func Problem1(data *[]string) int {
	maze := linesToMaze(data)
	loop := maze.FindLoop()
	q, r := utils.DivMod(len(loop), 2)
	return q + r
}

func Problem2(data *[]string) int {
	var images []*image.Paletted
	maze := linesToMaze(data)
	inclosed := maze.FindInclosed(&images)
	if render {
		images = append(images, drawMaze(&maze))
		delays := make([]int, len(images))
		delays[0], delays[len(delays)-1] = 100, 100
		utils.WriteAGif(&images, &delays, "problem2.gif")
	}
	return len(inclosed)
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
