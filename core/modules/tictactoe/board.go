package tictactoe

import "math"

type Board struct {
	board [Size][Size]int
}

func (b *Board) Init() {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			b.Set(i, j, 0)
		}
	}
}

func (b *Board) FromMatrix(board [Size][Size]int) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			b.Set(i, j, board[i][j])
		}
	}
}

func (b *Board) ToMatrix() [Size][Size]int {
	return b.board
}

func (b *Board) Get(row int, col int) int {
	return b.board[row][col]
}

func (b *Board) Set(row int, col int, value int) {
	b.board[row][col] = value
}

func (b *Board) CheckWinner() GameStatus {
	for i := 0; i < Size; i++ {
		if b.Get(i, 0) != 0 && b.Get(i, 0) == b.Get(i, 1) && b.Get(i, 1) == b.Get(i, 2) {
			return GameStatus{Winner: b.Get(i, 1), Coords: []Coord{{i, 0}, {i, 1}, {i, 2}}}
		}

		if b.Get(0, i) != 0 && b.Get(0, i) == b.Get(1, i) && b.Get(1, i) == b.Get(2, i) {
			return GameStatus{Winner: b.Get(1, i), Coords: []Coord{{0, i}, {1, i}, {2, i}}}
		}
	}

	if b.Get(1, 1) != 0 {
		if b.Get(0, 0) == b.Get(1, 1) && b.Get(1, 1) == b.Get(2, 2) {
			return GameStatus{Winner: b.Get(1, 1), Coords: []Coord{{0, 0}, {1, 1}, {2, 2}}}
		}

		if b.Get(0, 2) == b.Get(1, 1) && b.Get(1, 1) == b.Get(2, 0) {
			return GameStatus{Winner: b.Get(1, 1), Coords: []Coord{{0, 2}, {1, 1}, {2, 0}}}
		}
	}

	return GameStatus{Winner: 0}
}

func (b *Board) IsDone() bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if b.Get(i, j) == 0 {
				return false
			}
		}
	}

	return true
}

func (b *Board) Hash() int {
	var hash int = 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			var idx int = i*Size + j
			var val int = b.Get(i, j)
			if val == -1 {
				val = 2
			}

			hash += val * int(math.Pow(3, float64(idx)))
		}
	}

	return hash
}

func (b *Board) GetOpenSpots() []Coord {
	var spots []Coord
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if b.Get(i, j) == 0 {
				spots = append(spots, Coord{i, j})
			}
		}
	}

	return spots
}
