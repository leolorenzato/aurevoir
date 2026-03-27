package items

const (
	shutdownLabel string = "Shutdown"
	shutdownCmd   string = "shutdown -h now"
	rebootLabel   string = "Reboot"
	rebootCmd     string = "shutdown -r now"
)

func Build(cfg Cfg) []Item {

	return []Item{
		buildShutdownItem(cfg.Shutdown),
		buildRebootItem(cfg.Reboot),
	}
}

func buildShutdownItem(cfg ShutdownCfg) Item {
	return Item{Enable: true, Icon: cfg.icon, Label: shutdownLabel, Cmd: shutdownCmd}
}

func buildRebootItem(cfg RebootCfg) Item {
	return Item{Enable: true, Icon: cfg.icon, Label: rebootLabel, Cmd: rebootCmd}
}
