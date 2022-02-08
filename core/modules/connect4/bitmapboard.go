package connect4

/*
 * How a connect4 - bitmap works
 * https:github.com/denkspuren/BitboardC4/blob/master/BitboardDesign.md
 */

const BOTTOM uint64 = 0b_0000001_0000001_0000001_0000001_0000001_0000001_0000001
const THIRD uint64 = 0b_1111111_0000000_0000000
const SECOND uint64 = 0b_0000000_1111111_0000000
const FIRST uint64 = 0b_0000000_0000000_1111111

type BitmapBoard struct {
	Pos  uint64
	Mask uint64
}

func (bb *BitmapBoard) Init() {
	bb.Pos = 0
	bb.Mask = 0
}

func (bb *BitmapBoard) GetPlayerBitmap(color int8) uint64 {
	return Ternary(color == 1, bb.Pos, bb.Pos^bb.Mask).(uint64)
}

func (bb *BitmapBoard) CanPlay(col int8) bool {
	return (bb.Mask>>(col*7+5))&1 == 0
}

func (bb *BitmapBoard) ToMatrix(weight int) [Height][Width]int {
	var board [Height][Width]int
	opponent := bb.GetPlayerBitmap(-1)
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			idx := i*Width + j

			if (bb.Pos>>idx)&1 == 1 {
				board[j][i] = -weight
			} else if (opponent>>idx)&1 == 1 {
				board[j][i] = weight
			}
		}
	}

	return board
}

func (bb *BitmapBoard) MakeMove(col int8) {
	bb.Mask |= bb.Mask + (1 << (col * 7))
	bb.Pos ^= bb.Mask
}

func (bb *BitmapBoard) Hash() uint64 {
	return bb.Pos + bb.Mask + BOTTOM
}

func (bb *BitmapBoard) CheckWinner() bool {
	var options = [4]int8{1, 6, 8, 7}
	var pos uint64

	for _, dir := range options {
		pos = bb.Pos & (bb.Pos >> dir)
		if pos&(pos>>(dir*2)) > 0 {
			return true
		}
	}

	return false
}

/*
 @result => tuple (code - int, all moves - [7]array, size - amount of valid moves)
 @code => 0: ok, 1: empty, 2: forced move
*/
func (bb *BitmapBoard) SortedMoves(hash uint64) (int8, [7]MoveEval, int) {
	var moves [7]MoveEval
	var validCount int = 0
	isSymmetrical := IsSymmetrical(hash)

	// fill table with valid moves
	for _, option := range ExploreOrder {
		if (isSymmetrical && option > 3) || !bb.CanPlay(option) {
			continue
		}

		tmpBoard := BitmapBoard{Pos: bb.Pos, Mask: bb.Mask}
		tmpBoard.MakeMove(option)
		winner := tmpBoard.CheckWinner()
		if winner {
			return 2, [7]MoveEval{{Col: option, Board: tmpBoard, Winner: winner}}, 0 // forced move
		}

		moves[validCount] = MoveEval{Col: option, Board: tmpBoard, Winner: winner}
		validCount++
	}

	if validCount == 0 {
		return 1, [7]MoveEval{}, 0
	}

	return 0, moves, validCount
}
