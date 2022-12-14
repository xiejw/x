/*
[Tictactoe] is a program to play tic-tac-toe with RL algorithrm.

[Tictactoe]: https://en.wikipedia.org/wiki/Tic-tac-toe
*/
package main

import (
	"log"
	"os"

	"github.com/xiejw/x/rl/internal/drawing"
	"github.com/xiejw/x/rl/tictactoe/internal/board"
	"github.com/xiejw/x/rl/tictactoe/internal/policy"
)

func main() {
	// Prepare logging
	fo, err := os.Create("/tmp/tictactoe_log.txt")
	if err != nil {
		panic(err)
	}

	log.SetOutput(fo)
	defer fo.Close()

	// policy1 := &policy.DummyPolicy{
	// 	SleepSeconds: 1,
	// 	AIFirst:      true,
	// }

	policy2 := &policy.MCPolicy{
		SimulationCount: 1000,
		Seed:            123,
		AIFirst:         true,
	}
	policy2.Init()

	p := policy2

	drawing.BoardLoop("Tic Tac Toe", board.NewBoard(p))
}
