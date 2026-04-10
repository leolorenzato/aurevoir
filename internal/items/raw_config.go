package items

type RawCfg struct {
	Shutdown *RawShutdownCfg `toml:"shutdown"`
	Reboot   *RawRebootCfg   `toml:"reboot"`
}

type RawShutdownCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawRebootCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}
