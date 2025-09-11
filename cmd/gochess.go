package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	// for making a text prompt.
	"github.com/charmbracelet/bubbles/textinput"

	// for making a nice centred box.
	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"

	// basics.

	// local packages.
	chess "github.com/kraasch/gochess/pkg/chess"
)

var (
	// return value.
	output = ""
	// flags.
	verbose  = false
	suppress = false
	// styles.
	styleBox = lip.NewStyle().
			BorderStyle(lip.NormalBorder()).
			BorderForeground(lip.Color("56"))
	NL = fmt.Sprintln()
	// makes it possible to show a two-line help in a one-line prompt.
	toggleHelp = true
)

// type ( // TODO: remove later.
// 	errMsg error // TODO: remove later.
// ) // TODO: remove later.

type model struct {
	width     int
	height    int
	textInput textinput.Model
	cb        chess.Chessboard
	// err       error // TODO: remove later.
}

func (m model) Init() tea.Cmd {
	// return func() tea.Msg { return nil } // This line does nothing.
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			moveInput := m.textInput.Value()

			////////////////////////////////////////////////////////////
			// BEGIN
			// TODO: MAYBE let input be handled by a module called 'chessctrl'.
			// TODO: test as a separate module.
			// TODO: flip results in flipped board.
			// TODO: move results in change (test different options: castling, en-passant, etc)
			////////////////////////////////////////////////////////////
			pattern0 := `^insert [rnbqkpRNBQKP][a-h][1-8]$` // TODO: make this internal of the CHESSCTRL class.
			pattern1 := `^i [rnbqkpRNBQKP][a-h][1-8]$`      // TODO: make this internal of the CHESSCTRL class.
			pattern2 := `^insert [a-h][1-8]$`               // TODO: make this internal of the CHESSCTRL class.
			pattern3 := `^i [a-h][1-8]$`                    // TODO: make this internal of the CHESSCTRL class.
			regex0 := regexp.MustCompile(pattern0)
			regex1 := regexp.MustCompile(pattern1)
			regex2 := regexp.MustCompile(pattern2)
			regex3 := regexp.MustCompile(pattern3)
			if regex0.MatchString(moveInput) || regex1.MatchString(moveInput) {
				m.cb.Insert(moveInput[len(moveInput)-3:])
				m.textInput.SetValue("")
				m.textInput.Placeholder = "new piece..."
				return m, nil
			} else if regex2.MatchString(moveInput) || regex3.MatchString(moveInput) {
				m.cb.Insert(moveInput[len(moveInput)-2:])
				m.textInput.SetValue("")
				m.textInput.Placeholder = "removing..."
				return m, nil
			}

			switch moveInput {
			// TODO: implement these.
			case "s", "save":
				// TODO: implement.
				m.textInput.SetValue("")
				m.textInput.Placeholder = "implement: save"
			case "l", "load":
				// TODO: implement.
				m.textInput.SetValue("")
				m.textInput.Placeholder = "implement: load"
			case "f", "flip":
				// TODO: implement: flip the view.
				m.textInput.SetValue("")
				m.textInput.Placeholder = "implement: flip"
			case "u", "undo":
				// TODO: implement: undo the view.
				m.textInput.SetValue("")
				m.textInput.Placeholder = "implement: undo"
			case "r", "redo":
				// TODO: implement: redo the view.
				m.textInput.SetValue("")
				m.textInput.Placeholder = "implement: redo"

			case "h", "help":
				m.textInput.SetValue("")
				if toggleHelp {
					m.textInput.Placeholder = "move, insert, quit"
					toggleHelp = !toggleHelp
				} else {
					m.textInput.Placeholder = "'a7a6' 'i Qb5' 'q'"
					toggleHelp = !toggleHelp
				}
			case "q", "quit":
				output = "You quit (by command)!"
				return m, tea.Quit
			default:
				// string is not a baked command so it is a move command or an invalid command.
				pattern := `^[a-h][1-8][a-h][1-8]$` // TODO: make this internal of the CHESSCTRL class.
				regex := regexp.MustCompile(pattern)
				if regex.MatchString(moveInput) {
					m.cb.Move(moveInput) // TODO: implement: move the piece.
					m.textInput.SetValue("")
					m.textInput.Placeholder = "moving..."
				} else {
					m.textInput.SetValue("")
					m.textInput.Placeholder = "invalid command"
				}
			}
			////////////////////////////////////////////////////////////
			// END TODO: let input be handled by a module called 'chessctrl'.
			////////////////////////////////////////////////////////////

		case tea.KeyCtrlC, tea.KeyEsc:
			output = "You quit!"
			return m, tea.Quit
		}
		// case errMsg:     // TODO: remove later.
		// 	m.err = msg     // TODO: remove later.
		// 	return m, nil   // TODO: remove later.
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var str string
	// if verbose { // TODO: implement flags.
	// }
	str = m.cb.Display("entire", "pieces")
	str = styleBox.Render(str)
	str += NL + "  " + m.textInput.View()
	str = lip.Place(m.width, m.height, lip.Center, lip.Center, str)
	return str
}

func main() {
	// parse flags.
	flag.BoolVar(&verbose, "verbose", false, "Show info")
	flag.BoolVar(&suppress, "suppress", false, "Print nothing")
	flag.Parse()

	// make a new text input.
	ti := textinput.New()      // standard example.
	ti.Placeholder = "command" // standard example.
	ti.Focus()                 // standard example.
	ti.CharLimit = 156         // standard example.
	ti.Width = 20              // standard example.

	// add a chess board.
	cb := chess.NewBoard()

	// init model.
	m := model{0, 0, ti, cb}

	// start bubbletea.
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// print the last highlighted value in calendar to stdout.
	if !suppress {
		fmt.Println(output)
	}
} // fin.
