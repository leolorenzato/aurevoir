package title

import (
	"aurevoir/internal/layout"
	"fmt"

	"charm.land/lipgloss/v2"
)

func (m Model) View() (string, error) {
	rendered, err := m.render()
	if err != nil {
		return "", err
	}

	return rendered, nil
}

func (m Model) render() (string, error) {
	availableSize, err := m.getAvailableSize()
	if err != nil {
		return "", err
	}

	contentSize, err := layout.GetStyleContentSize(m.Style, availableSize)
	if err != nil {
		return "", err
	}

	availableContentSize, err := layout.GetStyleContentAvailableSize(
		m.Style,
		availableSize,
	)
	if err != nil {
		return "", err
	}

	if lipgloss.Width(m.text) > availableContentSize.Width {
		return "", fmt.Errorf("text too long")
	}

	return (m.Style.
		Width(contentSize.Width).
		Align(lipgloss.Center, lipgloss.Center).
		Render(m.text)), nil
}
