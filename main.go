package main

import (
	"fmt"
	"tictactoe-minmax/core"
)

func main() {
	myTurn := true
	board := [][]int8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for {
		myTurn = !myTurn
		var winner int8 = core.CheckWinner(board)
		if winner != 0 {
			fmt.Printf("winner is: %d\n", winner)
			break
		} else if core.IsStalemate(board) {
			fmt.Println("Stalemate")
			break
		}

		var x, y int
		if myTurn {
			for {
				fmt.Println("Enter next coords")
				fmt.Scanf("\n%d %d", &x, &y)
				fmt.Printf("You have entered : %d %d\n", x, y)
				if board[x][y] == 0 {
					board[x][y] = 1
					break
				}

				fmt.Println("Invalid args!!!")
			}
		} else {
			aiMove := core.CalcMove(board)
			board[aiMove.Coords.Row][aiMove.Coords.Col] = -1
		}

		fmt.Println(core.StringifyBoard(board))
	}

}
