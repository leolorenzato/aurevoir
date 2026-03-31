package menu

import (
	"aurevoir/internal/components/types"
	"log"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case LaunchSelectedItemCmdMsg:
		item := m.getSelectedItem()
		if m.dryRun {
			log.Printf("dry run: command to launch: %s", item.Cmd)
			return m, nil
		}
		return m, LaunchCmd(item.Cmd)
	case LockMsg:
		m.lock = true
		return m, nil
	case UnlockMsg:
		m.lock = false
		return m, nil
	case tea.KeyPressMsg:
		if m.lock {
			return m, nil
		}
		switch msg.String() {
		case "up":
			m.decrementCursor()
		case "down":
			m.incrementCursor()
		case "enter":
			return m, func() tea.Msg { return ItemSelectedMsg{} }
		}
	}

	return m, nil
}
