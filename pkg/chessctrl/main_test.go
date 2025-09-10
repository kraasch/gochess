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
	 * Test for Toast()
	 */
	{
		testingFunction: func(in TestList) (out string) {
			// board := in.inputArr[0]
			out = Toast()
			return out
		},
		tests: []TestList{
			{
				// turn in to these: ♔♕♖♗♘♙ ♚♛♜♝♞♟
				testName: "controller_flip_00",
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
				expectedValue:// this comment prevents start of string literal here.
				"Toast!",
				// "♜♞♝♛♚♝♞♜" + NL +
				// 	"♟♟♟♟♟♟♟♟" + NL +
				// 	"        " + NL +
				// 	"        " + NL +
				// 	"        " + NL +
				// 	"        " + NL +
				// 	"♙♙♙♙♙♙♙♙" + NL +
				// 	"♖♘♗♕♔♗♘♖",
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
