package title

import (
	"aurevoir/internal/layout"
	"log"
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
	contentSize, err := layout.GetStyleContentSize(m.Style, m.AvailableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := layout.GetStyleContentAvailableSize(m.Style, m.AvailableSize)
	if err != nil {
		return "", err
	}
	truncText := layout.Truncate(layout.StripNonSpaceWhitespace(m.text), availableContentSize.Width, "")

	return (m.Style.
		Width(contentSize.Width).
		Render(truncText)), nil
}
