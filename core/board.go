package core

func checkWinner(board [][]int) int {
	for i := 0; i < Size; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[1][i]
		}

		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[1][i]
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[1][1]
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[1][1]
	}

	return 0
}
