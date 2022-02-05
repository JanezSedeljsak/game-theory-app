package connect4

const MaxScore int8 = 100
const MinScore int8 = -100

// Prioritize moves in the center
var ExploreOrder = [Width]int8{3, 4, 2, 5, 1, 6, 0}

func CalcMove(board BitmapBoard, maxDepth int8) MiniMaxState {
	return newdp(maxDepth).negaMax(board, 0, -1, MinScore, MaxScore, false)
}

func newdp(maxDepth int8) *dp {
	return &dp{Memo: make(map[uint64]int8), MaxDepth: maxDepth}
}

func (dp *dp) negaMax(board BitmapBoard, depth int8, color int8, alpha int8, beta int8, winner bool) MiniMaxState {
	if winner {
		var weight int8 = 1
		if depth%2 == 0 {
			weight = -1
		}
		var endEval int8 = (50 * weight) - (depth * weight)
		return MiniMaxState{Value: endEval}
	} else if depth == dp.MaxDepth {
		return MiniMaxState{Value: 0}
	}

	hash := board.Hash()
	if _, ok := dp.Memo[hash]; ok {
		return MiniMaxState{Value: dp.Memo[hash]}
	}

	var bestVal int8 = MinScore * color
	var bestMove int8

	moves := board.SortedMoves(hash)
	if moves == nil {
		return MiniMaxState{Value: 0}
	}

	for _, move := range moves {
		newVal := dp.negaMax(move.Board, depth+1, -color, alpha, beta, move.Winner).Value

		if color == 1 && newVal > bestVal {
			bestVal = newVal
			bestMove = move.Col
			if bestVal > alpha {
				alpha = newVal
			}
		} else if color == -1 && newVal < bestVal {
			bestVal = newVal
			bestMove = move.Col
			if bestVal < beta {
				beta = newVal
			}
		}

		if alpha >= beta {
			break
		}
	}

	res := MiniMaxState{Col: bestMove, Value: bestVal}
	if bestVal != 0 {
		dp.Memo[hash] = res.Value
	}

	return res
}
