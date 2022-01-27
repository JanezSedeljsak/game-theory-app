package core

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = int(-MaxInt - 1)
const Size = 3

func MinMax(board [][]int, depth int, isMax bool, alpha int, beta int) Response {
	var winner int = CheckWinner(board)
	if winner != 0 {
		return Response{Coord{}, winner}
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
			current := MinMax(board, depth+1, false, alpha, beta)
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

		return Response{bestMove, bestVal}
	}

	bestVal = MaxInt
	for _, option := range moveOptions {
		board[option.Row][option.Col] = -1
		current := MinMax(board, depth+1, true, alpha, beta)
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

	return Response{bestMove, bestVal}
}
