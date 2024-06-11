package components

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/make_request_cli/activities"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func CreateTemplatePage(app *tview.Application, pages *tview.Pages, db *gorm.DB) *tview.Flex {

	topView := tview.NewTextView().
		SetText(CenterTextVertically("A template is used to create a request. A template is just a reusable object that contains url, port number.", 3)).
		SetDynamicColors(true).SetTextAlign(tview.AlignCenter)

	bottomView := tview.NewTextView().
		SetText("To go back to the home page (ESCAPE) Key").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	urlInput := tview.NewInputField().
		SetLabel("URL: ").
		SetPlaceholder("Enter URL (i.e localhost)").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorYellow).
		SetLabelColor(tcell.ColorGreen).
		SetPlaceholderTextColor(tcell.ColorGray).
		SetFieldWidth(30)

	portInput := tview.NewInputField().
		SetLabel("Port Number: ").
		SetPlaceholder("Enter Port Number (i.e 8080)").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorYellow).
		SetLabelColor(tcell.ColorGreen).
		SetPlaceholderTextColor(tcell.ColorGray).
		SetFieldWidth(30)
	portInput.SetChangedFunc(func(text string) {
		if _, err := strconv.Atoi(text); err != nil && text != "" {
			ShowPopup(app, pages, "Invalid port number. Please enter a valid integer", func() {
				portInput.SetText("")
				app.SetFocus(portInput)
			})
		}
	})

	var checkedBox = false
	checkbox := tview.NewCheckbox().
		SetLabel("Check To Enable HTTPS: ").
		SetChecked(checkedBox).
		SetChangedFunc(func(checked bool) {
			checkedBox = checked
		})

	form := tview.NewForm().
		AddFormItem(urlInput).
		AddFormItem(portInput).
		AddFormItem(checkbox).
		AddButton("Create Template", func() {
			//Validation required before proceseeding.
			notEmptyUrl := urlInput.GetText() != ""
			notEmptyPort := portInput.GetText() != ""
			if !notEmptyPort || !notEmptyUrl {
				ShowPopup(app, pages, "There are missing required details.", func() {
					app.SetFocus(urlInput)
				})
			}

			db.Create(&activities.Template{
				URL:   urlInput.GetText(),
				PORT:  portInput.GetText(),
				HTTPS: checkedBox,
			})

			ShowPopup(app, pages, "You have successfully created a new template!", func() {
				//Set Checkbox to false
				checkbox.SetChecked(false)
				checkedBox = false
				//Clear inputs
				urlInput.SetText("")
				portInput.SetText("")
				app.SetFocus(urlInput)
				GoToPage(HOME_PAGE, app, pages, db, true)
			})
		})

	templatePage := tview.NewFlex().
		AddItem(topView, 3, 1, false).
		AddItem(form, 0, 1, true).
		AddItem(bottomView, 3, 1, false)

	templatePage.SetDirection(tview.FlexRow)
	templatePage.SetTitle("Create Template")
	templatePage.SetBorder(true)
	return templatePage
}
