package app

import (
	"aurevoir/internal/components/footer"
	"aurevoir/internal/components/menu"
	"aurevoir/internal/components/title"
	"aurevoir/internal/components/types"
	"aurevoir/internal/theme"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	appName        string
	termSize       types.Size
	errorStyle     lipgloss.Style
	containerStyle lipgloss.Style
	title          title.Model
	menu           menu.Model
	footer         footer.Model
}

func NewModel(
	appName string,
	styles theme.Styles,
) Model {
	m := Model{
		appName:        appName,
		errorStyle:     styles.Error,
		containerStyle: styles.Container,
		title:          title.NewModel(titletext, styles.Title),
		menu: menu.NewModel(
			items,
			styles.Menu.Container,
			styles.Menu.Item,
			styles.Menu.SelectedItem,
		),
		footer: footer.NewModel(styles.Footer),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	if cmd := m.title.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}

	if cmd := m.menu.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}

	if cmd := m.footer.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}
