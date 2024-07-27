package ui

import (
	"context"
	"fmt"
)

func (m Model) viewSelection() string {
	s := "Todo TUI app:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s [%s]\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m Model) viewList() string {
	s := "To Do Items:\n\n"

	for i, todo := range m.ToDos {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %d. %s \n", cursor, i+1, todo.title)
	}

	s += "\nPress esc to go back.\n"
	return s
}

func (m Model) viewToDo() string {
	var s string

	s += fmt.Sprintf(
		"Title: %s\n\nDescription:\n%s\n\nDue: %s\nCreated: %s",
		m.CurrentTodo.title,
		m.CurrentTodo.description,
		m.CurrentTodo.due,
		m.CurrentTodo.created,
	)

	s += "\n\nPress ctrl+d to delete"

	return s
}

func (m AddTodoModel) Save() error {
	sql := `INSERT INTO todos (title, description, due) VALUES ($1, $2, $3)`
	_, err := m.conn.Exec(context.Background(), sql, m.Title, m.Description, m.DueDate)
	if err != nil {
		return fmt.Errorf("could not save todo item: %w", err)
	}
	return nil
}
