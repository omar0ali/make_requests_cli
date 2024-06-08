package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	activities "github.com/omar0ali/make_request_cli/activities"
	components "github.com/omar0ali/make_request_cli/components"
	"github.com/omar0ali/make_request_cli/database"
	models "github.com/omar0ali/make_request_cli/models"
	"github.com/rivo/tview"
)

var (
	pages *tview.Pages
)

func main() {
	//connection to database.
	database.Connect()
	sqlDB, err := database.DB.DB()
	if err != nil {
		panic("Problem: with getting db.")
	}
	defer sqlDB.Close()

	//starting the app with menu.
	activities.Draw([]models.MenuItem{
		models.CreateItem("Quit"),
		models.CreateItem("Create Template"),
		models.CreateItem("Create Request"),
		models.CreateItem("Read   Templates"),
		models.CreateItem("Read   Requests"),
		models.CreateItem("Delete Template"),
		models.CreateItem("Delete Request"),
		models.CreateItem("Post   HTTP request"),
		models.CreateItem("Get    HTTP request"),
		models.CreateItem("Delete HTTP request"),
		models.CreateItem("Update HTTP request"),
	}, func(item models.MenuItem) bool {
		exitSignal := false
		item.OnClick(func() {
			switch item.ID {
			case 0:
				if err := activities.CreateDialogYesNo("Are you sure?", func() error {
					exitSignal = true
					return nil
				}); err != nil {
					fmt.Println(err)
				}
			case 1:
				activities.StartCreateTemplate(database.DB)
			case 2:
				activities.StartCreateRequest(database.DB)
			case 3:
				activities.ClearScreen()
				data := activities.GetTemplates(database.DB)
				activities.DisplayTemplates(data)
			case 4:
				activities.ClearScreen()
				data := activities.GetRequests(database.DB)
				activities.DisplayRequests(data)
			case 5:
				activities.DeleteTemplate(database.DB)
			case 6:
				activities.DeleteRequest(database.DB)
			case 7:
				activities.CreatePostRequest(database.DB)
			case 8:
				activities.CreateGetRequest(database.DB)
			case 9:
				activities.CreateDeleteRequest(database.DB)
			case 10:
				activities.CreateUpdateRequest(database.DB)
			}
		})
		return exitSignal
	})

	//Starting the application using TView
	app := tview.NewApplication()
	pages = tview.NewPages()
	components.GoToPage("", app, pages, database.DB)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			components.GoToPage("HomePage", app, pages, database.DB)
		}
		return event
	})
	// if err := app.SetRoot(pages, true).Run(); err != nil {
	// 	panic(err)
	// }
}
