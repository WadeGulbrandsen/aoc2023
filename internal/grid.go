package internal

type GridDirection rune

const (
	N GridDirection = 'N'
	E GridDirection = 'E'
	S GridDirection = 'S'
	W GridDirection = 'W'
)

func (d GridDirection) TurnL() GridDirection {
	switch d {
	case N:
		return W
	case E:
		return N
	case S:
		return E
	case W:
		return S
	}
	return d
}

func (d GridDirection) TurnR() GridDirection {
	switch d {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	return d
}

func (d GridDirection) Reverse() GridDirection {
	switch d {
	case N:
		return S
	case E:
		return W
	case S:
		return N
	case W:
		return E
	}
	return d
}

type GridVector struct {
	Point     GridPoint
	Direction GridDirection
}

func (v *GridVector) TurnL() GridVector {
	return GridVector{v.Point, v.Direction.TurnL()}
}

func (v *GridVector) TurnR() GridVector {
	return GridVector{v.Point, v.Direction.TurnR()}
}

func (v *GridVector) Reverse() GridVector {
	return GridVector{v.Point, v.Direction.Reverse()}
}

func (v *GridVector) Next() GridVector {
	return GridVector{v.Point.Move(v.Direction), v.Direction}
}

func (v *GridVector) MoveL() GridVector {
	return GridVector{v.Point.Move(v.Direction.TurnL()), v.Direction.TurnL()}
}

func (v *GridVector) MoveR() GridVector {
	return GridVector{v.Point.Move(v.Direction.TurnR()), v.Direction.TurnR()}
}

type GridPoint struct {
	X, Y int
}

func (p *GridPoint) Move(d GridDirection) GridPoint {
	switch d {
	case N:
		return p.N()
	case E:
		return p.E()
	case S:
		return p.S()
	case W:
		return p.W()
	}
	return *p
}

func (p *GridPoint) Distance(o *GridPoint) int {
	dx := AbsDiff[int](p.X, o.X)
	dy := AbsDiff[int](p.Y, o.Y)
	return dx + dy
}

func (p *GridPoint) N() GridPoint {
	return GridPoint{p.X, p.Y - 1}
}

func (p *GridPoint) S() GridPoint {
	return GridPoint{p.X, p.Y + 1}
}

func (p *GridPoint) W() GridPoint {
	return GridPoint{p.X - 1, p.Y}
}

func (p *GridPoint) E() GridPoint {
	return GridPoint{p.X + 1, p.Y}
}

type Grid struct {
	Size   GridPoint
	Points map[GridPoint]rune
}

func (g *Grid) InBounds(p GridPoint) bool {
	return p.X >= 0 && p.X < g.Size.X && p.Y >= 0 && p.Y < g.Size.Y
}

func (g *Grid) At(p GridPoint) rune {
	return g.Points[p]
}

func (g Grid) String() string {
	s := ""
	for y := 0; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
			r := g.At(GridPoint{x, y})
			if r == 0 {
				r = '.'
			}
			s += string(r)
		}
		s += "\n"
	}
	return s
}

func (g *Grid) Rows() []string {
	lines := make([]string, g.Size.Y)
	for y := 0; y < g.Size.Y; y++ {
		line := make([]rune, g.Size.X)
		for x := 0; x < g.Size.X; x++ {
			if c := g.At(GridPoint{X: x, Y: y}); c != 0 {
				line[x] = c
			} else {
				line[x] = '.'
			}
		}
		lines[y] = string(line)
	}
	return lines
}

func (g *Grid) Columns() []string {
	lines := make([]string, g.Size.X)
	for x := 0; x < g.Size.X; x++ {
		line := make([]rune, g.Size.Y)
		for y := 0; y < g.Size.Y; y++ {
			if c := g.At(GridPoint{X: x, Y: y}); c != 0 {
				line[y] = c
			} else {
				line[y] = '.'
			}
		}
		lines[x] = string(line)
	}
	return lines
}

func MakeGridFromLines(lines *[]string) Grid {
	g := Grid{Points: make(map[GridPoint]rune)}
	g.Size.Y = len(*lines)
	if g.Size.Y > 0 {
		g.Size.X = len((*lines)[0])
		for y, l := range *lines {
			for x, r := range l {
				if r != '.' {
					g.Points[GridPoint{x, y}] = r
				}
			}
		}
	}
	return g
}
