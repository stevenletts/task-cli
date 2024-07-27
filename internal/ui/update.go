package ui

import (
	"context"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":

			var maxLen int

			switch m.ViewState {
			case ViewSelection:
				maxLen = len(m.choices) - 1
			case ViewAdd:
				maxLen = 0
			case ViewList:
				maxLen = len(m.ToDos) - 1
			}

			if m.cursor < maxLen {
				m.cursor++
			}

		case "ctrl+d":
			if m.ViewState == ViewTodo {

				sql := fmt.Sprintf("DELETE FROM todos WHERE id = %d", m.CurrentTodo.id)
				_, err := m.conn.Exec(context.Background(), sql)
				if err != nil {
					log.Printf("Error deleting todo item: %v", err)
				} else {
					log.Printf("Todo item with ID %d deleted successfully.", m.CurrentTodo.id)

					// Find the index of the current todo in the slice
					var index int
					for i, todo := range m.ToDos {
						if todo.id == m.CurrentTodo.id {
							index = i
							break
						}
					}

					m.ToDos = append(m.ToDos[:index], m.ToDos[index+1:]...)
				}
			}
		case "enter", " ":
			switch m.ViewState {
			case ViewSelection:
				m.ViewState = m.cursor + 1
				m.cursor = 0
			case ViewList:
				m.CurrentTodo = m.ToDos[m.cursor]
				m.ViewState = ViewTodo
				m.cursor = 0
			}

		case "esc":
			if m.ViewState == ViewTodo {
				m.ViewState = ViewList
			} else {
				m.ViewState = ViewSelection
			}
			m.cursor = 0
		}
	}

	return m, nil
}
