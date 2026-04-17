package title

import (
	"aurevoir/internal/layout"
	"log"

	"charm.land/lipgloss/v2"
)

func (m Model) View() (string, error) {
	rendered, err := m.render()
	if err != nil {
		log.Printf("title render error: %v", err)
		return "", err
	}

	return rendered, nil
}

func (m Model) render() (string, error) {
	containerAvailableSize, err := m.getAvailableSize()
	if err != nil {
		return "", err
	}
	contentSize, err := layout.GetStyleContentSize(m.Style, containerAvailableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := layout.GetStyleContentAvailableSize(m.Style, containerAvailableSize)
	if err != nil {
		return "", err
	}
	truncText := layout.Truncate(
		layout.StripNonSpaceWhitespace(m.text),
		availableContentSize.Width,
		"",
	)

	return (m.Style.
		Width(contentSize.Width).
		Align(lipgloss.Center, lipgloss.Center).
		Render(truncText)), nil
}
