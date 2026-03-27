package items

type Cfg struct {
	Shutdown ShutdownCfg `toml:"shutdown"`
	Reboot   RebootCfg   `toml:"reboot"`
}

type ShutdownCfg struct {
	icon string `toml:"icon"`
}

type RebootCfg struct {
	icon string `toml:"icon"`
}
