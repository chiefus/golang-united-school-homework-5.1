package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square *Square) Perimeter() float64 {
	return float64(square.a * square.a)
}

func (square *Square) Area() float64 {
	return float64(square.a * 4)
}

func (square *Square) End() Point {
	endX := square.start.x + int(square.a)
	endY := square.start.y - int(square.a)
	return Point{x: endX, y: endY}
}
