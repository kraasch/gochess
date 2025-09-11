package chessctrl

import (
	"strconv"
	"strings"
)

const (
	NL2 = "\n" // TODO: generalize end-of-line sequence.
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func flipGrid(input string) string {
	// Split the input into lines.
	lines := strings.Split(input, "\n")
	// Reverse the order of lines.
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	// Reverse each line.
	for i := range lines {
		lines[i] = reverseString(lines[i])
	}
	// Join the lines back into a string.
	return strings.Join(lines, "\n")
}

func Flip(board string) string {
	return flipGrid(board)
}

func readCharAt(s *string, x, y int) string {
	// Assumes x and y are within bounds.
	ss := *s
	lines := strings.Split(ss, "\n")
	line := lines[y]
	runes := []rune(line)
	return string(runes[x])
}

func replaceCharAt(s *string, x, y int, insert string) {
	// Assumes x and y are within bounds.
	// Assumes insert is exactly of length 1.
	bytes := []byte(*s)
	bytes[x+(y*8)+y] = insert[0]
	*s = string(bytes)
}

// TODO: make this return an error, if move is invalid.
// TODO: 1) not a valid move string (invalid length, ie not equal).
// TODO: 2) not a valid move string (origin or destination invalid).
// TODO: 3) no piece at origin.
func validate(board *string, move string) bool {
	hasCorrectLength := true
	followsRegex := true
	hasPiece := true
	return hasCorrectLength && followsRegex && hasPiece
}

func validateInsertion(board *string, insertion string) bool {
	return true // TODO: implement.
}

func apply(board *string, move string) {
	// origin = o, destination = d, alphabetic = a; numeric = n.
	// a simple move instruction in chess follows the form:
	// oa + on + da + dn, ie A1B2 moves a piece from A1 to B2.
	upper := strings.ToUpper(move) // use upper case in order to subtract ASCII A = 65 from alphabetic coordinates.
	oa := int(upper[0]) - 65
	on, _ := strconv.Atoi(upper[1:2])
	on = 7 - (on - 1)
	da := int(upper[2]) - 65
	dn, _ := strconv.Atoi(upper[3:4])
	dn = 7 - (dn - 1)
	origin := readCharAt(board, oa, on)
	replaceCharAt(board, oa, on, " ")
	replaceCharAt(board, da, dn, origin)
}

// // NOTE: use later maybe.
// func convertPieceCode() {
// convert piece.
// str1 := "KQRBNPkqrbnp"
// str2 := "♔♕♖♗♘♙♚♛♜♝♞♟"
// found := ""
// for i, r := range str1 {
// 	if r == rune(pieceCode[0]) {
// 		found = str2[i : i+1]
// 		break
// 	}
// }
// }

func applyInsert(board *string, insertion string) {
	pieceCode := insertion[0:1]
	destinationCode := insertion[1:3]
	upper := strings.ToUpper(destinationCode) // use upper case in order to subtract ASCII A = 65 from alphabetic coordinates.
	da := int(upper[0]) - 65
	dn, _ := strconv.Atoi(upper[1:2])
	dn = 7 - (dn - 1)
	// replace destination with piece.
	replaceCharAt(board, da, dn, pieceCode)
}

func Move(board *string, moves string) {
	// TODO: Assumes moves are split by commas. (make tests).
	movesArr := strings.Split(moves, ",")
	for _, move := range movesArr {
		if validate(board, move) {
			apply(board, move)
		}
	}
}

func Insert(board *string, insertion string) {
	// TODO: First char: Assumes insertion starts with one of "rnbqkpRNBQKP"
	// TODO: Second and 3rd char: Assumes insertion ends with destination in the form of B1.
	if validateInsertion(board, insertion) {
		applyInsert(board, insertion)
	}
}
