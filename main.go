package main

import (
	"fmt"
	"os"
	"time"
	"uuid"

	// "charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type client struct {
	currentUser user
}

type server struct {
	port        int
	activeUsers []user
}

type clientModel struct {
	viewport    viewport.Model
	chat        []string
	sidebar     viewport.Model
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

type user struct {
	id       uuid.UUID
	username string
}

type message struct {
	author    user
	timestamp time.Time
	text      string
}

func (m message) String() string {
	return m.timestamp.String() + " " + m.author.username + ": " + m.text
}

type userInput struct {
	input string
}

func (c *client) SendMessage(userInput string) string {
	sentMessage := message{
		author:    c.currentUser,
		timestamp: time.Now(),
		text:      userInput,
	}

	return sentMessage.String()
}

func initialModel() clientModel {
	vp := viewport.New(viewport.WithWidth(300), viewport.WithHeight(50))
	vp.MouseWheelEnabled = true
	vp.AtTop()
	vp.SetContent(`Welcome!
	Type a message and hit enter to send a message.`)

	sb := viewport.New(viewport.WithWidth(50), viewport.WithHeight(300))

	ta := textarea.New()
	ta.Focus()

	return clientModel{
		viewport:    vp,
		chat:        []string{}, // TODO: This needs to pull chat history somehow
		sidebar:     sb,
		textarea:    ta,
		senderStyle: lipgloss.Style{},
		err:         nil,
	}
}

func (m clientModel) Init() tea.Cmd {
	return nil
}

func (m clientModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.textarea.Value() != "" {
				m.textarea.Value()
				return m, nil
			}
		default:
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m clientModel) View() tea.View {
	viewportView := m.viewport.View()
	v := tea.NewView(viewportView + "\n" + m.textarea.View())
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
