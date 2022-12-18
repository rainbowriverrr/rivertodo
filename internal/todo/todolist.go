package todo

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Implements the bubbletea.Model interface
type Model struct {
	Items  []TodoItem
	cursor int //which item is selected
}

func (m *Model) Init() tea.Cmd {
	return nil //nil means no command
}

func (m *Model) AddItem(item TodoItem) {
	m.Items = append(m.Items, item)
}
