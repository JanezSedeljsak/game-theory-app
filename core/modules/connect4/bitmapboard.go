package connect4

/*
 * How a connect4 - bitmap workshttps:
 * github.com/denkspuren/BitboardC4/blob/master/BitboardDesign.md
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

func (bb *BitmapBoard) HeuristicEvaluation() int8 {
	player2 := bb.Pos ^ bb.Mask
	player2Options := ^bb.Pos
	player1Options := ^player2

	player1Wins := CountWinnablePositions(player1Options)
	player2Wins := CountWinnablePositions(player2Options)

	var hScore int8 = 0
	for i := 0; i < 4; i++ {
		hScore += PopulationCount(player1Wins[i]) - PopulationCount(player2Wins[i])
	}

	return hScore
}

func CountWinnablePositions(bmap uint64) [4]uint64 {
	var options = [...]int8{7, 6, 8, 1}
	var pos uint64
	var countDir [4]uint64

	for idx, dir := range options {
		pos = bmap & (bmap >> dir)
		countDir[idx] = pos & (pos >> (dir * 2))
	}

	return countDir
}

// population count -> count 1s in binary representation of a bitmap
func PopulationCount(bmap uint64) int8 {
	var count int8
	for count = 0; bmap > 0; count++ {
		bmap &= bmap - 1
	}

	return count
}
