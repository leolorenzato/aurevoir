package theme

type Cfg struct {
	Container     ContainerCfg
	Title         TitleCfg
	Menu          MenuCfg
	ConfirmDialog ConfirmDialogCfg
	Footer        FooterCfg
}

func (c *Cfg) MergeRaw(r RawCfg) {
	if r.Container != nil {
		c.Container.MergeRaw(*r.Container)
	}
	if r.Title != nil {
		c.Title.MergeRaw(*r.Title)
	}
	if r.Menu != nil {
		c.Menu.MergeRaw(*r.Menu)
	}
	if r.ConfirmDialog != nil {
		c.ConfirmDialog.MergeRaw(*r.ConfirmDialog)
	}
	if r.Footer != nil {
		c.Footer.MergeRaw(*r.Footer)
	}
}

type ContainerCfg struct {
	Border        bool
	BorderColor   string
	BorderRounded bool
}

func (c *ContainerCfg) MergeRaw(r RawContainerCfg) {
	if r.Border != nil {
		c.Border = *r.Border
	}
	if r.BorderColor != nil {
		c.BorderColor = *r.BorderColor
	}
	if r.BorderRounded != nil {
		c.BorderRounded = *r.BorderRounded
	}
}

type TitleCfg struct {
	Border        bool
	BorderColor   string
	BorderRounded bool
	TextColor     string
}

func (c *TitleCfg) MergeRaw(r RawTitleCfg) {
	if r.Border != nil {
		c.Border = *r.Border
	}
	if r.BorderColor != nil {
		c.BorderColor = *r.BorderColor
	}
	if r.BorderRounded != nil {
		c.BorderRounded = *r.BorderRounded
	}
	if r.TextColor != nil {
		c.TextColor = *r.TextColor
	}
}

type MenuCfg struct {
	Border                bool
	BorderColor           string
	BorderRounded         bool
	TextColor             string
	SelectedItemTextColor string
}

func (c *MenuCfg) MergeRaw(r RawMenuCfg) {
	if r.Border != nil {
		c.Border = *r.Border
	}
	if r.BorderColor != nil {
		c.BorderColor = *r.BorderColor
	}
	if r.BorderRounded != nil {
		c.BorderRounded = *r.BorderRounded
	}
	if r.TextColor != nil {
		c.TextColor = *r.TextColor
	}
	if r.SelectedItemTextColor != nil {
		c.SelectedItemTextColor = *r.SelectedItemTextColor
	}
}

type ConfirmDialogCfg struct {
	Border                  bool
	BorderColor             string
	BorderRounded           bool
	TextColor               string
	SelectedOptionTextColor string
}

func (c *ConfirmDialogCfg) MergeRaw(r RawConfirmDialogCfg) {
	if r.Border != nil {
		c.Border = *r.Border
	}
	if r.BorderColor != nil {
		c.BorderColor = *r.BorderColor
	}
	if r.BorderRounded != nil {
		c.BorderRounded = *r.BorderRounded
	}
	if r.TextColor != nil {
		c.TextColor = *r.TextColor
	}
	if r.SelectedOptionTextColor != nil {
		c.SelectedOptionTextColor = *r.SelectedOptionTextColor
	}
}

type FooterCfg struct {
	Border        bool
	BorderColor   string
	BorderRounded bool
	TextColor     string
}

func (c *FooterCfg) MergeRaw(r RawFooterCfg) {
	if r.Border != nil {
		c.Border = *r.Border
	}
	if r.BorderColor != nil {
		c.BorderColor = *r.BorderColor
	}
	if r.BorderRounded != nil {
		c.BorderRounded = *r.BorderRounded
	}
	if r.TextColor != nil {
		c.TextColor = *r.TextColor
	}
}
