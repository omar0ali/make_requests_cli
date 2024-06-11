package components

import (
	"encoding/json"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/make_request_cli/activities"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func UpdateHTTPRequest(app *tview.Application, pages *tview.Pages, db *gorm.DB) *tview.Flex {
	topView := tview.NewTextView().
		SetText(CenterTextVertically("HTTP Update Request.", 4)).
		SetDynamicColors(true).SetTextAlign(tview.AlignCenter)

	// First Select a request
	listRequest := createListRequests(db)
	// Prepare midview
	midView := tview.NewFlex()
	midView.SetDirection(tview.FlexRow)
	midViewTextDetials := tview.NewTextView()
	midViewTextDetials.SetDynamicColors(true).SetTextColor(tcell.ColorDarkGreen)
	midView.AddItem(midViewTextDetials, 1, 1, false)

	listRequest.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			selectedRequest = activities.GetRequests(db)[listRequest.GetCurrentItem()]
			midViewTextDetials.SetText(fmt.Sprintf("(*) %v Selected", selectedRequest))

			//DrawMidView.
			db.First(&selectedRequest.Template, "id = ?", selectedRequest.TemplateID)
			requestData := selectedRequest.DATA
			// Parse JSON data into a slice of strings
			var headerData []string
			if err := json.Unmarshal([]byte(requestData), &headerData); err != nil {
				fmt.Println("Error:", err)
				return event
			}
			form := tview.NewForm()
			form.SetBorder(true)
			form.SetBorderColor(tcell.ColorGreenYellow)
			midView.AddItem(form, 0, 1, true)
			for i := 0; i < len(headerData); i++ {
				input := tview.NewInputField().SetLabel(headerData[i])
				form.AddFormItem(input)
				listInputs = append(listInputs, input)
			}
			app.SetFocus(form)
			form.AddButton("Call Request", func() {
				//Loop through the inputs and get all the results.
				contentData := make(map[string]interface{})
				for i, v := range listInputs {
					contentData[headerData[i]] = v.GetText()
				}
				jsonData, err := json.Marshal(contentData)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				//We make the call
				//Call request.
				body, err := activities.HttpRequest(activities.UPDATE, activities.ConstructURL(selectedRequest), jsonData)
				if err != nil {
					ShowPopup(app, pages, fmt.Sprintf("Error: %v", err), func() {
						GoToPage(HOME_PAGE, app, pages, db, true)
					})
					return
				}
				//Display PopUp
				ShowPopup(app, pages, fmt.Sprintf("URL: %v\nBody: %v", activities.ConstructURL(selectedRequest), string(body)), func() {
					GoToPage(HOME_PAGE, app, pages, db, true)
				})
			})
		}
		return event
	})

	bottomView := tview.NewTextView().
		SetText(CenterTextVertically("To go back to the home page (ESCAPE) Key", 3)).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	view := tview.NewFlex().
		AddItem(topView, 4, 1, false).
		AddItem(listRequest, 10, 1, true).
		AddItem(midView, 0, 1, true).
		AddItem(bottomView, 3, 1, false)
	view.SetDirection(tview.FlexRow)
	view.SetBorder(true)
	view.SetTitle("HTTP UPDATE REQUEST")
	return view
}
