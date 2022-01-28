package core

import (
	"encoding/json"
	"fmt"
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

func (ttt *tictactoe) Init(aiStart bool) [][]int8 {
	ttt.Lock()
	defer ttt.Unlock()
	ttt.board = [][]int8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	if aiStart {
		aiMove := CalcMove(ttt.board)
		ttt.board[aiMove.Coords.Row][aiMove.Coords.Col] = -1
	}

	return ttt.board
}

func (ttt *tictactoe) Mutate(board [][]int8) string {
	ttt.board = board
	var isDraw bool = IsStalemate(board)
	var winner int = int(CheckWinner(ttt.board))
	var isDone bool = isDraw || winner != 0

	if isDone {
		gs := GameStatus{ttt.board, winner, isDone}
		return gs.Stringify()
	}

	aiMove := CalcMove(ttt.board)
	ttt.board[aiMove.Coords.Row][aiMove.Coords.Col] = -1
	winner = int(CheckWinner(ttt.board))
	gs := GameStatus{ttt.board, winner, IsStalemate(board) || winner != 0}
	return gs.Stringify()
}

func (ttt *tictactoe) Status(board [][]int8) string {
	ttt.board = board
	var isDraw bool = IsStalemate(ttt.board)
	var winner int = int(CheckWinner(ttt.board))
	var isDone bool = isDraw || winner != 0

	gs := GameStatus{ttt.board, winner, isDone}
	return gs.Stringify()
}

func (gs *GameStatus) Stringify() string {
	str, err := json.Marshal(gs)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str)
}
