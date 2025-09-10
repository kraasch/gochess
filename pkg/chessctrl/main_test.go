package chessctrl

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
	 * Test errors of Move()
	 */
	{ // TODO: implement.
		testingFunction: func(in TestList) (out string) {
			board := in.inputArr[0]
			moves := in.inputArr[1]
			Move(&board, moves)
			out = board
			return out
		},
		tests: []TestList{
			{
				testName: "controller_error_00",
				isMulti:  true,
				inputArr: []string{
					"rnbqkbnr" + NL +
						"pppppppp" + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"PPPPPPPP" + NL +
						"RNBQKBNR",
					"a7a6",
				},
				expectedValue: // this comment prevents start of string literal here.
				"rnbqkbnr" + NL +
					" ppppppp" + NL +
					"p       " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"PPPPPPPP" + NL +
					"RNBQKBNR",
			},
		},
	},

	/*
	 * Test for Move()
	 */
	{
		testingFunction: func(in TestList) (out string) {
			board := in.inputArr[0]
			moves := in.inputArr[1]
			Move(&board, moves)
			out = board
			return out
		},
		tests: []TestList{
			{
				testName: "controller_move_one-move_00",
				isMulti:  true,
				inputArr: []string{
					"rnbqkbnr" + NL +
						"pppppppp" + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"PPPPPPPP" + NL +
						"RNBQKBNR",
					"a7a6",
				},
				expectedValue: // this comment prevents start of string literal here.
				"rnbqkbnr" + NL +
					" ppppppp" + NL +
					"p       " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"PPPPPPPP" + NL +
					"RNBQKBNR",
			},
			{
				testName: "controller_move_multiple-moves_00",
				isMulti:  true,
				inputArr: []string{
					"rnbqkbnr" + NL +
						"pppppppp" + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"PPPPPPPP" + NL +
						"RNBQKBNR",
					"a7a6,b7b6",
				},
				expectedValue: // this comment prevents start of string literal here.
				"rnbqkbnr" + NL +
					"  pppppp" + NL +
					"pp      " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"PPPPPPPP" + NL +
					"RNBQKBNR",
			},
			{
				testName: "controller_move_multiple-moves_01",
				isMulti:  true,
				inputArr: []string{
					"rnbqkbnr" + NL +
						"pppppppp" + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"PPPPPPPP" + NL +
						"RNBQKBNR",
					"a7a6,a6a5",
				},
				expectedValue: // this comment prevents start of string literal here.
				"rnbqkbnr" + NL +
					" ppppppp" + NL +
					"        " + NL +
					"p       " + NL +
					"        " + NL +
					"        " + NL +
					"PPPPPPPP" + NL +
					"RNBQKBNR",
			},
		},
	},

	/*
	 * Test for Flip()
	 */
	{
		testingFunction: func(in TestList) (out string) {
			board := in.inputArr[0]
			out = Flip(board)
			return out
		},
		tests: []TestList{
			{
				testName: "controller_flip_00",
				isMulti:  true,
				inputArr: []string{
					"rnbqkbnr" + NL +
						"pppppppp" + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"PPPPPPPP" + NL +
						"RNBQKBNR",
				},
				expectedValue: // this comment prevents start of string literal here.
				"RNBKQBNR" + NL +
					"PPPPPPPP" + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"pppppppp" + NL +
					"rnbkqbnr",
			},
			{
				// turn in to these: ♔♕♖♗♘♙ ♚♛♜♝♞♟
				testName: "controller_flip_01",
				isMulti:  true,
				inputArr: []string{
					"rnb     " + NL +
						"ppp     " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        " + NL +
						"        ",
				},
				expectedValue: // this comment prevents start of string literal here.
				"        " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"        " + NL +
					"     ppp" + NL +
					"     bnr",
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
