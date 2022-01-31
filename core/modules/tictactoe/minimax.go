package tictactoe

import "math/rand"

const MaxInt int8 = 100
const MinInt int8 = -100

func CalcMove(board Board) Response {
	dp := newdp()
	return dp.miniMax(board, 0, false, MinInt, MaxInt)
}

func newdp() *dp {
	return &dp{Memo: make(map[int]Response)}
}

func (d *dp) miniMax(board Board, depth int, isMax bool, alpha int8, beta int8) Response {
	var hash int = board.Hash()
	if _, ok := d.Memo[hash]; ok {
		return d.Memo[hash]
	}

	var winner int = board.CheckWinner().Winner
	if winner != 0 {
		var endEval int = winner*10 + (10-depth)*winner
		return Response{Value: int8(endEval)}
	} else if board.IsDone() {
		return Response{Value: 0}
	}

	var moveOptions []Coord = board.GetOpenSpots()
	var bestVal int8 = MinInt
	var swapThreshold float64 = 0.5
	var bestMove Coord

	if isMax {
		for _, option := range moveOptions {
			board.Set(option.Row, option.Col, 1)
			current := d.miniMax(board, depth+1, false, alpha, beta)
			board.Set(option.Row, option.Col, 0)

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

		res := Response{Row: int8(bestMove.Row), Col: int8(bestMove.Col), Value: bestVal}
		d.Memo[hash] = res
		return res
	}

	bestVal = MaxInt
	for _, option := range moveOptions {
		board.Set(option.Row, option.Col, -1)
		current := d.miniMax(board, depth+1, true, alpha, beta)
		board.Set(option.Row, option.Col, 0)

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

	res := Response{Row: int8(bestMove.Row), Col: int8(bestMove.Col), Value: bestVal}
	d.Memo[hash] = res
	return res
}
