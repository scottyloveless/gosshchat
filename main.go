package main

import (
	"fmt"
	"os"
	"time"
	"uuid"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport    viewport.Model
	chat        map[user]message
	sidebar     viewport.Model
	inputBox    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

type user struct {
	id       uuid.UUID
	username string
}

type message struct {
	timestamp time.Time
	text      string
}

type userInput struct {
	input string
}

func SendMessage(msg string) message {
	sentMessage := message{
		timestamp: time.Now(),
		text:      msg,
	}

	return sentMessage
}

func initialModel() model {
	vp := viewport.New(200, 50)
	vp.MouseWheelEnabled = true
	vp.AtTop()
	vp.SetContent(`Welcome!
	Type a message and hit enter to send a message.`)

	sb := viewport.New(50, 200)

	ta := textarea.New()
	ta.Focus()

	return model{
		viewport:    vp,
		chat:        map[user]message{},
		sidebar:     sb,
		inputBox:    ta,
		senderStyle: lipgloss.Style{},
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			if m.inputBox.Value() != "" {
				SendMessage(m.inputBox.Value())
				return m, nil
			}
		default:
			var cmd tea.Cmd
			m.inputBox, cmd = m.inputBox.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	viewportView := m.viewport.View()
	v := tea.NewView(viewportView + "\n" + m.inputBox.View())
	// c := m.inputBox.Cursor()
	// if c != nil {
	// 	c.Y += lipgloss.Height(viewportView)
	// }
	// v.Cursor = c
	v.AltScreen = true
	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
