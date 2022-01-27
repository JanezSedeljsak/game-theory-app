package main

import (
	"fmt"
	core "tictactoe-minmax/core"
)

func main() {
	board := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Println(board)
	val := core.MinMax(board, 0, true, core.MinInt, core.MaxInt)
	fmt.Println(val)
}
