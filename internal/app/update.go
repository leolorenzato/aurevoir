package app

import (
	"aurevoir/internal/components/confirm_dialog"
	"aurevoir/internal/components/footer"
	"aurevoir/internal/components/menu"
	"aurevoir/internal/components/title"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termSize.Width = msg.Width
		m.termSize.Height = msg.Height
		cmds = append(cmds, tea.ClearScreen)
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			cmds = append(cmds, tea.Quit)
		}
	case menu.ItemSelectedMsg:
		cmds = append(
			cmds,
			func() tea.Msg { return menu.LockMsg{} },
			func() tea.Msg { return confirm_dialog.ShowMsg{} },
		)
	case confirm_dialog.CancelActionMsg:
		cmds = append(
			cmds,
			func() tea.Msg { return confirm_dialog.HideMsg{} },
			func() tea.Msg { return menu.UnlockMsg{} },
		)
	case confirm_dialog.ConfirmActionMsg:
		cmds = append(
			cmds,
			func() tea.Msg { return menu.LaunchSelectedItemCmdMsg{} },
			func() tea.Msg { return confirm_dialog.HideMsg{} },
			func() tea.Msg { return menu.UnlockMsg{} },
		)
	}

	subModelsCmd := m.updateSubModels(msg)

	return m, tea.Batch(tea.Batch(cmds...), subModelsCmd)
}

func (m Model) updateSubModels(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	updated, cmd := m.title.Update(msg)
	m.title = updated.(title.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	updated, cmd = m.menu.Update(msg)
	m.menu = updated.(menu.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	updated, cmd = m.footer.Update(msg)
	m.footer = updated.(footer.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}
