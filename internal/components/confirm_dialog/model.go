package confirm_dialog

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Item struct {
	Name string
	Icon string
	Cmd  string
}

type Model struct {
	confirm             bool
	AvailableSize       types.Size
	ContainerStyle      lipgloss.Style
	OptionStyle         lipgloss.Style
	SelectedOptionStyle lipgloss.Style
}

func NewModel(
	items []Item,
	containerStyle lipgloss.Style,
	itemStyle lipgloss.Style,
	selectedItemStyle lipgloss.Style,
) Model {
	return Model{
		confirm:             false,
		ContainerStyle:      containerStyle,
		OptionStyle:         itemStyle,
		SelectedOptionStyle: selectedItemStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) selectLeft() {
	m.confirm = false
}

func (m *Model) selectRight() {
	m.confirm = true
}
