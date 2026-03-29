package footer

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const defaultText string = "• ↑/↓ to navigate • ↵ to select • ctrl+c to quit"

type Model struct {
	text          string
	AvailableSize types.Size
	Style         lipgloss.Style
}

func NewModel(style lipgloss.Style) Model {
	return Model{
		text:          defaultText,
		AvailableSize: types.Size{},
		Style:         style,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
