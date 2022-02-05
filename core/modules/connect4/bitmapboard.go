package connect4

/*
 * How a connect4 - bitmap works
 * https:github.com/denkspuren/BitboardC4/blob/master/BitboardDesign.md
 */

const BOTTOM uint64 = 0b_0000001_0000001_0000001_0000001_0000001_0000001_0000001

type BitmapBoard struct {
	Pos  uint64
	Mask uint64
}

func (bb *BitmapBoard) Init() {
	bb.Pos = 0
	bb.Mask = 0
}

func (bb *BitmapBoard) GetPlayerBitmap(color int8) uint64 {
	if color == 1 {
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

func (bb *BitmapBoard) MakeMove(col int8, color int8) {
	newMask := bb.Mask | (bb.Mask + (1 << (col * 7)))

	if color == 1 {
		opponent := bb.Mask ^ bb.Pos
		bb.Pos = newMask ^ opponent
	}

	bb.Mask = newMask
}

func (bb *BitmapBoard) Hash() uint64 {
	return bb.Pos + bb.Mask + BOTTOM
}

func (bb *BitmapBoard) CheckWinner(color int8) int8 {
	bmap := bb.GetPlayerBitmap(color)
	var options = [...]int8{1, 6, 8, 7}
	var pos uint64

	for _, dir := range options {
		pos = bmap & (bmap >> dir)
		if pos&(pos>>(dir*2)) > 0 {
			return color
		}
	}

	return 0
}

func (bb *BitmapBoard) SortedMoves(hash uint64, color int8) []MoveEval {
	var moves [7]MoveEval
	isSymmetrical := IsSymmetrical(hash)
	validCount := 0

	for _, option := range ExploreOrder {
		if (isSymmetrical && option > 3) || !bb.CanPlay(option) {
			continue
		}

		tmpBoard := BitmapBoard{Pos: bb.Pos, Mask: bb.Mask}
		tmpBoard.MakeMove(option, color)
		winner := tmpBoard.CheckWinner(color)
		moves[validCount] = MoveEval{Col: option, Board: tmpBoard, Winner: winner}
		validCount++
	}

	if validCount == 0 {
		return nil
	}

	// sort moves based on value with insertion sort (DESC -> 1, ASC -> -1)
	var sortedMoves = make([]MoveEval, validCount)
	sortedMoves[0] = moves[0]

	for i := 1; i < validCount; i++ {
		sortedMoves[i] = moves[i]
		j := i

		for j > 0 {
			if moves[j-1].Winner*color < moves[j].Winner*color {
				sortedMoves[j-1], sortedMoves[j] = sortedMoves[j], sortedMoves[j-1]
			}

			j--
		}
	}

	return sortedMoves
}
