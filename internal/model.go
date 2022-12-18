package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rainbowriver/rivertodo/internal/todo"
)

type Model struct {
	//The lists of items
	ToDoList, CompletedList todo.Model

	//which item is selected
	cursor int

	//The size of the terminal
	width, height int

	//Error message
	err error
}

func (m *Model) Init() tea.Cmd {
	return nil //nil means no command
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// on keypress
	case tea.KeyMsg:
		switch msg.String() {
		// Quit if the user presses ctrl+c or q
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	// on window resize
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m *Model) View() string {
	return "Hello, World!"
}
