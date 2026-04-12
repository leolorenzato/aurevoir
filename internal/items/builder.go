package items

const (
	shutdownLabel  string = "Shutdown"
	rebootLabel    string = "Reboot"
	logoutLabel    string = "Logout"
	lockLabel      string = "Lock"
	suspendLabel   string = "Suspend"
	hibernateLabel string = "Hibernate"
)

func Build(cfg Cfg) []Item {

	return []Item{
		buildLockItem(cfg.Lock),
		buildShutdownItem(cfg.Shutdown),
		buildRebootItem(cfg.Reboot),
		buildLogoutItem(cfg.Logout),
		buildSuspendItem(cfg.Suspend),
		buildHibernateItem(cfg.Hibernate),
	}
}

func buildShutdownItem(cfg ShutdownCfg) Item {
	return Item{Enable: cfg.Show, Icon: cfg.Icon, Label: shutdownLabel, Cmd: cfg.Cmd}
}

func buildRebootItem(cfg RebootCfg) Item {
	return Item{Enable: cfg.Show, Icon: cfg.Icon, Label: rebootLabel, Cmd: cfg.Cmd}
}

func buildLogoutItem(cfg LogoutCfg) Item {
	return Item{Enable: cfg.Show, Icon: cfg.Icon, Label: logoutLabel, Cmd: cfg.Cmd}
}

func buildLockItem(cfg LockCfg) Item {
	return Item{Enable: cfg.Show, Icon: cfg.Icon, Label: lockLabel, Cmd: cfg.Cmd}
}

func buildSuspendItem(cfg SuspendCfg) Item {
	return Item{Enable: cfg.Show, Icon: cfg.Icon, Label: suspendLabel, Cmd: cfg.Cmd}
}

func buildHibernateItem(cfg HibernateCfg) Item {
	return Item{Enable: cfg.Show, Icon: cfg.Icon, Label: hibernateLabel, Cmd: cfg.Cmd}
}
