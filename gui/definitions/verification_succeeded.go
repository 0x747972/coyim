package definitions

func init() {
	add(`VerificationSucceeded`, &defVerificationSucceeded{})
}

type defVerificationSucceeded struct{}

func (*defVerificationSucceeded) String() string {
	return `<interface>
  <object class="GtkDialog" id="dialog">
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <child internal-child="vbox">
      <object class="GtkBox" id="box">
        <property name="border-width">10</property>
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
        <child>
          <object  class="GtkImage">
            <property name="file">build/images/success.png</property>
          </object>
          <packing>
            <property name="padding">20</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="verification_message"/>
        </child>
        <child internal-child="action_area">
          <object class="GtkButtonBox" id="button_box">
            <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
            <child>
              <object class="GtkButton" id="button_ok">
                <property name="label" translatable="yes">OK</property>
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
