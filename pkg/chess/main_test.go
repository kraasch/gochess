package chess

import (

	// this is a test.
	"testing"

	// printing and formatting.
	"fmt"

	// other imports.
	"github.com/kraasch/godiff/godiff"
)

const (
	PW = "\x1b[1;38;2;255;0;0m"     // PW = player white, ANSI foreground color (= red).
	PB = "\x1b[1;38;2;100;100;100m" // PB = player black, ANSI foreground color (= gray).
	BW = "\x1b[1;38;2;100;100;100m" // BW = board white, ANSI background color (= white).
	BB = "\x1b[48;5;56m"            // BB = board black, ANSI background color (= purple).
	N  = "\x1b[0m"                  // ANSI clear formatting.
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
	 * Test for the function Toast().
	 */
	{
		testingFunction: func(in TestList) (out string) {
			inputValue := in.inputArr[0]
			out = ChessBoard(inputValue)
			return
		},
		tests: []TestList{
			{
				testName: "category_description_number00",
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
