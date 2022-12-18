package todo

import (
	_ "github.com/charmbracelet/bubbles"
)

type TodoItem struct {
	Title       string
	Description string
}

func NewTodoItem(title string, description string) TodoItem {
	return TodoItem{
		Title:       title,
		Description: description,
	}
}

func (t *TodoItem) FilterValue() string {
	return t.Title
}
