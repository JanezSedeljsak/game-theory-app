package core

type Coord struct {
	Row int
	Col int
}

type Response struct {
	Coords Coord
	Value  int
}

type dp struct {
	Memo map[int]Response
}
