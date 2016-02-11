package gui

import (
	gtk "github.com/gotk3/gotk3/gtk/iface"
	"github.com/twstrike/coyim/config"
)

type viewMenu struct {
	merge   gtk.CheckMenuItem
	offline gtk.CheckMenuItem
}

func (v *viewMenu) setFromConfig(c *config.ApplicationConfig) {
	doInUIThread(func() {
		v.merge.SetActive(c.Display.MergeAccounts)
		v.offline.SetActive(!c.Display.ShowOnlyOnline)
	})
}

func (u *gtkUI) toggleMergeAccounts() {
	if u.config != nil {
		u.config.Display.MergeAccounts = u.viewMenu.merge.GetActive()
		u.saveConfigOnly()
	}

	u.roster.redraw()
}

func (u *gtkUI) toggleShowOffline() {
	if u.config != nil {
		u.config.Display.ShowOnlyOnline = !u.viewMenu.offline.GetActive()
		u.saveConfigOnly()
	}

	u.roster.redraw()
}
