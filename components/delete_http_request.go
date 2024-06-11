package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/make_request_cli/activities"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func DeleteHTTPRequest(app *tview.Application, pages *tview.Pages, db *gorm.DB) *tview.Flex {

	topView := tview.NewTextView().
		SetText(CenterTextVertically("HTTP Delete Request.", 4)).
		SetDynamicColors(true).SetTextAlign(tview.AlignCenter)

	listRequest := createListRequests(db)
	listRequest.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			//Select one request and get Template
			requests := activities.GetRequests(db)
			selectedRequest = requests[listRequest.GetCurrentItem()]
			db.First(&selectedRequest.Template, "id = ?", selectedRequest.TemplateID)
			//Call request.
			body, err := activities.HttpRequest(activities.DELETE, activities.ConstructURL(selectedRequest), nil)
			if err != nil {
				ShowPopup(app, pages, fmt.Sprintf("Err: %v", err), func() {
					GoToPage(HOME_PAGE, app, pages, db, true)
				})
				return event
			}
			ShowPopup(app, pages, fmt.Sprintf("URL: %v\nBody: %v", activities.ConstructURL(selectedRequest), string(body)), func() {
				GoToPage(HOME_PAGE, app, pages, db, true)
			})
		}
		return event
	})
	mainFlex := tview.NewFlex()
	mainFlex.SetBorder(true).SetTitle("DELETE HTTP REQUEST")
	mainFlex.SetDirection(tview.FlexRow)
	mainFlex.AddItem(topView, 4, 1, false).
		AddItem(listRequest, 0, 1, true)
	return mainFlex
}
