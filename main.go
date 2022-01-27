package main

import (
	"fmt"
	"tictactoe-minmax/core"
)

func main() {
	myTurn := false
	board := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for {
		myTurn = !myTurn
		var winner int = core.CheckWinner(board)
		if winner != 0 {
			fmt.Printf("winner is: %d\n", winner)
			break
		}

		var x, y int
		if myTurn {
			fmt.Println("Enter next coords")
			fmt.Scanf("%d", &x)
			fmt.Scanf("%d", &y)
			board[x][y] = 1
		} else {
			ai := core.MinMax(board, 0, false, core.MinInt, core.MaxInt)
			fmt.Println(ai)
			board[ai.Coords.X][ai.Coords.Y] = -1
		}

		fmt.Println(core.StringifyBoard(board))
	}

}
