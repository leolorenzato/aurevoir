package items

const (
	defaultShutdownCmd string = "shutdown -h now"
	defaultRebootCmd   string = "shutdown -r now"
)

func DefaultCfg() Cfg {
	return Cfg{
		Shutdown: ShutdownCfg{
			Icon: "",
			Cmd:  defaultShutdownCmd,
		},
		Reboot: RebootCfg{
			Icon: "",
			Cmd:  defaultRebootCmd,
		},
	}
}
