package menu

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case BlockMsg:
		m.block = true
		return m, nil
	case tea.KeyPressMsg:
		if m.block {
			return m, nil
		}
		switch msg.String() {
		case "up":
			m.decrementCursor()
		case "down":
			m.incrementCursor()
		case "enter":
			item := m.getSelectedItem()
			return m, func() tea.Msg { return SelectedItemMsg{Item: item} }
		}
	}

	return m, nil
}
