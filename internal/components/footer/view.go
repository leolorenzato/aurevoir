package footer

import (
	"aurevoir/internal/layout"
	"log"

	"charm.land/lipgloss/v2"
)

func (m Model) View() (string, error) {
	rendered, err := m.render()
	if err != nil {
		log.Printf("footer render error: %v", err)
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

	return (m.Style.
		Width(contentSize.Width).
		Align(lipgloss.Center, lipgloss.Center).
		Render(m.text)), nil
}
