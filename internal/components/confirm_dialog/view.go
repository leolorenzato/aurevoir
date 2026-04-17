package confirm_dialog

import (
	"aurevoir/internal/layout"
	"fmt"
	"log"

	"charm.land/lipgloss/v2"
)

func (m Model) View() (string, error) {
	if !m.show {
		return "", nil
	}
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

	var yesStyle lipgloss.Style
	var noStyle lipgloss.Style
	if m.confirm {
		yesStyle = m.SelectedOptionStyle
		noStyle = m.OptionStyle
	} else {
		yesStyle = m.OptionStyle
		noStyle = m.SelectedOptionStyle
	}

	question := "Are you sure?"
	choices := noStyle.Render("No") + "  " + yesStyle.Render("Yes")
	text := lipgloss.JoinVertical(lipgloss.Center, question, choices)

	textWidth := lipgloss.Width(text)
	if textWidth > availableContentSize.Width {
		return "", fmt.Errorf(
			"invalid width %d, must be >= %d",
			availableContentSize.Width,
			textWidth,
		)
	}

	textHeight := lipgloss.Height(text)
	if textHeight > availableContentSize.Height {
		return "", fmt.Errorf(
			"invalid height %d, must be >= %d",
			availableContentSize.Height,
			textHeight,
		)
	}

	return m.ContainerStyle.
		Width(contentSize.Width).
		Height(contentSize.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(text), nil
}
