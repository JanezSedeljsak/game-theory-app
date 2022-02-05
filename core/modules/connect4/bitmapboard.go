package connect4

/*
 * How a connect4 - bitmap workshttps:
 * github.com/denkspuren/BitboardC4/blob/master/BitboardDesign.md
 */

const BOTTOM uint64 = 0b_0000001_0000001_0000001_0000001_0000001_0000001_0000001

// each cell indicates how many winning positions can be made with that specific cell
var EvaluationTable = [Height][Width]int8{
	{3, 4, 5, 7, 5, 4, 3},
	{4, 6, 8, 10, 8, 6, 4},
	{5, 8, 11, 13, 11, 8, 5},
	{5, 8, 11, 13, 11, 8, 5},
	{4, 6, 8, 10, 8, 6, 4},
	{3, 4, 5, 7, 5, 4, 3},
}

type BitmapBoard struct {
	Pos        uint64
	Mask       uint64
	lastPlayer int8
	lastCol    int8
}

func (bb *BitmapBoard) Init() {
	bb.Pos = 0
	bb.Mask = 0
	bb.lastCol = -1
	bb.lastPlayer = 0
}

func (bb *BitmapBoard) GetPlayerBitmap(player int8) uint64 {
	if player == 1 {
		return bb.Pos
	}

	return bb.Pos ^ bb.Mask
}

func (bb *BitmapBoard) Evaluate() int8 {
	opponent := bb.GetPlayerBitmap(-1)
	sum := 0

	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			idx := i*Width + j

			if (bb.Pos>>idx)&1 == 1 {
				sum += int(EvaluationTable[i][j])
			} else if (opponent>>idx)&1 == 1 {
				sum -= int(EvaluationTable[i][j])
			}
		}
	}

	return int8(sum / 3)
}

func (bb *BitmapBoard) CanPlay(col int8) bool {
	return (bb.Mask>>(col*7+5))&1 == 0
}

func (bb *BitmapBoard) ToMatrix() [Height][Width]int {
	var board [Height][Width]int
	opponent := bb.GetPlayerBitmap(-1)
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			idx := i*Width + j

			if (bb.Pos>>idx)&1 == 1 {
				board[j][i] = 1
			} else if (opponent>>idx)&1 == 1 {
				board[j][i] = -1
			}
		}
	}

	return board
}

func (bb *BitmapBoard) MakeMove(col int8, player int8) {
	newMask := bb.Mask | (bb.Mask + (1 << (col * 7)))
	bb.lastPlayer = player
	bb.lastCol = col

	if player == 1 {
		opponent := bb.Mask ^ bb.Pos
		bb.Pos = newMask ^ opponent
	}

	bb.Mask = newMask
}

func (bb *BitmapBoard) ReverseMove() {
	bb.Mask ^= (bb.Mask - (1 << (bb.lastCol * 7)))
	if bb.lastPlayer == 1 {
		bb.Pos &= bb.Mask
	}
}

func (bb *BitmapBoard) Hash() uint64 {
	return bb.Pos + bb.Mask + BOTTOM
}

func (bb *BitmapBoard) CheckWinner() int8 {
	bmap := bb.GetPlayerBitmap(bb.lastPlayer)

	// horizontal
	var pos uint64 = bmap & (bmap >> 7)
	if pos&(pos>>14) > 0 {
		return bb.lastPlayer
	}

	// \ diagonal
	pos = bmap & (bmap >> 6)
	if pos&(pos>>12) > 0 {
		return bb.lastPlayer
	}

	// / diagonal
	pos = bmap & (bmap >> 8)
	if pos&(pos>>16) > 0 {
		return bb.lastPlayer
	}

	// vertical
	pos = bmap & (bmap >> 1)
	if pos&(pos>>2) > 0 {
		return bb.lastPlayer
	}

	return 0
}
