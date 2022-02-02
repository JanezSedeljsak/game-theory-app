package connect4

import "fmt"

const MaxScore int8 = 100
const MinScore int8 = -100

// Prioritize moves in the center
var ExploreOrder = [...]int8{3, 4, 2, 5, 1, 6, 0}

func CalcMove(board Board) MiniMaxState {
	return iterativeDeepening(board)
}

func newdp() *dp {
	return &dp{Memo: make(map[uint64]int8)}
}

func iterativeDeepening(board Board) MiniMaxState {
	dp := newdp()
	curBest := MaxScore
	var curRes MiniMaxState

	for i := 3; i < 15; i++ {
		res := dp.negaMax(board, 0, -1, i, MinScore, MaxScore)
		if res.Value < curBest {
			curBest = res.Value
			curRes = res
		}
	}

	return curRes
}

func (dp *dp) negaMax(board Board, depth int, color int, maxDepth int, alpha int8, beta int8) MiniMaxState {
	hash := board.Hash()
	if _, ok := dp.Memo[hash]; ok {
		return MiniMaxState{Value: dp.Memo[hash]}
	}

	var winner int = board.CheckWinner().Winner
	if winner != 0 {
		var endEval int = winner*50 + (50-depth)*winner
		fmt.Println(endEval * color)
		return MiniMaxState{Value: int8(endEval * color)}
	} else if depth == maxDepth {
		return MiniMaxState{Value: 0}
	}

	var bestVal int8 = MinScore
	var foundOption bool = false
	var bestMove int8

	for _, option := range ExploreOrder {
		if board.Cols[option].IsFull() {
			continue
		}

		foundOption = true
		board.Drop(int(option), color)
		newVal := -dp.negaMax(board, depth+1, -color, maxDepth, -beta, -alpha).Value
		board.Pop()

		if newVal > bestVal {
			bestVal = newVal
			bestMove = option
		}

		if newVal > alpha {
			alpha = newVal
		}

		if alpha >= beta {
			break
		}
	}

	if !foundOption {
		return MiniMaxState{Value: 0}
	}

	res := MiniMaxState{Col: bestMove, Value: bestVal * int8(color)}
	dp.Memo[hash] = res.Value
	return res
}
