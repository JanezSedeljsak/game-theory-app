package connect4

const MaxScore int8 = 100
const MinScore int8 = -100

// Prioritize moves in the center
var ExploreOrder = [Width]int8{3, 4, 2, 5, 1, 6, 0}

func CalcMove(board Board, maxDepth int8) (int8, int8) {
	bmap := board.ToBitmap()
	return newdp(maxDepth).negaMax(bmap, 0, MinScore, MaxScore, false)
}

func newdp(maxDepth int8) *dp {
	return &dp{Memo: make(map[uint64]int8), MaxDepth: maxDepth}
}

func (dp *dp) negaMax(board BitmapBoard, depth int8, alpha int8, beta int8, winner bool) (int8, int8) {
	if depth == dp.MaxDepth {
		return 0, 0
	}

	hash := board.Hash()
	if _, ok := dp.Memo[hash]; ok {
		return 0, dp.Memo[hash]
	}

	var bestVal int8 = MinScore
	var bestMove int8

	code, moves, size := board.SortedMoves(hash)
	switch code {
	case 1:
		return 0, 0
	case 2:
		return moves[0].Col, 49 - depth
	}

	for index, move := range moves {
		_, newVal := dp.negaMax(move.Board, depth+1, -beta, -alpha, move.Winner)
		newVal *= -1

		if newVal > bestVal {
			bestVal = newVal
			bestMove = move.Col
		}

		if newVal >= beta {
			break
		}

		if bestVal > alpha {
			alpha = newVal
		}

		if index+1 == size {
			break
		}
	}

	if bestVal != 0 {
		dp.Memo[hash] = bestVal
	}

	return bestMove, bestVal
}
