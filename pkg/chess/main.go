package chess

import (
	"fmt"
)

const (
	// from https://en.wikipedia.org/wiki/Chess_symbols_in_Unicode
	// ♔♕♖♗♘♙
	// ♚♛♜♝♞♟
	// 🨀🨁🨂🨃🨄🨅🨆🨇🨈🨉🨊🨋🨌🨍🨎🨏🨐🨑🨒🨓🨔🨕🨖🨗🨘🨙🨚🨛🨜🨝🨞🨟🨠🨡🨢🨣🨤🨥🨦🨧🨨🨩🨪🨫🨬🨭🨮🨯🨰🨱🨲🨳🨴🨵🨶🨷🨸🨹🨺🨻🨼🨽🨾🨿🩀🩁🩂🩃🩄🩅🩆🩇🩈🩉🩊🩋🩌🩍🩎🩏🩐🩑🩒🩓
	PW    = "\x1b[0;36m" // PW = player white, ANSI color.
	PB    = "\x1b[0;31m" // PB = player black, ANSI color.
	BW    = "\x1b[47m"   // BW = board white, ANSI color.
	BB    = "\x1b[40m"   // BB = board black, ANSI color.
	N     = "\x1b[0m"    // ANSI clear formatting.
	NL2   = "\n"         // TODO: generalize end-of-line sequence.
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
	board2 = "  abcdefgh  " + NL2 +
		"8 " + BW + "r" + BB + "n" + BW + "b" + BB + "q" + BW + "k" + BB + "b" + BW + "n" + BB + "r" + N + " 8" + NL2 +
		"7 " + BB + "p" + BW + "p" + BB + "p" + BW + "p" + BB + "p" + BW + "p" + BB + "p" + BW + "p" + N + " 7" + NL2 +
		"6 " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + N + " 6" + NL2 +
		"5 " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + N + " 5" + NL2 +
		"4 " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + N + " 4" + NL2 +
		"3 " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + BB + " " + BW + " " + N + " 3" + NL2 +
		"2 " + BW + "P" + BB + "P" + BW + "P" + BB + "P" + BW + "P" + BB + "P" + BW + "P" + BB + "P" + N + " 2" + NL2 +
		"1 " + BB + "R" + BW + "N" + BB + "B" + BW + "Q" + BB + "K" + BW + "B" + BB + "N" + BW + "R" + N + " 1" + NL2 +
		"  abcdefgh  "
)

func Format(inputBoard, format string) string {
	str := ""
	if format == "pieces" {
		// Make "KQRBNPkqrbnp" into "♚♛♜♝♞♟♔♕♖♗♘♙".
		str1 := "KQRBNPkqrbnp"
		str2 := "♚♛♜♝♞♟♔♕♖♗♘♙"
		runes := []rune(str2)
		// loop over every character and replace.
		for _, currentRune := range inputBoard {
			found := ""
			for i, r := range str1 {
				if r == currentRune {
					found = string(runes[i])
				}
			}
			if found != "" {
				str += found
			} else {
				str += string(currentRune)
			}
		}
	}
	return str
}

func ChessBoard(in string) string {
	str := ""
	if in == "standard" {
		str = fmt.Sprintf("%v", board)
	} else if in == "color" {
		str = fmt.Sprintf("%v", board2)
	}
	return str
}
