package theme

type RawCfg struct {
	Container     *RawContainerCfg     `toml:"container"`
	Title         *RawTitleCfg         `toml:"title"`
	Menu          *RawMenuCfg          `toml:"menu"`
	ConfirmDialog *RawConfirmDialogCfg `toml:"confirm_dialog"`
	Footer        *RawFooterCfg        `toml:"footer"`
}

type RawContainerCfg struct {
	Border        *bool   `toml:"border"`
	BorderColor   *string `toml:"border_color"`
	BorderRounded *bool   `toml:"border_rounded"`
}

type RawTitleCfg struct {
	Border        *bool   `toml:"border"`
	BorderColor   *string `toml:"border_color"`
	BorderRounded *bool   `toml:"border_rounded"`
	TextColor     *string `toml:"text_color"`
}

type RawMenuCfg struct {
	Border                *bool   `toml:"border"`
	BorderColor           *string `toml:"border_color"`
	BorderRounded         *bool   `toml:"border_rounded"`
	TextColor             *string `toml:"text_color"`
	SelectedItemTextColor *string `toml:"selected_item_text_color"`
}

type RawConfirmDialogCfg struct {
	Border                  *bool   `toml:"border"`
	BorderColor             *string `toml:"border_color"`
	BorderRounded           *bool   `toml:"border_rounded"`
	TextColor               *string `toml:"text_color"`
	SelectedOptionTextColor *string `toml:"selected_option_text_color"`
}

type RawFooterCfg struct {
	Border        *bool   `toml:"border"`
	BorderColor   *string `toml:"border_color"`
	BorderRounded *bool   `toml:"border_rounded"`
	TextColor     *string `toml:"text_color"`
}
