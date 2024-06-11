package components

import (
	"encoding/json"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/make_request_cli/activities"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func CreateRequestPage(app *tview.Application, pages *tview.Pages, db *gorm.DB) *tview.Flex {

	var (
		selectedTemplate activities.Template
	)
	topView := tview.NewTextView().
		SetText(CenterTextVertically("A Request is used to make a HTTP request. A request is just a reusable object that contains all required resources to make an HTTP request. To shift foucs CTRL+R", 4)).
		SetDynamicColors(true).SetTextAlign(tview.AlignCenter)

	templateSelectedView := tview.NewTextView().
		SetDynamicColors(true).SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorGreenYellow)

	bottomView := tview.NewTextView().
		SetText(CenterTextVertically("To go back to the home page (ESCAPE) Key", 3)).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	nameInput := tview.NewInputField().
		SetLabel("Name i.e GetProjects").
		SetPlaceholder("Name").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorYellow).
		SetLabelColor(tcell.ColorGreen).
		SetPlaceholderTextColor(tcell.ColorGray).
		SetFieldWidth(30)

	pathInput := tview.NewInputField().
		SetLabel("Path").
		SetPlaceholder("i.e /projects").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorYellow).
		SetLabelColor(tcell.ColorGreen).
		SetPlaceholderTextColor(tcell.ColorGray).
		SetFieldWidth(30)

	//In the right side, we will need a list and input that take text and each will be inserted into the list.
	listData := tview.NewList()
	listData.ShowSecondaryText(false)
	listData.SetMainTextColor(tcell.ColorGreen)
	inputData := tview.NewInputField()
	inputDataTitle := tview.NewTextView().SetDynamicColors(true)
	inputDataTitle.SetText("Enter Data: ").SetTextColor(tcell.ColorYellow)
	inputData.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			listData.AddItem(inputData.GetText(), "", 0, nil)
			inputData.SetText("")
		}
		return event
	})
	//Form
	form := tview.NewForm().
		AddFormItem(nameInput).
		AddFormItem(pathInput).
		AddButton("Create Request", func() {
			//TODO: Validate Data...
			//Check if there is items more than 0 in the DATA.
			var data []string
			for i := 0; i < listData.GetItemCount(); i++ {
				itemtxt, _ := listData.GetItemText(i)
				data = append(data, itemtxt)
			}
			// Convert the list to JSON
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error converting list to JSON:", err)
				return
			}
			//TODO: Create a new request, DATA IS MISSING.
			db.Create(&activities.Request{
				TemplateID: selectedTemplate.ID,
				NAME:       nameInput.GetText(),
				PATH:       pathInput.GetText(),
				Template:   selectedTemplate,
				DATA:       string(jsonData),
			})
			ShowPopup(app, pages, "You have successfully created a new request!", func() {
				GoToPage(HOME_PAGE, app, pages, db, true)
			})
		})
	//Select a template from a list.
	list := getTemplates(db)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyEnter:
			selectedIndex := list.GetCurrentItem()
			selectedTemplate = activities.GetTemplates(db)[selectedIndex] //Getting the selected tempalte
			templateSelectedView.SetText(fmt.Sprintf("(*) Template Selected: %v", selectedTemplate))
			app.SetFocus(form)
		}

		return event
	})

	//RightFlex
	rightFlex := tview.NewFlex().
		AddItem(tview.NewTextView().SetTextColor(tcell.ColorYellow).SetText("Include data to be sent with this request (Optional):"), 1, 1, false).
		AddItem(listData, 0, 1, false).
		AddItem(inputDataTitle, 1, 1, false).
		AddItem(inputData, 1, 1, true)

	rightFlex.SetDirection(tview.FlexRow)
	rightFlex.SetTitle("Request Data (Optionl) CTRL+E")
	rightFlex.SetBorder(true)

	//LeftFlex
	leftFlex := tview.NewFlex().
		AddItem(topView, 4, 1, false).
		AddItem(tview.NewTextView().SetText("Select one of the following templates: ").SetTextColor(tcell.ColorYellow), 1, 1, false).
		AddItem(list, 0, 1, true).
		AddItem(form, 0, 1, true).
		AddItem(templateSelectedView, 1, 1, false).
		AddItem(bottomView, 3, 1, false)

	leftFlex.SetDirection(tview.FlexRow)
	leftFlex.SetTitle("Create Request CTRL+R")
	leftFlex.SetBorder(true)

	//MainFlex
	mainFlex := tview.NewFlex().
		AddItem(leftFlex, 0, 1, true).
		AddItem(rightFlex, 0, 1, false)

	mainFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlE:
			app.SetFocus(inputData)
		case tcell.KeyCtrlR:
			if selectedTemplate.ID > -0 {
				app.SetFocus(form)
			} else {
				app.SetFocus(list)
			}
		}

		return event
	})

	return mainFlex
}

func getTemplates(db *gorm.DB) *tview.List {
	data := activities.GetTemplates(db)
	list := tview.NewList()

	for _, v := range data {
		list.AddItem(fmt.Sprintf("(%d) URL: %v", v.ID, v.URL), fmt.Sprintf("PORT: %v HTTPS: %v", v.PORT, v.HTTPS), 0, nil)
	}

	return list
}
