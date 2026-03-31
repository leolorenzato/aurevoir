package confirm_dialog

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	show                bool
	confirm             bool
	AvailableSize       types.Size
	ContainerStyle      lipgloss.Style
	OptionStyle         lipgloss.Style
	SelectedOptionStyle lipgloss.Style
}

func NewModel(
	containerStyle lipgloss.Style,
	itemStyle lipgloss.Style,
	selectedItemStyle lipgloss.Style,
) Model {
	return Model{
		show:                false,
		confirm:             false,
		AvailableSize:       types.Size{},
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

func (m *Model) isShown() bool {
	return m.show
}
