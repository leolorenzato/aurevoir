package footer

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
)

const (
	menuHintText          string = "• ↑/↓ to navigate • ↵ to select • ctrl+c to quit"
	confirmDialogHintText string = "• ←/→ to navigate • ↵ to select • ctrl+c to quit"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg.(type) {
	case ToggleConfirmDialogHintMsg:
		m.text = confirmDialogHintText
	case ToggleMenuHintMsg:
		m.text = menuHintText
	}

	return m, nil
}
