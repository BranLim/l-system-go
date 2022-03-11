package painter

type Point struct {
	Symbol    rune
	X         int
	Y         int
	Direction int
}

const (
	N int = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func BuildPoint(symbol rune, x, y, direction int) *Point {
	return &Point{
		symbol,
		x,
		y,
		direction,
	}
}

func (p *Point) AsRune() rune {
	if p.Symbol == '0' {
		return 'o'
	}
	switch p.Direction {
	case N, S:
		return '|'
	case SE, NW:
		return '\\'
	case E, W:
		return '-'
	case NE, SW:
		return '/'
	}
	return ' '
}

func (p *Point) nextCoords() *Point {
	var x = p.X
	var y = p.Y

	if p.Direction == N || p.Direction == NW || p.Direction == NE {
		y--
	} else if p.Direction == S || p.Direction == SE || p.Direction == SW {
		y++
	}

	if p.Direction == W || p.Direction == SW || p.Direction == NW {
		x--
	} else if p.Direction == NE || p.Direction == SE || p.Direction == E {
		x++
	}
	return &Point{p.Symbol, x, y, p.Direction}
}

func (p *Point) rotateLeft() *Point {
	direction := p.Direction - 1
	x := p.X
	y := p.Y
	if direction < 0 {
		direction += 8
	}
	if p.Direction >= S || p.Direction == N {
		x--
	} else {
		x++
	}
	return BuildPoint(p.Symbol, x, y, direction)
}

func (p *Point) rotateRight() *Point {
	var direction = p.Direction + 1
	direction = direction % 8
	x := p.X
	if p.Direction >= S {
		x--
	} else {
		x++
	}
	return BuildPoint(p.Symbol, x, p.Y, direction)
}

func (p *Point) Next(symbol rune) *Point {
	switch symbol {
	case '1':
		return p.nextCoords()
	case '0':
		return BuildPoint(symbol, p.X, p.Y, p.Direction)
	case '[':
		return p.rotateLeft()
	case ']':
		return p.rotateRight()
	default:
		return p
	}
}
