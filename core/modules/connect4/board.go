package connect4

type Board struct {
	Cols         [Width]Stack
	lastInserted Coord
}

func (b *Board) Init() bool {
	for i := 0; i < Width; i++ {
		b.Cols[i].Init(Height)
	}

	return true
}

func (b *Board) SetLastInserted(row int, col int) {
	b.lastInserted = Coord{Row: row, Col: col}
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
	for j := 0; j < Width; j++ {
		b.Cols[j].Init(Height)
		for i := 0; i < Height && board[i][j] != 0; i++ {
			b.Cols[j].Push(board[i][j])
		}
	}
}

func (b *Board) Drop(col int, player int) bool {
	if b.Cols[col].IsFull() {
		return false
	}

	row := b.Cols[col].Push(player)
	b.SetLastInserted(row, col)
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

func (b *Board) IsDone() bool {
	for _, col := range b.Cols {
		if !col.IsFull() {
			return false
		}
	}

	return true
}

func (b *Board) cmp(row1 int, col1 int, row2 int, col2 int) bool {
	return b.Cols[col1].Peek(row1) == b.Cols[col2].Peek(row2)
}

func (b *Board) CheckWinner() GameStatus {
	r := b.lastInserted.Row
	c := b.lastInserted.Col

	if r > 2 {
		// Check vertical (down)
		if b.cmp(r, c, r-1, c) && b.cmp(r, c, r-2, c) && b.cmp(r, c, r-3, c) {
			return GameStatus{Winner: b.Cols[c].Peek(r), Coords: []Coord{{r, c}, {r - 1, c}, {r - 2, c}, {r - 3, c}}}
		}

		// Check left diagonal
		if c > 2 && b.cmp(r, c, r-1, c-1) && b.cmp(r, c, r-2, c-2) && b.cmp(r, c, r-3, c-3) {
			return GameStatus{Winner: b.Cols[c].Peek(r), Coords: []Coord{{r, c}, {r - 1, c - 1}, {r - 2, c - 2}, {r - 3, c - 3}}}
		}

		// Check right diagonal
		if c < 4 && b.cmp(r, c, r-1, c+1) && b.cmp(r, c, r-2, c+2) && b.cmp(r, c, r-3, c+3) {
			return GameStatus{Winner: b.Cols[c].Peek(r), Coords: []Coord{{r, c}, {r - 1, c + 1}, {r - 2, c + 2}, {r - 3, c + 3}}}
		}
	}

	// check horizontal
	winningLine := []Coord{{r, c}}
	for i := c - 1; i >= 0; i-- {
		if len(winningLine) == 4 {
			break
		}

		if b.cmp(r, c, r, i) {
			winningLine = append(winningLine, Coord{r, i})
		} else {
			break
		}
	}

	curLen := len(winningLine)
	if curLen < 4 && curLen+Width-c-1 >= 4 {
		for i := c + 1; i < Width; i++ {
			if len(winningLine) == 4 {
				break
			}

			if b.cmp(r, c, r, i) {
				winningLine = append(winningLine, Coord{r, i})

			} else {
				break
			}
		}
	}

	if len(winningLine) == 4 {
		return GameStatus{Winner: b.Cols[c].Peek(r), Coords: winningLine}
	}

	return GameStatus{Winner: 0}
}
