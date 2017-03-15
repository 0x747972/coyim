package gui

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/twstrike/coyim/i18n"
	"github.com/twstrike/gotk3adapter/gtki"
)

func authorizePresenceSubscriptionDialog(parent gtki.Window, from string) gtki.MessageDialog {
	builder := newBuilder("AuthorizeSubscription")

	confirmDialog := builder.getObj("dialog").(gtki.MessageDialog)
	text := fmt.Sprintf(i18n.Local("%s wants to talk to you. Is that ok?"), from)
	confirmDialog.SetProperty("text", text)
	confirmDialog.SetTransientFor(parent)

	return confirmDialog
}

type addContactDialog struct {
	builder                *builder
	dialog                 gtki.Window
	model                  gtki.ListStore
	accountInput           gtki.ComboBox
	contactInput           gtki.Entry
	notificationArea       gtki.Box
	notification           gtki.InfoBar
	subscriptionAskMessage gtki.TextBuffer
	nickname               gtki.Entry
	autoAuth               gtki.CheckButton
}

func (acd *addContactDialog) getVerifiedContact() (string, bool) {
	contact, _ := acd.contactInput.GetText()
	isJid, errmsg := verifyXmppAddress(contact)

	if !isJid {
		if acd.notification != nil {
			acd.notificationArea.Remove(acd.notification)
		}
		acd.notification = buildBadUsernameNotification(errmsg)
		acd.notificationArea.Add(acd.notification)
		acd.notification.ShowAll()
		log.Printf(errmsg)
		return "", false
	}

	return contact, true
}

func (acd *addContactDialog) getCurrentAccount() (string, error) {
	iter, err := acd.accountInput.GetActiveIter()
	if err != nil {
		return "", err
	}
	val, err := acd.model.GetValue(iter, 1)
	if err != nil {
		return "", err
	}
	return val.GetString()
}

func (acd *addContactDialog) getCurrentMessage() string {
	return acd.subscriptionAskMessage.GetText(
		acd.subscriptionAskMessage.GetStartIter(),
		acd.subscriptionAskMessage.GetEndIter(),
		false,
	)
}

func (acd *addContactDialog) getCurrentNickname() string {
	txt, _ := acd.nickname.GetText()
	return txt
}

func (acd *addContactDialog) getAutoAuthorize() bool {
	return acd.autoAuth.GetActive()
}

func (acd *addContactDialog) initAccounts(accounts []*account) {
	for _, acc := range accounts {
		iter := acd.model.Append()
		acd.model.SetValue(iter, 0, acc.session.GetConfig().Account)
		acd.model.SetValue(iter, 1, acc.session.GetConfig().ID())
	}

	if len(accounts) > 0 {
		acd.accountInput.SetActive(0)
	}
}

func (acd *addContactDialog) init() {
	acd.builder = newBuilder("AddContact")
	acd.builder.getItems(
		"AddContact", &acd.dialog,
		"accounts-model", &acd.model,
		"accounts", &acd.accountInput,
		"notification-area", &acd.notificationArea,
		"address", &acd.contactInput,
		"subscriptionAskMessage", &acd.subscriptionAskMessage,
		"nickname", &acd.nickname,
		"auto_authorize_checkbutton", &acd.autoAuth,
	)
}

func presenceSubscriptionDialog(accounts []*account, sendSubscription func(accountID, peer, msg, nick string, autoauth bool) error) gtki.Window {
	acd := &addContactDialog{}
	acd.init()
	acd.initAccounts(accounts)

	acd.builder.ConnectSignals(map[string]interface{}{
		"on_cancel_signal": func() {
			acd.dialog.Destroy()
		},
		"on_save_signal": func() {
			contact, ok := acd.getVerifiedContact()
			if !ok {
				return
			}

			accountID, err := acd.getCurrentAccount()
			if err != nil {
				log.Printf("Error encountered when getting account: %v", err)
				return
			}

			err = sendSubscription(accountID, contact, acd.getCurrentMessage(), acd.getCurrentNickname(), acd.getAutoAuthorize())
			if err != nil {
				log.Printf("Error encountered when sending subscription: %v", err)
				return
			}

			acd.dialog.Destroy()
			showSMPWizard()
		},
	})

	return acd.dialog
}

type smpWizard struct {
	builder *builder
	wizard  gtki.Assistant
	pin     gtki.Label
}

func showSMPWizard() {
	w := &smpWizard{}
	w.builder = newBuilder("SMPWizard")
	w.builder.getItems(
		"SMPWizard", &w.wizard,
		"PinLabel", &w.pin,
	)
	pin, err := createPIN()
	if err != nil {
		log.Printf("Cannot recover. Quitting SMP Wizard.")
		w.wizard.Destroy()
	}
	w.pin.SetText(pin)
	w.wizard.ShowAll()

	cur := w.wizard.GetCurrentPage()
	page, err := w.wizard.GetNthPage(cur)
	if err != nil {
		log.Printf("Error encountered when getting current page: %v", err)
	}
	w.wizard.SetPageComplete(page, true)

	w.builder.ConnectSignals(map[string]interface{}{
		"on_gen_pin": func() {
			pin, err := createPIN()
			if err != nil {
				log.Printf("Cannot recover. Quitting SMP Wizard.")
				w.wizard.Destroy()
			}
			w.pin.SetText(pin)
		},
		"on_close_signal": func() {
			w.wizard.Destroy()
		},
		"on_cancel_signal": func() {
			w.wizard.Destroy()
		},
		"on_escape_signal": func() {
			w.wizard.Destroy()
		},
	})
}

func createPIN() (string, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(1000000)))
	if err != nil {
		log.Printf("Error encountered when creating a new PIN: %v", err)
		return "", err
	}
	return fmt.Sprintf("%06d", val), nil
}
