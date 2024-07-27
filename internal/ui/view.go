package ui

import (
	"fmt"
)

func (m Model) View() string {
	switch m.ViewState {
	case ViewSelection:
		return m.viewSelection()
	case ViewAdd:
		return m.AddModel.View()
	case ViewList:
		return m.viewList()
	case ViewTodo:
		return m.viewToDo()
	default:
		return "Unknown view"
	}
}

func (m AddTodoModel) View() string {
	clearScreen := "\033[H\033[2J"
	cursor := " "
	switch m.cursor {
	case 0:
		cursor = ">"
	}
	titleLine := fmt.Sprintf("%s Title: %s", cursor, m.Title)

	cursor = " "
	switch m.cursor {
	case 1:
		cursor = ">"
	}
	descriptionLine := fmt.Sprintf("%s Description: %s", cursor, m.Description)

	cursor = " "
	switch m.cursor {
	case 2:
		cursor = ">"
	}
	dueDateLine := fmt.Sprintf("%s Due Date: %s", cursor, m.DueDate)

	return fmt.Sprintf("%s%s\n%s\n%s\n\nTab to switch fields, Enter to save", clearScreen, titleLine, descriptionLine, dueDateLine)
}
