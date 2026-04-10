package items

type Cfg struct {
	Shutdown ShutdownCfg
	Reboot   RebootCfg
}

func (c *Cfg) MergeRaw(r RawCfg) {
	if r.Shutdown != nil {
		c.Shutdown.MergeRaw(*r.Shutdown)
	}
	if r.Reboot != nil {
		c.Reboot.MergeRaw(*r.Reboot)
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
