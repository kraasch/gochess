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
	str := Color(cb.Board, mode, format, cb.Flipped)
	return str
}

func (cb *Chessboard) Move(move string) {
	ctrl.Move(&cb.Board, move)
}

func (cb *Chessboard) Insert(insertion string) {
	ctrl.Insert(&cb.Board, insertion)
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
	if format == "filled" || format == "pieces" { // TODO: remove these two options from the code, just let the function do its job once called independent from the string which is passed in.
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

func insertSpacesAndColor(input string, flipped bool) string {
	var builder strings.Builder
	isEven := true
	if flipped { // if the board is flipped start with the other color initially.
		isEven = !isEven
	}
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

func surround(s, headerFooter string, flipped bool) string {
	res := ""
	mid := ""
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if !flipped {
			mid += fmt.Sprintf("%d %s %d\n", 8-i, line, 8-i)
		} else {
			mid += fmt.Sprintf("%d %s %d\n", i+1, line, i+1)
		}
	}
	res += headerFooter + "\n"
	res += mid
	res += headerFooter
	return res
}

func flipLines(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = reverseString(line)
	}
	return strings.Join(lines, "\n")
}

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

func Color(in, mode, format string, flipped bool) string { // TODO: IMPLEMENT THIS NEXT.
	str := ""
	narrow := ""
	wide := ""
	replacable := "" // Hacky string with z instead of b.
	if !flipped {
		narrow = "  abcdefgh  "
		wide = "   a b c d e f g h  "
		replacable = "   a z c d e f g h  " // Hacky string with z instead of b.
	} else {
		narrow = "  hgfedcba  "
		wide = "   h g f e d c b a  "
		replacable = "   h g f e d c z a  " // Hacky string with z instead of b.
		// flip the inner board too.
		in = flipLines(in)
	}
	if mode == "standard" {
		str = surround(in, narrow, flipped)
	} else if mode == "color" {
		in = insertSpacesAndColor(in, flipped)
		str = surround(in, wide, flipped)
	} else if mode == "entire" {
		in = insertSpacesAndColor(in, flipped)
		str = surround(in, replacable, flipped)
		str = Format(str, format)
		// Hacky function call, replacing all zs with bs, so that bishops can be distinguished from the b in the coordinate label.
		str = strings.ReplaceAll(str, "z", "b")
	}
	return str
}
