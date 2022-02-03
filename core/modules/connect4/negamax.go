package connect4

const MaxScore int8 = 100
const MinScore int8 = -100

// Prioritize moves in the center
var ExploreOrder = [...]int8{3, 4, 2, 5, 1, 6, 0}

func CalcMove(board Board) MiniMaxState {
	board.CalcHash()
	return newdp().iterativeDeepening(board)
}

func newdp() *dp {
	return &dp{Memo: make(map[uint64]int8)}
}

func (dp *dp) iterativeDeepening(board Board) MiniMaxState {
	best := MiniMaxState{Value: MaxScore}

	for depth := 5; depth < 15; depth++ {
		dp.MaxDepth = depth
		curRes := dp.negaMax(board, 0, -1, MinScore, MaxScore)
		if curRes.Value < best.Value {
			best = curRes
		}

		if best.Value < 0 {
			break
		}
	}

	return best
}

func (dp *dp) negaMax(board Board, depth int, color int, alpha int8, beta int8) MiniMaxState {
	if _, ok := dp.Memo[board.hashValue]; ok {
		return MiniMaxState{Value: dp.Memo[board.hashValue]}
	}

	var winner int = board.CheckWinner().Winner
	if winner != 0 {
		endEval := (50 * winner) - (depth * winner)
		return MiniMaxState{Value: int8(endEval)}
	} else if depth == dp.MaxDepth {
		return MiniMaxState{Value: 0}
	}

	var bestVal int8 = MinScore * int8(color)
	var foundOption bool = false
	var bestMove int8

	for _, option := range ExploreOrder {
		if board.Cols[option].IsFull() {
			continue
		}

		foundOption = true
		board.Drop(int(option), color)
		newVal := dp.negaMax(board, depth+1, -color, alpha, beta).Value
		board.Pop()

		if color == 1 && newVal > bestVal {
			bestVal = newVal
			bestMove = option
			if bestVal > alpha {
				alpha = newVal
			}
		} else if color == -1 && newVal < bestVal {
			bestVal = newVal
			bestMove = option
			if bestVal < beta {
				beta = newVal
			}
		}

		if alpha >= beta {
			break
		}
	}

	if !foundOption {
		return MiniMaxState{Value: 0}
	}

	res := MiniMaxState{Col: bestMove, Value: bestVal}
	if bestVal != 0 {
		dp.Memo[board.hashValue] = res.Value
	}

	return res
}
