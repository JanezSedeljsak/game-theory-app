package main

import (
	"fmt"
	"tictactoe-minmax/core"
)

func main() {
	myTurn := true
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
		} else {
			var isDone bool = core.IsStalemate(board)
			if isDone {
				fmt.Println("Stalemate")
				break
			}
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
			boardCpy := board // copy by value
			ai := core.MinMax(boardCpy, 0, false, core.MinInt, core.MaxInt)
			//fmt.Println(ai)
			board[ai.Coords.X][ai.Coords.Y] = -1
		}

		fmt.Println(core.StringifyBoard(board))
	}

}
