package menu

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
	lock              bool
	cursor            int
	items             []Item
	AvailableSize     types.Size
	ContainerStyle    lipgloss.Style
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
}

func NewModel(
	items []Item,
	containerStyle lipgloss.Style,
	itemStyle lipgloss.Style,
	selectedItemStyle lipgloss.Style,
) Model {
	return Model{
		lock:              false,
		items:             items,
		AvailableSize:     types.Size{},
		ContainerStyle:    containerStyle,
		ItemStyle:         itemStyle,
		SelectedItemStyle: selectedItemStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) incrementCursor() {
	m.cursor++
	m.clipCursor()
}

func (m *Model) decrementCursor() {
	m.cursor--
	m.clipCursor()
}

func (m *Model) clipCursor() {
	n := len(m.items)
	if n == 0 {
		m.cursor = 0
		return
	}

	m.cursor = (m.cursor%n + n) % n
}

func (m *Model) getSelectedItem() Item {
	return m.items[m.cursor]
}

func (m *Model) IsLocked() bool {
	return m.lock
}
