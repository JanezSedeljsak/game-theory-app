package core

import (
	"strings"
)

func CheckWinner(board [][]int8) int8 {
	for i := 0; i < Size; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][1]
		}

		if board[0][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[1][i]
		}
	}

	if board[1][1] != 0 {
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
			return board[1][1]
		}

		if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
			return board[1][1]
		}
	}

	return 0
}

func IsStalemate(board [][]int8) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func StringifyBoard(board [][]int8) string {
	var sb strings.Builder
	for i := 0; i < Size; i++ {
		if i != 0 {
			sb.WriteString("\n")
		}
		for j := 0; j < Size; j++ {
			switch board[i][j] {
			case 0:
				sb.WriteString("_")
			case 1:
				sb.WriteString("O")
			case -1:
				sb.WriteString("X")
			}
		}

	}

	return sb.String()
}

func BoardHash(board [][]int8) int {
	var hash int = 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			var val int = int(board[i][j]) + 1
			hash = 31*hash + val
		}
	}

	return hash
}
