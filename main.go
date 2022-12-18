package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         //items on the to-do list
	cursor   int              //which item is selected
	selected map[int]struct{} //which items are selected
}

func initialModel() model {
	return model{
		// initial list of items
		choices: []string{
			"Check Amazon Order Status",
			"Study CS 3307",
			"Finish this demo",
		},
		// initially, nothing is selected
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil //nil means no command
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		//figure out which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "What should you do next?\n\n"
	for i, choice := range m.choices {
		cursor := "  "
		if i == m.cursor {
			cursor = "▸ "
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("[%s] %s %s\n", cursor, checked, choice)
	}

	s += "\nPress Q to quit"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh no! %v", err)
		os.Exit(1)
	}
}
