package menu

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			m.decrementCursor()
		case "down":
			m.incrementCursor()
		case "enter":
			item := m.getSelectedItem()
			return m, func() tea.Msg { return MenuSelectedItemMsg{Item: item} }
		}
	}

	return m, nil
}
