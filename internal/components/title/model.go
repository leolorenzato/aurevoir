package title

import (
	"aurevoir/internal/components/types"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	text          string
	AvailableSize types.Size
	Style         lipgloss.Style
}

func NewModel(text string, style lipgloss.Style) Model {
	return Model{
		text:  text,
		Style: style,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
