package connect4

type Board struct {
	Cols [Width]Stack
}

func (b *Board) Init() bool {
	for i := 0; i < Width; i++ {
		b.Cols[i].Init(Height)
	}

	return true
}

func (b *Board) ToMatrix() [Height][Width]int {
	var board [Height][Width]int
	for i, col := range b.Cols {
		for j := 0; j < Height; j++ {
			board[j][i] = col.Peek(j)
		}
	}

	return board
}

func (b *Board) FromMatrix(board [Height][Width]int) {
	for i, col := range b.Cols {
		for j := 0; j < Height; j++ {
			col.Push(board[j][i])
		}
	}
}

func (b *Board) Drop(i int, player int) bool {
	if b.Cols[i].IsFull() {
		return false
	}

	b.Cols[i].Push(player)
	return true
}
