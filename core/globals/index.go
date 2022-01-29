package globals

const MaxInt = 100
const MinInt = -100

type Coord struct {
	Row int
	Col int
}

type Response struct {
	Coords Coord
	Value  int
}
