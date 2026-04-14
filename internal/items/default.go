package items

import "aurevoir/internal/utils"

func DefaultOrder() []ItemId {
	return []ItemId{
		lockId,
		shutdownId,
		rebootId,
		logoutId,
		suspendId,
		hibernateId,
	}
}

func DefaultCmds() map[ItemId]string {
	return map[ItemId]string{
		lockId:      "loginctl lock-session",
		shutdownId:  "shutdown -h now",
		rebootId:    "shutdown -r now",
		logoutId:    "loginctl terminate-session $XDG_SESSION_ID",
		suspendId:   "systemctl suspend",
		hibernateId: "systemctl hibernate",
	}
}

func DefaultCfg() Cfg {
	defaultCmds := DefaultCmds()

	return Cfg{
		Order: buildDefaultOrder(),
		Lock: LockCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultCmds[lockId],
		},
		Shutdown: ShutdownCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultCmds[shutdownId],
		},
		Reboot: RebootCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultCmds[rebootId],
		},
		Logout: LogoutCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultCmds[logoutId],
		},
		Suspend: SuspendCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultCmds[suspendId],
		},
		Hibernate: HibernateCfg{
			Show: true,
			Icon: "",
			Cmd:  defaultCmds[hibernateId],
		},
	}
}

func buildDefaultOrder() OrderCfg {
	itemIds := DefaultOrder()

	itemIdsStr := make([]string, len(itemIds))
	for i, v := range itemIds {
		itemIdsStr[i] = string(v)
	}

	expectedItemIdsStr := utils.SliceToSet(itemIdsStr)

	expectedItemIds := make(map[ItemId]struct{}, len(expectedItemIdsStr))
	for k := range expectedItemIdsStr {
		expectedItemIds[ItemId(k)] = struct{}{}
	}

	return OrderCfg{
		expectedItemIds: expectedItemIds,
		ItemIds:         itemIds,
	}
}
