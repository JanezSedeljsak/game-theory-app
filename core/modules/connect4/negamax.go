package connect4

const MaxScore int8 = 100
const MinScore int8 = -100

// Prioritize moves in the center
var ExploreOrder = [...]int8{3, 4, 2, 5, 1, 6, 0}

func CalcMove(board Board) MiniMaxState {
	dp := newdp()
	return dp.negaMax(board, 0, -1)
}

func newdp() *dp {
	return &dp{Memo: make(map[uint64]int8)}
}

func (dp *dp) negaMax(board Board, depth int, color int) MiniMaxState {
	var winner int = board.CheckWinner().Winner
	if winner != 0 {
		var endEval int = winner*50 + (50-depth)*winner
		return MiniMaxState{Value: int8(endEval)}
	} else if depth == 7 {
		return MiniMaxState{Value: 0}
	}

	var bestVal int8 = MinScore
	var bestMove int8

	for _, option := range ExploreOrder {
		if board.Cols[option].IsFull() {
			continue
		}

		board.Drop(int(option), color)
		newVal := -dp.negaMax(board, depth+1, -color).Value
		board.Pop()

		if newVal > bestVal {
			bestVal = newVal
			bestMove = option
		}
	}

	res := MiniMaxState{Col: bestMove, Value: bestVal * int8(color)}
	return res
}
