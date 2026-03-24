package confirm_dialog

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
		case "left":
			m.selectLeft()
		case "right":
			m.selectRight()
		case "enter":
			if m.confirm {
				return m, func() tea.Msg { return ConfirmAction{} }
			} else {
				return m, func() tea.Msg { return CancelAction{} }
			}
		}
	}

	return m, nil
}
