package items

const (
	defaultLockCmd      string = "loginctl lock-session"
	defaultShutdownCmd  string = "shutdown -h now"
	defaultRebootCmd    string = "shutdown -r now"
	defaultLogoutCmd    string = "loginctl terminate-session $XDG_SESSION_ID"
	defaultSuspendCmd   string = "systemctl suspend"
	defaultHibernateCmd string = "systemctl hibernate"
)

func DefaultCfg() Cfg {
	return Cfg{
		Lock: LockCfg{
			Icon: "",
			Cmd:  defaultLockCmd,
		},
		Shutdown: ShutdownCfg{
			Icon: "",
			Cmd:  defaultShutdownCmd,
		},
		Reboot: RebootCfg{
			Icon: "",
			Cmd:  defaultRebootCmd,
		},
		Logout: LogoutCfg{
			Icon: "",
			Cmd:  defaultLogoutCmd,
		},
		Suspend: SuspendCfg{
			Icon: "",
			Cmd:  defaultSuspendCmd,
		},
		Hibernate: HibernateCfg{
			Icon: "",
			Cmd:  defaultHibernateCmd,
		},
	}
}
