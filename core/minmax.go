package core

const MaxInt = 100
const MinInt = -100
const Size = 3

func CalcMove(board [][]int8) Response {
	dp := newdp()
	return dp.minMax(board, 0, false, MinInt, MaxInt)
}

func newdp() *dp {
	return &dp{Memo: make(map[int]Response)}
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

func (d *dp) minMax(board [][]int8, depth int, isMax bool, alpha int, beta int) Response {
	var bHash int = BoardHash(board)
	var hash = fullParamsHash(bHash, isMax, alpha, beta)
	if _, ok := d.Memo[hash]; ok {
		return d.Memo[hash]
	}

	var winner int8 = CheckWinner(board)
	if winner != 0 {
		return Response{Coord{}, int(winner)}
	} else if IsStalemate(board) {
		return Response{Coord{}, 0}
	}

	var moveOptions []Coord
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if board[i][j] == 0 {
				moveOptions = append(moveOptions, Coord{i, j})
			}
		}
	}

	var bestVal int = MinInt
	var bestMove Coord

	if isMax {
		for _, option := range moveOptions {
			board[option.Row][option.Col] = 1
			current := d.minMax(board, depth+1, false, alpha, beta)
			board[option.Row][option.Col] = 0

			if current.Value > bestVal {
				bestVal = current.Value
				bestMove = option
			}

			if bestVal > alpha {
				alpha = bestVal
			}

			if beta <= alpha {
				break
			}
		}

		res := Response{bestMove, bestVal}
		d.Memo[hash] = res
		return res
	}

	bestVal = MaxInt
	for _, option := range moveOptions {
		board[option.Row][option.Col] = -1
		current := d.minMax(board, depth+1, true, alpha, beta)
		board[option.Row][option.Col] = 0

		if current.Value < bestVal {
			bestVal = current.Value
			bestMove = option
		}

		if bestVal < beta {
			beta = bestVal
		}

		if beta <= alpha {
			break
		}
	}

	res := Response{bestMove, bestVal}
	d.Memo[hash] = res
	return res
}
