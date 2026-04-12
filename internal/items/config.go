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
	if r.Lock != nil {
		c.Lock.MergeRaw(*r.Lock)
	}
	if r.Shutdown != nil {
		c.Shutdown.MergeRaw(*r.Shutdown)
	}
	if r.Reboot != nil {
		c.Reboot.MergeRaw(*r.Reboot)
	}
	if r.Logout != nil {
		c.Logout.MergeRaw(*r.Logout)
	}
	if r.Suspend != nil {
		c.Suspend.MergeRaw(*r.Suspend)
	}
	if r.Hibernate != nil {
		c.Hibernate.MergeRaw(*r.Hibernate)
	}
}

type LockCfg struct {
	Show bool
	Icon string
	Cmd  string
}

func (c *LockCfg) MergeRaw(r RawLockCfg) {
	if r.Show != nil {
		c.Show = *r.Show
	}
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type ShutdownCfg struct {
	Show bool
	Icon string
	Cmd  string
}

func (c *ShutdownCfg) MergeRaw(r RawShutdownCfg) {
	if r.Show != nil {
		c.Show = *r.Show
	}
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type RebootCfg struct {
	Show bool
	Icon string
	Cmd  string
}

func (c *RebootCfg) MergeRaw(r RawRebootCfg) {
	if r.Show != nil {
		c.Show = *r.Show
	}
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type LogoutCfg struct {
	Show bool
	Icon string
	Cmd  string
}

func (c *LogoutCfg) MergeRaw(r RawLogoutCfg) {
	if r.Show != nil {
		c.Show = *r.Show
	}
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type SuspendCfg struct {
	Show bool
	Icon string
	Cmd  string
}

func (c *SuspendCfg) MergeRaw(r RawSuspendCfg) {
	if r.Show != nil {
		c.Show = *r.Show
	}
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}

type HibernateCfg struct {
	Show bool
	Icon string
	Cmd  string
}

func (c *HibernateCfg) MergeRaw(r RawHibernateCfg) {
	if r.Show != nil {
		c.Show = *r.Show
	}
	if r.Icon != nil {
		c.Icon = *r.Icon
	}
	if r.Cmd != nil {
		c.Cmd = *r.Cmd
	}
}
