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
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			b.Cols[j].Push(board[i][j])
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

func (b *Board) GetOpenSpots() []int {
	var options []int
	for i, col := range b.Cols {
		if !col.IsFull() {
			options = append(options, i)
		}
	}

	return options
}
