package tictactoe

import (
	"math/rand"
	"strings"
)

type dp struct {
	Memo map[int]Response
}

func CheckWinner(board [][]int) GameStatus {
	for i := 0; i < Size; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return GameStatus{Winner: board[i][1], Coords: []Coord{{i, 0}, {i, 1}, {i, 2}}}
		}

		if board[0][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return GameStatus{Winner: board[1][i], Coords: []Coord{{0, i}, {1, i}, {2, i}}}
		}
	}

	if board[1][1] != 0 {
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
			return GameStatus{Winner: board[1][1], Coords: []Coord{{0, 0}, {1, 1}, {2, 2}}}
		}

		if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
			return GameStatus{Winner: board[1][1], Coords: []Coord{{0, 2}, {1, 1}, {2, 0}}}
		}
	}

	return GameStatus{Winner: 0}
}

func IsDone(board [][]int) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func StringifyBoard(board [][]int) string {
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

func BoardHash(board [][]int) int {
	var hash int = 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			var val int = board[i][j] + 1
			hash = 31*hash + val
		}
	}

	return hash
}

func GetOpenSpots(board [][]int) []Coord {
	var spots []Coord
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if board[i][j] == 0 {
				spots = append(spots, Coord{i, j})
			}
		}
	}

	return spots
}

func (s *Actions) Init(aiStart bool, isAdvanced bool) [][]int {
	s.Lock()
	defer s.Unlock()

	s.board = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	if aiStart {
		if isAdvanced {
			aiMove := CalcMove(s.board)
			s.board[aiMove.Coords.Row][aiMove.Coords.Col] = -1
		} else {
			moveOptions := GetOpenSpots(s.board)
			randCoord := moveOptions[rand.Intn(len(moveOptions))]
			s.board[randCoord.Row][randCoord.Col] = -1
		}
	}

	return s.board
}

func (s *Actions) Mutate(board [][]int) string {
	s.Lock()
	defer s.Unlock()

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	s.board = board
	gs = getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	aiMove := CalcMove(s.board)
	s.board[aiMove.Coords.Row][aiMove.Coords.Col] = -1
	gs = getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) Multiplayer(board [][]int) string {
	s.Lock()
	defer s.Unlock()

	s.board = board
	gs := getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) RandomMove(board [][]int) string {
	s.Lock()
	defer s.Unlock()

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	s.board = board
	gs = getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	moveOptions := GetOpenSpots(s.board)
	randCoord := moveOptions[rand.Intn(len(moveOptions))]
	s.board[randCoord.Row][randCoord.Col] = -1
	gs = getGameStatus(s.board)
	return gs.Stringify()
}

func getGameStatus(board [][]int) GameStatus {
	var isDraw bool = IsDone(board)
	var gs GameStatus = CheckWinner(board)
	var isDone bool = isDraw || gs.Winner != 0

	return GameStatus{board, gs.Winner, isDone, gs.Coords}
}
