package app

import (
	"aurevoir/internal/components/types"
	"aurevoir/internal/layout"
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	v := tea.NewView(m.view())
	v.AltScreen = true
	v.WindowTitle = m.appName

	return v
}

func (m Model) view() string {
	containerAvailableSize := m.termSize
	containerContentSize, err := layout.GetStyleContentSize(
		m.containerStyle,
		containerAvailableSize,
	)
	if err != nil {
		log.Printf("get container content size error: %v", err)
		return m.viewErr("terminal size too small", containerAvailableSize)
	}
	containerContentAvailableSize, err := layout.GetStyleContentAvailableSize(
		m.containerStyle,
		containerAvailableSize,
	)
	if err != nil {
		log.Printf("get container content available size error: %v", err)
		return m.viewErr("terminal size too small", containerAvailableSize)
	}

	m.title.AvailableSize = containerContentAvailableSize
	renderedTitle, err := m.title.View()
	if err != nil {
		return m.viewErr("title rendering error", containerAvailableSize)
	}

	m.footer.AvailableSize = containerContentAvailableSize
	renderedFooter, err := m.footer.View()
	if err != nil {
		return m.viewErr("footer rendering error", containerAvailableSize)
	}

	containerContentFreeHeight := containerContentAvailableSize.Height -
		lipgloss.Height(renderedTitle) -
		lipgloss.Height(renderedFooter)
	containerContentFreeSize := types.Size{
		Width:  containerContentAvailableSize.Width,
		Height: containerContentFreeHeight,
	}

	m.menu.AvailableSize = containerContentFreeSize
	renderedMenu, err := m.menu.View()
	if err != nil {
		return m.viewErr("menu rendering error", containerAvailableSize)
	}

	joinedContent := lipgloss.JoinVertical(
		lipgloss.Center,
		renderedTitle,
		renderedMenu,
		renderedFooter,
	)

	renderedContainer := m.containerStyle.
		Width(containerContentSize.Width).
		Height(containerContentSize.Height).
		Render(joinedContent)

	return lipgloss.Place(
		containerAvailableSize.Width,
		containerAvailableSize.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedContainer,
	)
}

func (m Model) viewErr(err string, size types.Size) string {
	renderedErr := m.errorStyle.
		Width(size.Width).
		Render(err)

	return lipgloss.Place(
		size.Width,
		size.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedErr,
	)
}
