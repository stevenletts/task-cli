package ui

import (
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
			current := m.ViewState

			var maxLen int

			switch current {
			case 0:
				maxLen = len(m.choices) - 1
			case 1:
				maxLen = 0
			case 2:
				maxLen = len(m.ToDos) - 1
			}

			if m.cursor < maxLen {
				m.cursor++
			}

		case "enter", " ":
			switch m.ViewState {
			case ViewSelection:
				m.ViewState = m.cursor + 1
				m.cursor = 0
			}

		case "esc":
			m.ViewState = ViewSelection
			m.cursor = 0
		}
	}

	return m, nil
}
