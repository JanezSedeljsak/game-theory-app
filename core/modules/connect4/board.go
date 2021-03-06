package connect4

type Board struct {
	Cols         [Width]Stack
	lastInserted int
}

func (b *Board) Init() bool {
	b.lastInserted = -1
	for i := 0; i < Width; i++ {
		b.Cols[i] = *newIntStack(Height)
	}

	return true
}

func (b *Board) SetLastInserted(col int) {
	b.lastInserted = col
}

func (b *Board) ToMatrix() [Height][Width]int {
	var board [Height][Width]int
	for i, col := range b.Cols {
		for j := 0; j < Height; j++ {
			board[j][i] = col.Peek(j).(int)
		}
	}

	return board
}

func (b *Board) ToBitmap() BitmapBoard {
	var bmap BitmapBoard
	bmap.Init()

	for i, col := range b.Cols {
		for j := 0; j < col.Count(); j++ {
			cell := col.Peek(j).(int)
			idx := i*7 + j

			if cell != 0 {
				bmap.Mask |= 1 << idx
				if cell == 1 {
					bmap.Pos |= 1 << idx
				}
			}
		}
	}

	return bmap
}

func (b *Board) FromMatrix(board [Height][Width]int) {
	for j := 0; j < Width; j++ {
		b.Cols[j] = *newIntStack(Height)
		for i := 0; i < Height && board[i][j] != 0; i++ {
			b.Cols[j].Push(board[i][j])
		}
	}
}

func (b *Board) Drop(col int, player int) bool {
	if b.Cols[col].IsFull() {
		return false
	}

	b.Cols[col].Push(player)
	b.SetLastInserted(col)
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
	if row2 < 0 || row2 >= Height {
		return false
	}

	if col2 < 0 || col2 >= Width {
		return false
	}

	return b.Cols[col1].Peek(row1) == b.Cols[col2].Peek(row2)
}

func (b *Board) checkDirection(r int, c int, dr int, dc int) GameStatus {
	winningLine := []Coord{{r, c}}

	for i := 1; len(winningLine) != 4; i++ {
		if b.cmp(r, c, r+i*dr, c+i*dc) {
			winningLine = append(winningLine, Coord{r + i*dr, c + i*dc})
		} else {
			break
		}
	}

	if len(winningLine) < 4 {
		for i := -1; len(winningLine) != 4; i-- {
			if b.cmp(r, c, r+i*dr, c+i*dc) {
				winningLine = append(winningLine, Coord{r + i*dr, c + i*dc})
			} else {
				break
			}
		}
	}

	if len(winningLine) == 4 {
		return GameStatus{Winner: b.Cols[c].Peek(r).(int), Coords: winningLine}
	}

	return GameStatus{Winner: 0}
}

func (b *Board) CountMoves() int8 {
	sum := 0
	for _, col := range b.Cols {
		sum += col.Count()
	}

	return int8(sum)
}

func (b *Board) CheckWinner() GameStatus {
	c := b.lastInserted
	if c == -1 {
		return GameStatus{Winner: 0, Coords: []Coord{}}
	}

	r := b.Cols[c].top

	// Check vertical (down)
	if r > 2 && b.cmp(r, c, r-1, c) && b.cmp(r, c, r-2, c) && b.cmp(r, c, r-3, c) {
		return GameStatus{Winner: b.Cols[c].Peek(r).(int), Coords: []Coord{{r, c}, {r - 1, c}, {r - 2, c}, {r - 3, c}}}
	}

	// Check left diagonal
	gs := b.checkDirection(r, c, -1, -1)
	if gs.Winner != 0 {
		return gs
	}

	// Check right diagonal
	gs = b.checkDirection(r, c, -1, 1)
	if gs.Winner != 0 {
		return gs
	}

	// Check horizontal
	return b.checkDirection(r, c, 0, -1)
}
