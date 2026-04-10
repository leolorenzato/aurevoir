package theme

func DefaultCfg() Cfg {
	return Cfg{
		Container: ContainerCfg{
			Border:        false,
			BorderColor:   "#9399b2",
			BorderRounded: true,
		},
		Title: TitleCfg{
			Border:        false,
			BorderColor:   "#9399b2",
			BorderRounded: true,
			TextColor:     "#f2cdcd",
		},
		Menu: MenuCfg{
			Border:                true,
			BorderColor:           "#9399b2",
			BorderRounded:         true,
			TextColor:             "#a6adc8",
			SelectedItemTextColor: "#A6E3A1",
		},
		ConfirmDialog: ConfirmDialogCfg{
			Border:                  true,
			BorderColor:             "#9399b2",
			BorderRounded:           true,
			TextColor:               "#a6adc8",
			SelectedOptionTextColor: "#A6E3A1",
		},
		Footer: FooterCfg{
			Border:        false,
			BorderColor:   "#9399b2",
			BorderRounded: true,
			TextColor:     "#6c7086",
		},
	}
}
