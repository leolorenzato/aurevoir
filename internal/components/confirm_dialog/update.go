package confirm_dialog

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case ShowMsg:
		m.show = true
		m.selectLeft()
		return m, nil
	case HideMsg:
		m.show = false
		return m, nil
	case tea.KeyPressMsg:
		if !m.show {
			return m, nil
		}
		switch msg.String() {
		case "left":
			m.selectLeft()
		case "right":
			m.selectRight()
		case "enter":
			if m.confirm {
				return m, func() tea.Msg { return ConfirmActionMsg{} }
			} else {
				return m, func() tea.Msg { return CancelActionMsg{} }
			}
		}
	}

	return m, nil
}
