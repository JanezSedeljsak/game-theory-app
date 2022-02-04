package connect4

import "fmt"

func IsSymmetrical(bitmap uint64) bool {
	for i := 1; i <= 3; i++ {
		for j := 0; j < 7; j++ {
			leftIdx := 21 - i*7 + j
			rightIdx := 21 + i*7 + j
			if ((bitmap >> leftIdx) & 1) != ((bitmap >> rightIdx) & 1) {
				return false
			}
		}
	}

	return true
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

func PrintBitmap(bitmap uint64) {
	matrix := BitmapToMatrix(bitmap)
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			fmt.Printf("%d ", matrix[Height-i-1][j])
		}
		fmt.Println()
	}
}
