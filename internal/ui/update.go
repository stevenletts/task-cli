package ui

import (
	"context"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.ViewState {
	case ViewAdd:
		// Delegate all messages to AddModel's Update function
		var cmd tea.Cmd
		m.AddModel, cmd = m.AddModel.Update(msg)

		// Check if the message is the "enter" key to trigger saving
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
			err := m.AddModel.Save()
			if err != nil {
				log.Printf("Error saving todo item: %v", err)
			} else {
				m.ToDos = append(m.ToDos, TodoItem{
					title:       m.AddModel.Title,
					description: m.AddModel.Description,
					due:         m.AddModel.DueDate,
				})
				// Reset the AddModel and switch views
				m.AddModel = AddTodoModel{conn: m.conn}
				m.ViewState = ViewList
			}
		}
		return m, cmd
	}

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

func (m AddTodoModel) Update(msg tea.Msg) (AddTodoModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// case "esc":

		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			m.cursor = (m.cursor + 1) % 3
		case "shift+tab":
			m.cursor = (m.cursor - 1 + 3) % 3
		case "enter":

			err := m.Save()
			if err != nil {
				log.Printf("Error saving todo item: %v", err)
			}

		case "backspace":
			switch m.cursor {
			case 0:
				if len(m.Title) > 0 {
					m.Title = m.Title[:len(m.Title)-1]
				}
			case 1:
				if len(m.Description) > 0 {
					m.Description = m.Description[:len(m.Description)-1]
				}
			case 2:
				if len(m.DueDate) > 0 {
					m.DueDate = m.DueDate[:len(m.DueDate)-1]
				}
			}
		default:
			switch m.cursor {
			case 0:
				m.Title += msg.String()
			case 1:
				m.Description += msg.String()
			case 2:
				m.DueDate += msg.String()
			}
		}
	}

	return m, nil
}
