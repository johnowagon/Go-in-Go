package Engine

import (
	"fmt"

)

type Game struct {     // Main board, where pieces will live.
	board *Board
	blackScore int
	whiteScore int
}



func CreateGame() *Game {
	fmt.Println("Welcome to Go in Go!")
	fmt.Println("Input number for board size (X by X): ")
	var input int
	fmt.Scanln(&input)

	return &Game{
		board: NewBoard(input),
	}
}

func (g *Game) PerformMove(m *Move) {
	if m.color == white {
		g.board.pieces[m.xpos][m.ypos] = white
	}else{
		g.board.pieces[m.xpos][m.ypos] = black
	}
}