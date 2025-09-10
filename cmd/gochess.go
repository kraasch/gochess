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
	engine "github.com/kraasch/gochess/pkg/chess"
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
)

// type ( // TODO: remove later.
// 	errMsg error // TODO: remove later.
// ) // TODO: remove later.

type model struct {
	width     int
	height    int
	textInput textinput.Model
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
			text := m.textInput.Value()

			////////////////////////////////////////////////////////////
			// BEGIN TODO: let input be handled by a module called 'chessctrl'.
			// TODO: test as a separate module.
			// TODO: flip results in flipped board.
			// TODO: move results in change (test different options: castling, en-passant, etc)
			////////////////////////////////////////////////////////////
			switch text {
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
			case "h", "help":
				m.textInput.SetValue("")
				m.textInput.Placeholder = "move: a7a6 quit: q"
			case "q", "quit":
				output = "You quit (by command)!"
				return m, tea.Quit
			default:
				// string is not a baked command so it is a move command or an invalid command.
				pattern := `^[a-h][1-8][a-h][1-8]$`
				regex := regexp.MustCompile(pattern)
				if regex.MatchString(text) {
					m.textInput.SetValue("")
					m.textInput.Placeholder = "moving..."
					// TODO: implement: move the piece.
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
	str = engine.ChessBoard()
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

	// init model.
	m := model{0, 0, ti}

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
