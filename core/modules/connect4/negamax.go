package connect4

const MaxScore int8 = 100
const MinScore int8 = -100

// Prioritize moves in the center
var ExploreOrder = [...]int8{3, 4, 2, 5, 1, 6, 0}

func CalcMove(board BitmapBoard) MiniMaxState {
	return newdp().iterativeDeepening(board)
}

func newdp() *dp {
	return &dp{Memo: make(map[uint64]int8)}
}

func (dp *dp) iterativeDeepening(board BitmapBoard) MiniMaxState {
	best := MiniMaxState{Value: MaxScore}
	var depth int8

	for depth = 5; depth < 19; depth++ {
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

func (dp *dp) negaMax(board BitmapBoard, depth int8, color int8, alpha int8, beta int8) MiniMaxState {
	hash := board.Hash()
	if _, ok := dp.Memo[hash]; ok {
		return MiniMaxState{Value: dp.Memo[hash]}
	}

	var winner int8 = board.CheckWinner()
	if winner != 0 {
		var endEval int8 = (50 * winner) - (depth * winner)
		return MiniMaxState{Value: endEval}
	} else if depth == dp.MaxDepth {
		return MiniMaxState{Value: 0}
	}

	var bestVal int8 = MinScore * color
	var foundOption bool = false
	var bestMove int8

	prevPos, prevMask := board.Pos, board.Mask
	isSymmetrical := IsSymmetrical(hash)

	for _, option := range ExploreOrder {
		if !board.CanPlay(option) || (isSymmetrical && option > 3) {
			continue
		}

		foundOption = true
		board.MakeMove(option, color)
		newVal := dp.negaMax(board, depth+1, -color, alpha, beta).Value

		// reverse move
		board.Mask = prevMask
		board.Pos = prevPos

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
		dp.Memo[hash] = res.Value
	}

	return res
}
