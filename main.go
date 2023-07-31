package main

import (
	"strings"

	//"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

type phase int

const (
	welcome phase = iota
	username
	host
)

type model struct {
	phase          phase
	focusOnButton  bool
	username, host string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if !m.focusOnButton {
				m.focusOnButton = true
				return m, nil
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var builder strings.Builder

	switch m.phase {
	case welcome:
		background := gloss.Color("8")
		if m.focusOnButton {
			background = gloss.Color("13")
		}

		title := gloss.NewStyle().Bold(true)
		button := gloss.NewStyle().
			Background(background).
			Padding(0, 2)

		builder.WriteString(title.Render("Hello! Welcome to the Beepy Setup Wizard™"))
		builder.WriteString("\nA quick guide to navigating the Wizard:\n")
		builder.WriteString("\t↹ Tab|⏎ Return|↓ Down\n")
		builder.WriteString("\t\tMove focus downward, or move to the next page\n")
		builder.WriteString("\t⇧↹ Shift-Tab|⇧⏎ Shift-Return|↑ Up\n")
		builder.WriteString("\t\tMove focus upward, or move to the previous page\n")
		builder.WriteString("\t^C Ctrl-C\n")
		builder.WriteString("\t\tQuit\n")
		builder.WriteString("We hope you enjoy your time with the Wizard 🧙!\n\n")
		builder.WriteString(button.Render("Next"))
	default:
		builder.WriteString("How did we get here?")
	}

	return builder.String()
}

func main() {
	m := model{
		phase: welcome,
	}
	tea.NewProgram(m).Run()
}
