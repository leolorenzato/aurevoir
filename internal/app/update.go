package app

import (
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
	}

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

	return m, tea.Batch(cmds...)
}
