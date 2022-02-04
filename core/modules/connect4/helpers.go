package connect4

func reverse(x uint64) uint64 {
	x = ((x & 0xaaaaaaaa) >> 1) | ((x & 0x55555555) << 1)
	x = ((x & 0xcccccccc) >> 2) | ((x & 0x33333333) << 2)
	x = ((x & 0xf0f0f0f0) >> 4) | ((x & 0x0f0f0f0f) << 4)
	x = ((x & 0xff00ff00) >> 8) | ((x & 0x00ff00ff) << 8)

	return (x >> 16) | (x << 16)
}

func getGameStatusBitmap(board BitmapBoard) GameStatus {
	return GameStatus{board.ToMatrix(), 0, false, []Coord{}}
}

func getGameStatus(board Board) GameStatus {
	var isDraw bool = board.IsDone()
	var gs GameStatus = board.CheckWinner()
	var isDone bool = isDraw || gs.Winner != 0

	return GameStatus{board.ToMatrix(), gs.Winner, isDone, gs.Coords}
}

func BitmapToMatrix(bitmap uint64) [Height][Width]int {
	var board [Height][Width]int
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			idx := i*Width + j

			if (bitmap>>idx)&1 == 1 {
				board[j][i] = 1
			} else {
				board[j][i] = 0
			}
		}
	}

	return board
}
