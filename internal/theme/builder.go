package theme

import "charm.land/lipgloss/v2"

const (
	defaulHorizPadding int = 2
	defaulVertPadding  int = 0
)

func Build(cfg Cfg) Styles {

	return Styles{
		Error:         buildErrorStyle(),
		Container:     buildContainerStyle(cfg.Container),
		Title:         buildTitleStyle(cfg.Title),
		Menu:          buildMenuStyle(cfg.Menu),
		ConfirmDialog: buildConfirmDialogStyle(cfg.ConfirmDialog),
		Footer:        buildFooterStyle(cfg.Footer),
	}
}

func buildErrorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f38ba8"))
}

func buildContainerStyle(cfg ContainerCfg) lipgloss.Style {
	style := lipgloss.NewStyle().
		Padding(defaulVertPadding, defaulHorizPadding)

	if cfg.Border {
		var border lipgloss.Border
		if cfg.BorderRounded {
			border = lipgloss.RoundedBorder()
		} else {
			border = lipgloss.NormalBorder()
		}
		borderColor := lipgloss.Color(cfg.BorderColor)
		style = style.
			Border(border).
			BorderForeground(borderColor)
	}

	return style
}

func buildTitleStyle(cfg TitleCfg) lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor)).
		Align(lipgloss.Center).
		Bold(true)
	if cfg.Border {
		var border lipgloss.Border
		if cfg.BorderRounded {
			border = lipgloss.RoundedBorder()
		} else {
			border = lipgloss.NormalBorder()
		}
		borderColor := lipgloss.Color(cfg.BorderColor)
		style = style.
			Border(border).
			BorderForeground(borderColor)
	} else {
		var border lipgloss.Border
		border = lipgloss.Border{
			Top: " ", Bottom: " ",
			Left: " ", Right: " ",
			TopLeft: " ", TopRight: " ",
			BottomLeft: " ", BottomRight: " ",
		}
		style = style.
			Border(border)
	}

	return style
}

func buildMenuStyle(cfg MenuCfg) MenuStyles {
	containerStyle := lipgloss.NewStyle().
		Padding(defaulVertPadding, defaulHorizPadding)
	if cfg.Border {
		var border lipgloss.Border
		if cfg.BorderRounded {
			border = lipgloss.RoundedBorder()
		} else {
			border = lipgloss.NormalBorder()
		}
		borderColor := lipgloss.Color(cfg.BorderColor)
		containerStyle = containerStyle.
			Border(border).
			BorderForeground(borderColor)
	} else {
		var border lipgloss.Border
		border = lipgloss.Border{
			Top: " ", Bottom: " ",
			Left: " ", Right: " ",
			TopLeft: " ", TopRight: " ",
			BottomLeft: " ", BottomRight: " ",
		}
		containerStyle = containerStyle.
			Border(border)
	}

	itemStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor))

	selectedItemStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.SelectedItemTextColor))
	return MenuStyles{
		Container:    containerStyle,
		Item:         itemStyle,
		SelectedItem: selectedItemStyle,
	}
}

func buildConfirmDialogStyle(cfg ConfirmDialogCfg) ConfirmDialogStyles {
	containerStyle := lipgloss.NewStyle().
		Padding(defaulVertPadding, defaulHorizPadding)
	if cfg.Border {
		var border lipgloss.Border
		if cfg.BorderRounded {
			border = lipgloss.RoundedBorder()
		} else {
			border = lipgloss.NormalBorder()
		}
		borderColor := lipgloss.Color(cfg.BorderColor)
		containerStyle = containerStyle.
			Border(border).
			BorderForeground(borderColor)
	} else {
		var border lipgloss.Border
		border = lipgloss.Border{
			Top: " ", Bottom: " ",
			Left: " ", Right: " ",
			TopLeft: " ", TopRight: " ",
			BottomLeft: " ", BottomRight: " ",
		}
		containerStyle = containerStyle.
			Border(border)
	}

	optionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor))

	selectedOptionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.SelectedOptionTextColor))
	return ConfirmDialogStyles{
		Container:      containerStyle,
		Option:         optionStyle,
		SelectedOption: selectedOptionStyle,
	}
}

func buildFooterStyle(cfg FooterCfg) lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor)).
		Align(lipgloss.Center).
		Padding(defaulVertPadding, defaulHorizPadding)
	if cfg.Border {
		var border lipgloss.Border
		if cfg.BorderRounded {
			border = lipgloss.RoundedBorder()
		} else {
			border = lipgloss.NormalBorder()
		}
		borderColor := lipgloss.Color(cfg.BorderColor)
		style = style.
			Border(border).
			BorderForeground(borderColor)
	} else {
		var border lipgloss.Border
		border = lipgloss.Border{
			Top: " ", Bottom: " ",
			Left: " ", Right: " ",
			TopLeft: " ", TopRight: " ",
			BottomLeft: " ", BottomRight: " ",
		}
		style = style.
			Border(border)
	}

	return style
}
