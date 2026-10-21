package main

import (
	"time"
	"uuid"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	header         string
	welcomeMessage string
	textInput      string
	chat           map[user]message
	sidebar        []user
	inputBox       string
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


func (m message) SendMessage(msg string) message {
	sentMessage := message{
		timestamp: time.Now(),
		text:      msg,
	}

	return sentMessage
}

func initialModel() model {
	return model{
		welcomeMessage: "Welcome! Type /help for commands. Ctrl+C to leave.",
		chat:           map[user]message{},
		sidebar:        []user{},
		inputBox:       "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		
}

func (m model) View() tea.View {
}

func main() {
	var chatHistory map[user]message

	var loggedInUsers []user

	var chatbox string
	tea.ClearScreen()
}
