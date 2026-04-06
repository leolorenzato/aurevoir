package theme

type Cfg struct {
	Container     ContainerCfg     `toml:"container"`
	Title         TitleCfg         `toml:"title"`
	Menu          MenuCfg          `toml:"menu"`
	ConfirmDialog ConfirmDialogCfg `toml:"confirm_dialog"`
	Footer        FooterCfg        `toml:"footer"`
}

type ContainerCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"border_color"`
	BorderRounded bool   `toml:"border_rounded"`
}

type TitleCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"border_color"`
	BorderRounded bool   `toml:"border_rounded"`
	TextColor     string `toml:"text_color"`
}

type MenuCfg struct {
	Border                bool   `toml:"border"`
	BorderColor           string `toml:"border_color"`
	BorderRounded         bool   `toml:"border_rounded"`
	TextColor             string `toml:"text_color"`
	SelectedItemTextColor string `toml:"selected_item_text_color"`
}

type ConfirmDialogCfg struct {
	Border                  bool   `toml:"border"`
	BorderColor             string `toml:"border_color"`
	BorderRounded           bool   `toml:"border_rounded"`
	TextColor               string `toml:"text_color"`
	SelectedOptionTextColor string `toml:"selected_option_text_color"`
}

type FooterCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"border_color"`
	BorderRounded bool   `toml:"border_rounded"`
	TextColor     string `toml:"text_color"`
}
