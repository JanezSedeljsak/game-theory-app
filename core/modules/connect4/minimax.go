package connect4

import "math/rand"

const MaxInt int8 = 100
const MinInt int8 = -100

func CalcMove(board Board) MiniMaxState {
	dp := newdp()
	return dp.miniMax(board, 0, false, MinInt, MaxInt)
}

func newdp() *dp {
	return &dp{Memo: make(map[uint64]MiniMaxState)}
}

func (d *dp) miniMax(board Board, depth int, isMax bool, alpha int8, beta int8) MiniMaxState {
	var hash uint64 = board.Hash()
	if _, ok := d.Memo[hash]; ok {
		return d.Memo[hash]
	}

	var winner int = board.CheckWinner().Winner
	if winner != 0 {
		var endEval int = winner*50 + (50-depth)*winner
		return MiniMaxState{Value: int8(endEval)}
	} else if board.IsDone() || depth == 8 {
		return MiniMaxState{Value: 0}
	}

	var moveOptions []int = board.GetOpenSpots()
	var bestVal int8 = MinInt
	var swapThreshold float64 = 0.5
	var bestMove int

	if isMax {
		for _, option := range moveOptions {
			board.Drop(option, 1)
			current := d.miniMax(board, depth+1, false, alpha, beta)
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

			if beta < alpha {
				break
			}
		}

		res := MiniMaxState{Col: int8(bestMove), Value: bestVal}
		d.Memo[hash] = res
		return res
	}

	bestVal = MaxInt
	for _, option := range moveOptions {
		board.Drop(option, -1)
		current := d.miniMax(board, depth+1, true, alpha, beta)
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

		if beta < alpha {
			break
		}
	}

	res := MiniMaxState{Col: int8(bestMove), Value: bestVal}
	d.Memo[hash] = res
	return res
}
