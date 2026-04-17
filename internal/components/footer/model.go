package footer

import (
	"aurevoir/internal/components/types"
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	text          string
	AvailableSize types.Size
	Style         lipgloss.Style
}

func NewModel(style lipgloss.Style) Model {
	return Model{
		text:          "",
		AvailableSize: types.Size{},
		Style:         style,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) getAvailableSize() (types.Size, error) {
	if m.AvailableSize.Width <= 0 || m.AvailableSize.Height <= 0 {
		return types.Size{}, fmt.Errorf(
			"invalid available size, width: %d height %d",
			m.AvailableSize.Width,
			m.AvailableSize.Height,
		)
	}

	return m.AvailableSize, nil
}
