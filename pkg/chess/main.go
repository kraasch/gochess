package chess

import (
	"fmt"
	"strings"

	ctrl "github.com/kraasch/gochess/pkg/chessctrl" // TODO: import control module.
)

type Chessboard struct {
	Board   string
	Flipped bool
}

func NewBoard() Chessboard {
	return Chessboard{start, false}
}

func (cb *Chessboard) Flip() {
	cb.Flipped = true
}

func (cb *Chessboard) Unflip() {
	cb.Flipped = false
}

func (cb *Chessboard) Display(mode, format string) string {
	if cb.Flipped {
		return flipped0
	} else {
		// str := Color(cb.Board, "entire", "filled")
		str := Color(cb.Board, mode, format)
		return str
	}
}

func (cb *Chessboard) Move(move string) {
	ctrl.Move(&cb.Board, move)
}

const (
	// from https://en.wikipedia.org/wiki/Chess_symbols_in_Unicode
	// ♔♕♖♗♘♙
	// ♚♛♜♝♞♟
	// 🨀🨁🨂🨃🨄🨅🨆🨇🨈🨉🨊🨋🨌🨍🨎🨏🨐🨑🨒🨓🨔🨕🨖🨗🨘🨙🨚🨛🨜🨝🨞🨟🨠🨡🨢🨣🨤🨥🨦🨧🨨🨩🨪🨫🨬🨭🨮🨯🨰🨱🨲🨳🨴🨵🨶🨷🨸🨹🨺🨻🨼🨽🨾🨿🩀🩁🩂🩃🩄🩅🩆🩇🩈🩉🩊🩋🩌🩍🩎🩏🩐🩑🩒🩓
	PW    = "\x1b[38;5;244m" // PW = player white, ANSI color (light gray).
	PB    = "\x1b[38;5;236m" // PB = player black, ANSI color (dark gray).
	BW    = "\x1b[47m"       // BW = board white, ANSI color (white).
	BB    = "\x1b[40m"       // BB = board black, ANSI color (black).
	N     = "\x1b[0m"        // ANSI clear formatting.
	NL2   = "\n"             // TODO: generalize end-of-line sequence.
	start = "rnbqkbnr" + NL2 +
		"pppppppp" + NL2 +
		"        " + NL2 +
		"        " + NL2 +
		"        " + NL2 +
		"        " + NL2 +
		"PPPPPPPP" + NL2 +
		"RNBQKBNR"
	flipped0 = "   h g f e d c b a  " + NL2 + // TODO: remove this variable.
		"1 " + BW + " ♜" + BB + " ♞" + BW + " ♝" + BB + " ♚" + BW + " ♛" + BB + " ♝" + BW + " ♞" + BB + " ♜" + N + " 1" + NL2 +
		"2 " + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + N + " 2" + NL2 +
		"3 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 3" + NL2 +
		"4 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 4" + NL2 +
		"5 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 5" + NL2 +
		"6 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 6" + NL2 +
		"7 " + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + N + " 7" + NL2 +
		"8 " + BB + " ♖" + BW + " ♘" + BB + " ♗" + BW + " ♔" + BB + " ♕" + BW + " ♗" + BB + " ♘" + BW + " ♖" + N + " 8" + NL2 +
		"   h g f e d c b a  "
)

func Format(inputBoard, format string) string {
	if format == "none" {
		return inputBoard
	}
	str := ""
	str1 := "KQRBNPkqrbnp"
	str2 := "♔♕♖♗♘♙♚♛♜♝♞♟"
	if format == "filled" {
		str2 = "♚♛♜♝♞♟♚♛♜♝♞♟"
	}
	if format == "filled" || format == "pieces" {
		// Make "KQRBNPkqrbnp" into "♚♛♜♝♞♟♔♕♖♗♘♙".
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

func insertSpacesAndColor(input string) string {
	var builder strings.Builder
	isEven := true
	for _, ch := range input {
		if ch == '\n' {
			builder.WriteString(N)
			builder.WriteRune(ch)
			isEven = !isEven
		} else {
			if isEven {
				builder.WriteString(BW)
				isEven = false
			} else {
				builder.WriteString(BB)
				isEven = true
			}
			builder.WriteRune(' ')
			builder.WriteRune(ch)
		}
	}
	builder.WriteString(N)
	return builder.String()
}

func surround(s, headerFooter string) string {
	res := ""
	mid := ""
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		mid += fmt.Sprintf("%d %s %d\n", 8-i, line, 8-i)
	}
	res += headerFooter + "\n"
	res += mid
	res += headerFooter
	return res
}

func Color(in, mode, format string) string { // TODO: IMPLEMENT THIS NEXT.
	str := ""
	narrow := "  abcdefgh  "
	wide := "   a b c d e f g h  "
	replacable := "   a z c d e f g h  " // Hacky string with z instead of b.
	if mode == "standard" {
		str = surround(in, narrow)
	} else if mode == "color" {
		in = insertSpacesAndColor(in)
		str = surround(in, wide)
	} else if mode == "entire" {
		in = insertSpacesAndColor(in)
		str = surround(in, replacable)
		str = Format(str, format)
		// Hacky function call, replacing all zs with bs, so that bishops can be distinguished from the b in the coordinate label.
		str = strings.ReplaceAll(str, "z", "b")
	}
	return str
}
