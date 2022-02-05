package connect4

/*
 * How a connect4 - bitmap works
 * https:github.com/denkspuren/BitboardC4/blob/master/BitboardDesign.md
 */

const BOTTOM uint64 = 0b_0000001_0000001_0000001_0000001_0000001_0000001_0000001

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

func (bb *BitmapBoard) Hash() uint64 {
	return bb.Pos + bb.Mask + BOTTOM
}

func (bb *BitmapBoard) CheckWinner() int8 {
	bmap := bb.GetPlayerBitmap(bb.lastPlayer)
	var options = [...]int8{7, 6, 8, 1}
	var pos uint64

	for _, dir := range options {
		pos = bmap & (bmap >> dir)
		if pos&(pos>>(dir*2)) > 0 {
			return bb.lastPlayer
		}
	}

	return 0
}
