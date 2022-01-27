package core

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = int(-MaxInt - 1)
const Size = 3

func MinMax(board [][]int, depth int, isMax bool, alpha int, beta int) Response {
	var winner int = checkWinner(board)
	if winner != 0 {
		return Response{Coord{}, winner}
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
	if isMax {
		for _, option := range moveOptions {

			// update and reset board after minmax call
			board[option.x][option.y] = 1
			current := MinMax(board, depth+1, false, alpha, beta)
			board[option.x][option.y] = 0

			bestVal = max(bestVal, current.value)
			alpha = max(alpha, bestVal)
			if beta <= alpha {
				break
			}
		}

		return Response{Coord{}, bestVal}
	}

	bestVal = MaxInt
	for _, option := range moveOptions {

		// update and reset board after minmax call
		board[option.x][option.y] = -1
		current := MinMax(board, depth+1, true, alpha, beta)
		board[option.x][option.y] = 0

		bestVal = min(bestVal, current.value)
		beta = min(beta, bestVal)
		if beta <= alpha {
			break
		}
	}

	return Response{Coord{}, bestVal}
}

func min(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}
