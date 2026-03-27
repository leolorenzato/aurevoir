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
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
}

type TitleCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
	TextColor     string `toml:"textColor"`
}

type MenuCfg struct {
	Border                bool   `toml:"border"`
	BorderColor           string `toml:"borderColor"`
	BorderRounded         bool   `toml:"borderRounded"`
	TextColor             string `toml:"textColor"`
	SelectedItemTextColor string `toml:"selectedItemTextColor"`
}

type ConfirmDialogCfg struct {
	Border                  bool   `toml:"border"`
	BorderColor             string `toml:"borderColor"`
	BorderRounded           bool   `toml:"borderRounded"`
	TextColor               string `toml:"textColor"`
	SelectedOptionTextColor string `toml:"selectedOptionTextColor"`
}

type FooterCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
	TextColor     string `toml:"textColor"`
}
