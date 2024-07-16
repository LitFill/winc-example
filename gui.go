package main

import "github.com/tadvi/winc"

const (
	WIN_WIDTH  = 800
	WIN_HEIGHT = 600
)

func exitApp(arg *winc.Event) {
	winc.Exit()
}

func runGui() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(WIN_WIDTH, WIN_HEIGHT)
	mainWindow.SetText("LitFill's new Windows window")

	menuAtas := mainWindow.NewMenu()

	fileMn := menuAtas.AddSubMenu("File")
	fileMn.AddItem("New", winc.Shortcut{
		Modifiers: winc.ModControl,
		Key:       winc.KeyN,
	})
	exitMn := fileMn.AddItem("Exit", winc.Shortcut{
		Modifiers: winc.ModControl,
		Key:       winc.KeyQ,
	})

	editMn := menuAtas.AddSubMenu("Edit")
	cutMn := editMn.AddItem("Cut", winc.Shortcut{
		Modifiers: winc.ModControl,
		Key:       winc.KeyX,
	})
	copyMn := editMn.AddItem("Copy", winc.NoShortcut)
	pasteMn := editMn.AddItem("Paste", winc.NoShortcut)

	menuAtas.Show()

	exitMn.OnClick().Bind(exitApp)

	copyMn.SetCheckable(true)
	copyMn.SetChecked(true)
	pasteMn.SetEnabled(false)

	cutMn.OnClick().Bind(func(arg *winc.Event) {
		winc.MsgBoxOk(mainWindow, "Cut", "Click Event.")
	})

	edt := winc.NewEdit(mainWindow)
	edt.SetPos(10, 20)
	edt.SetText("edit this text: ")
	edt.OnChange().Bind(func(arg *winc.Event) {
		w, h := edt.Size()
		w++
		h++
		arg.Sender.SetSize(w, h)
	})

	btn := winc.NewPushButton(mainWindow)
	btn.SetText("show or hide")
	btn.SetPos(40, 50)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(func(arg *winc.Event) {
		if edt.Visible() {
			edt.Hide()
		} else {
			edt.Show()
		}
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(exitApp)

	winc.RunMainLoop()
}
