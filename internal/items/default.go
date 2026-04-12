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
			Show: true,
			Icon: "",
			Cmd:  defaultLockCmd,
		},
		Shutdown: ShutdownCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultShutdownCmd,
		},
		Reboot: RebootCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultRebootCmd,
		},
		Logout: LogoutCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultLogoutCmd,
		},
		Suspend: SuspendCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultSuspendCmd,
		},
		Hibernate: HibernateCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultHibernateCmd,
		},
	}
}
