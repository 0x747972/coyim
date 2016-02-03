package definitions

func init() {
	add(`Main`, &defMain{})
}

type defMain struct{}

func (*defMain) String() string {
	return `
<interface>
  <object class="GtkApplicationWindow" id="mainWindow">
    <property name="window-position">GTK_WIN_POS_NONE</property>
    <property name="default-height">600</property>
    <property name="default-width">200</property>
    <property name="title">CoyIM</property>
    <signal name="destroy" handler="on_close_window_signal" />
    <!-- <property name="icon">we dont know how to use it now</property> -->
    <child>
      <object class="GtkBox" id="Vbox">
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
        <child>
          <object class="GtkMenuBar" id="menubar">
            <child>
              <object class="GtkMenuItem" id="ContactsMenu">
                <property name="label" translatable="yes">Contacts</property>
                <child type="submenu">
                  <object class="GtkMenu" id="menu">
                    <child>
                      <object class="GtkMenuItem" id="addMenu">
                        <property name="label" translatable="yes">Add...</property>
                        <signal name="activate" handler="on_add_contact_window_signal" />
                      </object>
                    </child>
                  </object>
                </child>
              </object>
            </child>
            <child>
              <object class="GtkMenuItem" id="AccountsMenu">
                <property name="label" translatable="yes">Accounts</property>
              </object>
            </child>
            <child>
              <object class="GtkMenuItem" id="ViewMenu">
                <property name="label" translatable="yes">View</property>
                <child type="submenu">
                  <object class="GtkMenu" id="menu2">
                    <child>
                      <object class="GtkCheckMenuItem" id="CheckItemMerge">
                        <property name="label" translatable="yes">Merge Accounts</property>
                        <signal name="toggled" handler="on_toggled_check_Item_Merge_signal" />
                      </object>
                    </child>
                    <child>
                      <object class="GtkCheckMenuItem" id="CheckItemShowOffline">
                        <property name="label" translatable="yes">Show Offline Contacts</property>
                        <signal name="toggled" handler="on_toggled_check_Item_Show_Offline_signal" />
                      </object>
                    </child>
                  </object>
                </child>
              </object>
            </child>
            <child>
              <object class="GtkMenuItem" id="HelpMenu">
                <property name="label" translatable="yes">Help</property>
                <child type="submenu">
                  <object class="GtkMenu" id="menu3">
                    <child>
                      <object class="GtkMenuItem" id="feedbackMenu">
                        <property name="label" translatable="yes">Feedback</property>
                        <signal name="activate" handler="on_feedback_dialog_signal" />
                      </object>
                    </child>
                    <child>
                      <object class="GtkMenuItem" id="aboutMenu">
                        <property name="label" translatable="yes">About</property>
                        <signal name="activate" handler="on_about_dialog_signal" />
                      </object>
                    </child>
                  </object>
                </child>
              </object>
            </child>
          </object>
        </child>
        <child>
          <object class="GtkBox" id="notification-area">
            <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">false</property>
            <property name="pack-type">GTK_PACK_END</property>
            <property name="position">0</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>

`
}
