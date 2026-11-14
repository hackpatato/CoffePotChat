package client

import "fyne.io/fyne/app"


func GuiChat() {
	App := app.New()
	Window := App.NewWindow("Coffe Pot Chat")
	Window.ShowAndRun()
}