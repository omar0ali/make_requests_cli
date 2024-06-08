package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/make_request_cli/activities"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func HomePage(app *tview.Application, pages *tview.Pages, db *gorm.DB) *tview.Flex {
	leftTitle := tview.NewTextView().
		SetText(CenterTextVertically("MAIN MENU", 4)).
		SetTextColor(tcell.ColorGreen).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
		//------------------------------------------------------------------//
	rightTitleTop := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(readTemplates(db), 0, 1, true)
	rightTitleTop.SetBorder(true).
		SetTitle("Templates Details - CTRL+T").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tcell.ColorYellow)
		//------------------------------------------------------------------//

	rightTitleBottom := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(readRequests(db), 0, 1, true)
	rightTitleBottom.SetBorder(true).
		SetTitle("Requests Details - CTRL+R").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tcell.ColorYellow)

		//------------------------------------------------------------------//
	list := tview.NewList().
		AddItem("Create Template",
			"Template requires a url, port number. Will be used to create Requests.", '1', func() {
				pages.SwitchToPage("TemplatePage")
			}).
		AddItem("Create Request",
			"Create Request, must be already created template to choose from.", '2', func() {
				pages.SwitchToPage("RequestPage")
			}).
		AddItem("Post HTTP Request",
			"Make a post request. Note: Must already have Post requests created.", '3', func() {
				pages.SwitchToPage("PostPage")
			}).
		AddItem("Get HTTP Request",
			"Make a get request. Note: Must already have Get requests created.", '4', func() {
				pages.SwitchToPage("delete_request")
			}).
		AddItem("Delete HTTP Request",
			"Make a delete request.", '5', func() {
				pages.SwitchToPage("delete_request")
			}).
		AddItem("Update HTTP Request",
			"Make an update request.", '6', func() {
				pages.SwitchToPage("delete_request")
			}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	subTitle := tview.NewTextView().
		SetText("CLI tool built in Go that allows users to interactively create and manage HTTP requests.\nTo switch back to the main menu. Tab Key").
		SetDynamicColors(true).SetTextAlign(tview.AlignCenter)

	//------------------------------------------------------------------//
	boxMidText := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
			SetText("Select an option from 1 to 6, you can:\n(*) Create, read, and delete templates.\n(*) Create, read, and delete requests.\n(*) Send HTTP POST, GET, and DELETE requests.\n(*) Interactive user interface for easy operation.").SetTextColor(tcell.ColorGreenYellow), 0, 1, true)
	boxMidText.SetBorder(true).
		SetTitle("Instructions").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tcell.ColorDarkGreen)
		//------------------------------------------------------------------//

	bottomText := tview.NewTextView().
		SetText("To close the application (CTRL+C) or (q)").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	// Left side flex container
	leftFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(leftTitle, 4, 1, false).
		AddItem(subTitle, 4, 1, false).
		AddItem(boxMidText, 7, 1, false).
		AddItem(list, 0, 1, true).
		AddItem(bottomText, 2, 1, false)
	leftFlex.SetBorder(true)
	leftFlex.SetTitle("Home Page - Tab Key")

	// Right side flex container (split horizontally)
	rightFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(rightTitleTop, 0, 1, false).
		AddItem(rightTitleBottom, 0, 1, false)
	rightFlex.SetTitle("Display Details")

	// Main flex container (split vertically)
	mainFlex := tview.NewFlex().
		AddItem(leftFlex, 0, 1, true).
		AddItem(rightFlex, 0, 1, false)

	mainFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTAB:
			app.SetFocus(leftFlex)
		case tcell.KeyCtrlT:
			app.SetFocus(rightTitleTop)
		case tcell.KeyCtrlR:
			app.SetFocus(rightTitleBottom)
		}
		return event
	})

	return mainFlex
}

func readTemplates(db *gorm.DB) *tview.Flex {
	title := tview.NewTextView().
		SetText("Here are all the created templates. To scroll through them, you need to shift focus. Press CTRL+T to switch focus to the Templates.\n-----------").
		SetDynamicColors(true).SetTextColor(tcell.ColorGreenYellow)

	templates := activities.GetTemplates(db)
	list := tview.NewList()
	if len(templates) > 0 {
		for _, template := range templates {
			list.AddItem(fmt.Sprintf("%d URL: %s", template.ID, template.URL), fmt.Sprintf("(*) Port: %v HTTPS: %v", template.PORT, template.HTTPS), 0, nil)
		}
	} else {
		list.AddItem("Empty list", "", 0, nil)
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlD:
			//We need to remove it from the actuall database.
			list.RemoveItem(list.GetCurrentItem())
		}
		return event
	})

	readTemplates := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(title, 3, 1, false).
		AddItem(list, 0, 1, true)

	return readTemplates
}

func readRequests(db *gorm.DB) *tview.Flex {
	title := tview.NewTextView().
		SetText("Here are all the created requests. To scroll through them, you need to shift focus. Press CTRL+R to switch focus to the Requests.\n-----------").
		SetDynamicColors(true).SetTextColor(tcell.ColorGreenYellow)

	requests := activities.GetRequests(db)
	list := tview.NewList()
	if len(requests) > 0 {
		for _, request := range requests {
			list.AddItem(fmt.Sprintf("%d Name: %s", request.ID, request.NAME), fmt.Sprintf("(*) Data: %v Path: %v", request.DATA, request.PATH), 0, nil)
		}
	} else {
		list.AddItem("Empty list", "", 0, nil)
	}

	readRequests := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(title, 3, 1, false).
		AddItem(list, 0, 1, true)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlD:
			//We need to remove it from the actuall database.
			list.RemoveItem(list.GetCurrentItem())
		}
		return event
	})

	return readRequests
}
