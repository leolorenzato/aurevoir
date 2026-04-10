package items

const (
	shutdownLabel string = "Shutdown"
	rebootLabel   string = "Reboot"
)

func Build(cfg Cfg) []Item {

	return []Item{
		buildShutdownItem(cfg.Shutdown),
		buildRebootItem(cfg.Reboot),
	}
}

func buildShutdownItem(cfg ShutdownCfg) Item {
	return Item{Enable: true, Icon: cfg.Icon, Label: shutdownLabel, Cmd: cfg.Cmd}
}

func buildRebootItem(cfg RebootCfg) Item {
	return Item{Enable: true, Icon: cfg.Icon, Label: rebootLabel, Cmd: cfg.Cmd}
}
