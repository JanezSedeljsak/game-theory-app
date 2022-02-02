package testing

import (
	"fmt"
	"game-theory-app/core/modules/connect4"
	"time"
)

func Run() {
	var board connect4.Board
	var test Test

	board.Init()
	board.Drop(3, 1)
	fmt.Println(board)

	test.set(false, false, false, 8)
	test.runTest(board, "Default Min-Max")
}

type Test struct {
	isDp         bool
	isAlphaBeta  bool
	isCheckOrder bool
	depth        int
	dp           connect4.Dp
}

func (r *Test) set(isdp bool, isalphabeta bool, ischeckorder bool, depth int) {
	r.isDp = isdp
	r.isAlphaBeta = isalphabeta
	r.isCheckOrder = ischeckorder
	r.depth = depth

	if isdp {
		r.dp = *newdp()
	}
}

func (t *Test) runTest(board connect4.Board, funcName string) {
	start := time.Now()

	fmt.Printf("Ran test: %s\n", funcName)
	res := t.CalcMove(board)
	fmt.Println(res)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %d\n", elapsed)
}

func newdp() *connect4.Dp {
	return &connect4.Dp{Memo: make(map[uint64]connect4.MiniMaxState)}
}
