package tictactoe

import (
	"game-theory-app/core/globals"
	"math/rand"
)

func CalcMove(board [][]int8) globals.Response {
	dp := newdp()
	return dp.minMax(board, 0, false, globals.MinInt, globals.MaxInt)
}

func newdp() *dp {
	return &dp{Memo: make(map[int]globals.Response)}
}

func fullParamsHash(bHash int, isMax bool, alpha int, beta int) int {
	bHash = 31*bHash + alpha
	bHash = 31*bHash + beta
	bHash = 31 * bHash
	if isMax {
		bHash += 1
	}

	return bHash
}

func (d *dp) minMax(board [][]int8, depth int, isMax bool, alpha int, beta int) globals.Response {
	var bHash int = BoardHash(board)
	var hash = fullParamsHash(bHash, isMax, alpha, beta)
	if _, ok := d.Memo[hash]; ok {
		return d.Memo[hash]
	}

	var winner int8 = int8(CheckWinner(board).Winner)
	if winner != 0 {
		var endEval int = int(winner)*10 + (10-depth)*int(winner)
		return globals.Response{Coords: globals.Coord{}, Value: endEval}
	} else if IsStalemate(board) {
		return globals.Response{Coords: globals.Coord{}, Value: 0}
	}

	var moveOptions []globals.Coord = GetOpenSpots(board)
	var bestVal int = globals.MinInt
	var swapThreshold float64 = 0.5
	var bestMove globals.Coord

	if isMax {
		for _, option := range moveOptions {
			board[option.Row][option.Col] = 1
			current := d.minMax(board, depth+1, false, alpha, beta)
			board[option.Row][option.Col] = 0

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

		res := globals.Response{Coords: bestMove, Value: bestVal}
		d.Memo[hash] = res
		return res
	}

	bestVal = globals.MaxInt
	for _, option := range moveOptions {
		board[option.Row][option.Col] = -1
		current := d.minMax(board, depth+1, true, alpha, beta)
		board[option.Row][option.Col] = 0

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

	res := globals.Response{Coords: bestMove, Value: bestVal}
	d.Memo[hash] = res
	return res
}
