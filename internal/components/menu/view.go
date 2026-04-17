package menu

import (
	"aurevoir/internal/layout"
	"log"
	"slices"

	"charm.land/lipgloss/v2"
)

func (m Model) View() (string, error) {
	rendered, err := m.render()
	if err != nil {
		log.Printf("menu render error: %v", err)
		return "", err
	}

	return rendered, nil
}

func (m Model) render() (string, error) {
	containerAvailableSize, err := m.getAvailableSize()
	if err != nil {
		return "", err
	}
	contentSize, err := layout.GetStyleContentSize(m.ContainerStyle, containerAvailableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := layout.GetStyleContentAvailableSize(
		m.ContainerStyle,
		containerAvailableSize,
	)
	if err != nil {
		return "", err
	}

	var items []string

	// If cursor is visible, render from the top
	if m.cursor < availableContentSize.Height && m.cursor <= len(m.items)-1 {
		// The assumption is that an item has an height of 1
		items_num := min(availableContentSize.Height, len(m.items))
		for i := range items_num {
			cmd := m.items[i]
			item, err := m.renderMenuItem(cmd, i)
			if err != nil {
				return "", err
			}
			items = append(items, item)
		}
	}

	// If cursor is not visible, render from cursor backwards
	if m.cursor >= availableContentSize.Height && m.cursor <= len(m.items)-1 {
		// The assumption is that an item has an height of 1
		for i := m.cursor; i > m.cursor-availableContentSize.Height; i-- {
			cmd := m.items[i]
			item, err := m.renderMenuItem(cmd, i)
			if err != nil {
				return "", err
			}
			items = append(items, item)
		}
		slices.Reverse(items)
	}

	menuText := lipgloss.JoinVertical(lipgloss.Center, items...)

	return m.ContainerStyle.
		Width(contentSize.Width).
		Height(contentSize.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(menuText), nil
}

func (m Model) renderMenuItem(item Item, itemIndex int) (string, error) {
	containerAvailableSize, err := m.getAvailableSize()
	if err != nil {
		return "", err
	}
	availableContentWidth, err := layout.GetStyleContentAvailableWidth(
		m.ContainerStyle,
		containerAvailableSize.Width,
	)
	if err != nil {
		return "", err
	}

	var text string
	if item.Icon != "" {
		text = item.Icon + " " + item.Name
	} else {
		text = item.Name
	}

	truncText := layout.Truncate(
		layout.StripNonSpaceWhitespace(text),
		availableContentWidth,
		"...",
	)
	if itemIndex == m.cursor {
		return m.SelectedItemStyle.Render(truncText), nil
	}

	return m.ItemStyle.Render(truncText), nil
}
