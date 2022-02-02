package testing

import (
	"game-theory-app/core/modules/connect4"
	"math/rand"
)

const MaxInt int8 = 100
const MinInt int8 = -100

func (t *Test) CalcMove(board connect4.Board) connect4.MiniMaxState {
	res := t.miniMax(board, t.depth, false, MinInt, MaxInt)
	return res
}

func (t *Test) miniMax(board connect4.Board, depth int, isMax bool, alpha int8, beta int8) connect4.MiniMaxState {
	var hash uint64
	if t.isDp {
		hash = board.Hash()
		if _, ok := t.dp.Memo[hash]; ok {
			return t.dp.Memo[hash]
		}
	}

	var winner int = board.CheckWinner().Winner
	if winner != 0 {
		var endEval int = winner*50 + (50-depth)*winner
		return connect4.MiniMaxState{Value: int8(endEval)}
	} else if depth == 8 || board.IsDone() {
		return connect4.MiniMaxState{Value: 0}
	}

	var moveOptions []int = board.GetOpenSpots()
	var bestVal int8 = MinInt
	var swapThreshold float64 = 0.5
	var bestMove int

	if isMax {
		for _, option := range moveOptions {
			board.Drop(option, 1)
			current := t.miniMax(board, depth+1, false, alpha, beta)
			board.Pop(option)

			if current.Value > bestVal {
				swapThreshold = 0.5
				bestVal = current.Value
				bestMove = option
			}

			// Pick one of many best moves at random
			if current.Value == bestVal && rand.Float64() > swapThreshold {
				swapThreshold += (1 - swapThreshold) / 2
				bestMove = option
			}

			if bestVal > alpha {
				alpha = bestVal
			}

			if t.isAlphaBeta && beta < alpha {
				break
			}
		}

		res := connect4.MiniMaxState{Col: int8(bestMove), Value: bestVal}
		if t.isDp {
			t.dp.Memo[hash] = res
		}

		return res
	}

	bestVal = MaxInt
	for _, option := range moveOptions {
		board.Drop(option, -1)
		current := t.miniMax(board, depth+1, true, alpha, beta)
		board.Pop(option)

		if current.Value < bestVal {
			swapThreshold = 0.5
			bestVal = current.Value
			bestMove = option
		}

		// Pick one of many best moves at random
		if current.Value == bestVal && rand.Float64() > swapThreshold {
			swapThreshold += (1 - swapThreshold) / 2
			bestMove = option
		}

		if bestVal < beta {
			beta = bestVal
		}

		if t.isAlphaBeta && beta < alpha {
			break
		}
	}

	res := connect4.MiniMaxState{Col: int8(bestMove), Value: bestVal}
	if t.isDp {
		t.dp.Memo[hash] = res
	}

	return res
}
