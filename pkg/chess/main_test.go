package chess

import (

	// this is a test.
	"testing"

	// printing and formatting.
	"fmt"

	// other imports.
	"github.com/kraasch/godiff/godiff"
)

var NL = fmt.Sprintln()

type TestList struct {
	testName      string
	isMulti       bool
	inputArr      []string
	expectedValue string
}

type TestSuite struct {
	testingFunction func(in TestList) string
	tests           []TestList
}

var suites = []TestSuite{
	/*
	 * Test for the function Format-- for formatting the inner board.
	 */
	{
		testingFunction: func(in TestList) (out string) {
			board := in.inputArr[0]
			format := in.inputArr[1]
			out = Format(board, format)
			return
		},
		tests: []TestList{
			{
				// turn in to these: ♔♕♖♗♘♙ ♚♛♜♝♞♟
				testName: "board-inner_convert-characters_00",
				isMulti:  false,
				inputArr: []string{
					"rnbqkbnr" + NL +
						"pppppppp" + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"PPPPPPPP" + NL +
						"RNBQKBNR",
					"pieces",
				},
				expectedValue: // this comment prevents start of string literal here.
				"♖♘♗♕♔♗♘♖" + NL +
					"♙♙♙♙♙♙♙♙" + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"♟♟♟♟♟♟♟♟" + NL +
					"♜♞♝♛♚♝♞♜",
			},
		},
	},

	/*
	 * Test for the function Color -- for formatting the outer board.
	 */
	{
		testingFunction: func(in TestList) (out string) {
			inputValue := in.inputArr[0]
			out = Color(start, inputValue)
			return
		},
		tests: []TestList{
			{
				testName: "board-outer_print-as-text+no-formatting_00",
				isMulti:  false,
				inputArr: []string{"standard"},
				expectedValue: // this comment prevents start of string literal here.
				"  abcdefgh  " + NL +
					"8 rnbqkbnr 8" + NL +
					"7 pppppppp 7" + NL +
					"6          6" + NL +
					"5          5" + NL +
					"4          4" + NL +
					"3          3" + NL +
					"2 PPPPPPPP 2" + NL +
					"1 RNBQKBNR 1" + NL +
					"  abcdefgh  ",
			},
			{
				testName: "board-outer_print-as-text+with-color_00",
				isMulti:  false,
				inputArr: []string{"color"},
				expectedValue: // this comment prevents start of string literal here.
				"   a b c d e f g h  " + NL +
					"8 " + BW + " r" + BB + " n" + BW + " b" + BB + " q" + BW + " k" + BB + " b" + BW + " n" + BB + " r" + N + " 8" + NL +
					"7 " + BB + " p" + BW + " p" + BB + " p" + BW + " p" + BB + " p" + BW + " p" + BB + " p" + BW + " p" + N + " 7" + NL +
					"6 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 6" + NL +
					"5 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 5" + NL +
					"4 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 4" + NL +
					"3 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 3" + NL +
					"2 " + BW + " P" + BB + " P" + BW + " P" + BB + " P" + BW + " P" + BB + " P" + BW + " P" + BB + " P" + N + " 2" + NL +
					"1 " + BB + " R" + BW + " N" + BB + " B" + BW + " Q" + BB + " K" + BW + " B" + BB + " N" + BW + " R" + N + " 1" + NL +
					"   a b c d e f g h  ",
			},
		},
	},

	/*
	 * Test for the function ChessBoard -- for creating a chessboard.
	 */
	{
		testingFunction: func(in TestList) (out string) {
			out = ChessBoard()
			return
		},
		tests: []TestList{
			{
				testName: "create-full-board_simple-board_00",
				isMulti:  false,
				inputArr: []string{"color"},
				expectedValue: // this comment prevents start of string literal here.
				"   a b c d e f g h  " + NL +
					"8 " + BW + " ♖" + BB + " ♘" + BW + " ♗" + BB + " ♕" + BW + " ♔" + BB + " ♗" + BW + " ♘" + BB + " ♖" + N + " 8" + NL +
					"7 " + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + BB + " ♙" + BW + " ♙" + N + " 7" + NL +
					"6 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 6" + NL +
					"5 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 5" + NL +
					"4 " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + N + " 4" + NL +
					"3 " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + BB + "  " + BW + "  " + N + " 3" + NL +
					"2 " + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + BW + " ♟" + BB + " ♟" + N + " 2" + NL +
					"1 " + BB + " ♜" + BW + " ♞" + BB + " ♝" + BW + " ♛" + BB + " ♚" + BW + " ♝" + BB + " ♞" + BW + " ♜" + N + " 1" + NL +
					"   a b c d e f g h  ",
			},
		},
	},
}

func TestAll(t *testing.T) {
	for _, suite := range suites {
		for _, test := range suite.tests {
			name := test.testName
			t.Run(name, func(t *testing.T) {
				exp := test.expectedValue
				got := suite.testingFunction(test)
				if exp != got {
					if test.isMulti {
						t.Errorf("In '%s':\n", name)
						diff := godiff.CDiff(exp, got)
						t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
						t.Errorf("exp/got:\n%s\n", diff)
					} else {
						t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
					}
				}
			})
		}
	}
}
