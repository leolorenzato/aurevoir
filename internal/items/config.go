package items

import (
	"aurevoir/internal/utils"
	"fmt"
	"strings"
)

const (
	lockId      ItemId = "lock"
	shutdownId  ItemId = "shutdown"
	rebootId    ItemId = "reboot"
	logoutId    ItemId = "logout"
	suspendId   ItemId = "suspend"
	hibernateId ItemId = "hibernate"
)

type ItemId string

type Cfg struct {
	Order     OrderCfg
	Lock      LockCfg
	Shutdown  ShutdownCfg
	Reboot    RebootCfg
	Logout    LogoutCfg
	Suspend   SuspendCfg
	Hibernate HibernateCfg
}

func (c *Cfg) MergeRaw(r RawCfg) {
	if r.Order != nil {
		c.Order.MergeRaw(*r.Order)
	}
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

func (c *Cfg) Validate() error {
	if err := c.Order.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %v", err)
	}

	return nil
}

type OrderCfg struct {
	expectedItemIds map[ItemId]struct{}
	ItemIds         []ItemId
}

func (c *OrderCfg) MergeRaw(r RawOrderCfg) {
	c.ItemIds = []ItemId{}
	for _, v := range r {
		c.ItemIds = append(c.ItemIds, ItemId(v))
	}
}

func (c *OrderCfg) Validate() error {
	if err := c.checkDuplicateItemIds(); err != nil {
		return fmt.Errorf("invalid order: %v", err)
	}

	if err := c.ensureExpectedItemIds(); err != nil {
		return fmt.Errorf("invalid order: %v", err)
	}

	return nil
}

func (c *OrderCfg) checkDuplicateItemIds() error {
	seen := make(map[ItemId]struct{})
	for _, el := range c.ItemIds {
		if _, ok := seen[el]; ok {
			return fmt.Errorf("duplicate item: %s", el)
		}
		seen[el] = struct{}{}
	}

	return nil
}

func (c *OrderCfg) ensureExpectedItemIds() error {
	itemIdsStr := make([]string, len(c.ItemIds))
	for i, v := range c.ItemIds {
		itemIdsStr[i] = string(v)
	}

	expectedItemIdsStr := make(map[string]struct{}, len(c.expectedItemIds))
	for k := range c.expectedItemIds {
		expectedItemIdsStr[string(k)] = struct{}{}
	}

	missingItems := utils.GetMissingKeys(expectedItemIdsStr, utils.SliceToSet(itemIdsStr))
	if len(missingItems) > 0 {
		return fmt.Errorf("missing items [%s]", strings.Join(missingItems, ", "))
	}

	return nil
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
