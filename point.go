package main

type Point struct {
	symbol    int
	x         int
	y         int
	direction int
}

const (
	N int = iota
	NW
	W
	SW
	S
	SE
	E
	NE
)

func BuildPoint(symbol, x, y, direction int) *Point {
	return &Point{
		symbol,
		x,
		y,
		direction,
	}
}

//func (p *Point) nextCoords() *Point {
//	switch p.direction {
//	case N:
//		return BuildPoint(p.symbol, p.x, p.y-1, p.symbol)
//
//	}
//}

func (p *Point) next(symbol rune) *Point {
	switch symbol {
	case 1:

	}
	return p
}
