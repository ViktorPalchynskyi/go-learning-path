package exercise2

type Point struct{ X, Y float64 }

//go:noinline
func newPointValue(x, y float64) Point {
	return Point{X: x, Y: y}
}

//go:noinline
func newPointPtr(x, y float64) *Point {
	return &Point{X: x, Y: y}
}
