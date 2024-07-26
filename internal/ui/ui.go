package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TodoItem struct {
	title       string
	description string
	created     int
	due         int
}

type Model struct {
	ToDos     []TodoItem
	choices   []string
	cursor    int
	ViewState int
}

const (
	ViewSelection int = iota
	ViewAdd
	ViewDo
	ViewList
)

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func InitialModel() Model {
	return Model{
		choices: []string{"add", "do", "list"},
	}
}
