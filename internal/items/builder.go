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
	type builderFunc func(Cfg) Item

	var builders = map[ItemId]builderFunc{
		lockId: func(cfg Cfg) Item {
			return buildLockItem(cfg.Lock)
		},
		shutdownId: func(cfg Cfg) Item {
			return buildShutdownItem(cfg.Shutdown)
		},
		rebootId: func(cfg Cfg) Item {
			return buildRebootItem(cfg.Reboot)
		},
		logoutId: func(cfg Cfg) Item {
			return buildLogoutItem(cfg.Logout)
		},
		suspendId: func(cfg Cfg) Item {
			return buildSuspendItem(cfg.Suspend)
		},
		hibernateId: func(cfg Cfg) Item {
			return buildHibernateItem(cfg.Hibernate)
		},
	}

	var items []Item
	for _, itemId := range cfg.Order.ItemIds {
		items = append(items, builders[itemId](cfg))
	}

	return items
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
