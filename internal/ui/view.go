package ui

func (m Model) View() string {
	switch m.ViewState {
	case ViewSelection:
		return m.viewSelection()
	case ViewAdd:
		return m.viewAdd()
	case ViewDo:
		return m.viewDo()
	case ViewList:
		return m.viewList()
	default:
		return "Unknown view"
	}
}
