package components

import "github.com/rivo/tview"

// Function to show a popup modal with a given message
func ShowPopup(app *tview.Application, pages *tview.Pages, message string, btn func()) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("modal")
			btn()
		})

	// Add the modal to the pages
	pages.AddPage("modal", modal, true, true)
	app.SetFocus(modal)
}
