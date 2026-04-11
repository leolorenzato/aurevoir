package app

import (
	"aurevoir/internal/components/confirm_dialog"
	"aurevoir/internal/components/footer"
	"aurevoir/internal/components/menu"
	"aurevoir/internal/components/title"
	"aurevoir/internal/components/types"
	"aurevoir/internal/items"
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
	confirm_dialog confirm_dialog.Model
	footer         footer.Model
}

func NewModel(
	appName string,
	items []items.Item,
	styles theme.Styles,
	dryRun bool,
) Model {
	var menuItems []menu.Item
	for _, item := range items {
		if item.Enable {
			menuItems = append(menuItems, itemToMenuItem(item))
		}
	}

	m := Model{
		appName:        appName,
		errorStyle:     styles.Error,
		containerStyle: styles.Container,
		title:          title.NewModel(styles.Title),
		menu: menu.NewModel(
			menuItems,
			styles.Menu.Container,
			styles.Menu.Item,
			styles.Menu.SelectedItem,
			dryRun,
		),
		confirm_dialog: confirm_dialog.NewModel(
			styles.ConfirmDialog.Container,
			styles.ConfirmDialog.Option,
			styles.ConfirmDialog.SelectedOption,
		),
		footer: footer.NewModel(styles.Footer),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, func() tea.Msg { return footer.ToggleMenuHintMsg{} })

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

func itemToMenuItem(item items.Item) menu.Item {
	return menu.Item{Name: item.Label, Icon: item.Icon, Cmd: item.Cmd}
}
