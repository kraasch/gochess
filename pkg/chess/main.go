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

func NewBoardNew() Chessboard { // TODO: make the only contructur. (and rename it of course)
	return Chessboard{start, false}
}

func NewBoard() Chessboard { // TODO: remove this.
	str := Format(start, "filled")
	str = Color(str, "entire")
	return Chessboard{str, false}
}

func (cb *Chessboard) Flip() {
	cb.Flipped = true
}

func (cb *Chessboard) Unflip() {
	cb.Flipped = false
}

func (cb *Chessboard) Display() string {
	if cb.Flipped {
		return flipped0
	} else {
		str := Format(cb.Board, "filled")
		str = Color(str, "entire")
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
	board3 = "   a b c d e f g h  " + NL2 + // TODO: remove this variable.
		"8 " + BW + " ♖" + BB + " ♘" + BW + " ♗" + BB + " ♕" + BW + " ♔" + BB + " ♗" + BW + " ♘" + BB + " ♖" + N + " 8" + NL2 +
		"7 " + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + N + " 7" + NL2 +
		"6 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 6" + NL2 +
		"5 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 5" + NL2 +
		"4 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 4" + NL2 +
		"3 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 3" + NL2 +
		"2 " + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + N + " 2" + NL2 +
		"1 " + BB + " ♜" + BW + " ♞" + BB + " ♝" + BW + " ♛" + BB + " ♚" + BW + " ♝" + BB + " ♞" + BW + " ♜" + N + " 1" + NL2 +
		"   a b c d e f g h  "
	board4 = "   a b c d e f g h  " + NL2 + // TODO: remove this variable.
		"8 " + BW + " ♜" + BB + " ♞" + BW + " ♝" + BB + " ♛" + BW + " ♚" + BB + " ♝" + BW + " ♞" + BB + " ♜" + N + " 8" + NL2 +
		"7 " + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + N + " 7" + NL2 +
		"6 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 6" + NL2 +
		"5 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 5" + NL2 +
		"4 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 4" + NL2 +
		"3 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 3" + NL2 +
		"2 " + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + N + " 2" + NL2 +
		"1 " + BB + " ♜" + BW + " ♞" + BB + " ♝" + BW + " ♛" + BB + " ♚" + BW + " ♝" + BB + " ♞" + BW + " ♜" + N + " 1" + NL2 +
		"   a b c d e f g h  "
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
	str := ""
	str1 := "KQRBNPkqrbnp"
	str2 := ""
	str2 = "♔♕♖♗♘♙♚♛♜♝♞♟"
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

func Color(in, mode string) string { // TODO: IMPLEMENT THIS NEXT.
	str := ""
	narrow := "  abcdefgh  "
	wide := "   a b c d e f g h  "
	if mode == "standard" {
		str = surround(in, narrow)
	} else if mode == "color" {
		in = insertSpacesAndColor(in)
		str = surround(in, wide)
	} else if mode == "entire" {
		str = fmt.Sprintf("%v", board3)
	}
	return str
}
