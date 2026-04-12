package items

type Cfg struct {
	Lock      LockCfg
	Shutdown  ShutdownCfg
	Reboot    RebootCfg
	Logout    LogoutCfg
	Suspend   SuspendCfg
	Hibernate HibernateCfg
}

func (c *Cfg) MergeRaw(r RawCfg) {
	if r.Shutdown != nil {
		c.Shutdown.MergeRaw(*r.Shutdown)
	}
	if r.Reboot != nil {
		c.Reboot.MergeRaw(*r.Reboot)
	}
}

type LockCfg struct {
	Icon string
	Cmd  string
}

func (c *LockCfg) MergeRaw(r RawLockCfg) {
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type ShutdownCfg struct {
	Icon string
	Cmd  string
}

func (c *ShutdownCfg) MergeRaw(r RawShutdownCfg) {
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type RebootCfg struct {
	Icon string
	Cmd  string
}

func (c *RebootCfg) MergeRaw(r RawRebootCfg) {
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type LogoutCfg struct {
	Icon string
	Cmd  string
}

func (c *LogoutCfg) MergeRaw(r RawLogoutCfg) {
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type SuspendCfg struct {
	Icon string
	Cmd  string
}

func (c *SuspendCfg) MergeRaw(r RawSuspendCfg) {
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type HibernateCfg struct {
	Icon string
	Cmd  string
}

func (c *HibernateCfg) MergeRaw(r RawHibernateCfg) {
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}
