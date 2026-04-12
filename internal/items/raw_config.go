package items

type RawCfg struct {
	Lock      *RawLockCfg      `toml:"lock"`
	Shutdown  *RawShutdownCfg  `toml:"shutdown"`
	Reboot    *RawRebootCfg    `toml:"reboot"`
	Logout    *RawLogoutCfg    `toml:"logout"`
	Suspend   *RawSuspendCfg   `toml:"suspend"`
	Hibernate *RawHibernateCfg `toml:"hibernate"`
}

type RawLockCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawShutdownCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawRebootCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawLogoutCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawSuspendCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawHibernateCfg struct {
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}
