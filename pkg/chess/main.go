package chess

import (
	"fmt"
)

const (
	NL2   = "\n" // TODO: generalize end-of-line sequence.
	board = "  abcdefgh  " + NL2 +
		"8 rnbqkbnr 8" + NL2 +
		"7 pppppppp 7" + NL2 +
		"6          6" + NL2 +
		"5          5" + NL2 +
		"4          4" + NL2 +
		"3          3" + NL2 +
		"2 PPPPPPPP 2" + NL2 +
		"1 RNBQKBNR 1" + NL2 +
		"  abcdefgh  "
)

func ChessBoard(in string) string {
	return fmt.Sprintf("%v", board)
}
