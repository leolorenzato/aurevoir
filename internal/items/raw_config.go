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
	Show *bool   `toml:"show"`
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawShutdownCfg struct {
	Show *bool   `toml:"show"`
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawRebootCfg struct {
	Show *bool   `toml:"show"`
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawLogoutCfg struct {
	Show *bool   `toml:"show"`
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawSuspendCfg struct {
	Show *bool   `toml:"show"`
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}

type RawHibernateCfg struct {
	Show *bool   `toml:"show"`
	Icon *string `toml:"icon"`
	Cmd  *string `toml:"cmd"`
}
