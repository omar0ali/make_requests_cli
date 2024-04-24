package main

import (
	"fmt"
	"strings"

	activities "github.com/omar0ali/make_request_cli/activities"
	"github.com/omar0ali/make_request_cli/database"
	models "github.com/omar0ali/make_request_cli/models"
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
		models.CreateItem("Read   Templates"),
		models.CreateItem("Delete Template"),
		models.CreateItem("Create Request"),
		models.CreateItem("Read   Requests"),
		models.CreateItem("Delete Request"),
		models.CreateItem("Post   HTTP request"),
		models.CreateItem("Get    HTTP request"),
		models.CreateItem("Delete HTTP request"),
	}, func(item models.MenuItem) bool {
		var builder strings.Builder
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
				fmt.Println("We will create a template.")
				activities.StartCreateTemplate(database.DB)
			case 2:
				data := activities.GetTemplates(database.DB)
				for index, templates := range data {
					builder.WriteString(fmt.Sprintf("%v: URL:%v PORT:%v HTTP:%v\n", index, templates.URL, templates.PORT, templates.HTTPS))
				}
				activities.Display(builder.String())
			case 3:
				activities.Display("Delete Template")
				activities.DeleteTemplate(database.DB)
			case 4:
				activities.Display("Create Request")
				activities.StartCreateRequest(database.DB)
			case 5:
				data := activities.GetRequests(database.DB)
				for index, request := range data {
					builder.WriteString(fmt.Sprintf("%v:Path: %v\tData: %v\n", index, request.PATH, request.DATA))
				}
				activities.Display(builder.String())
			case 6:
				activities.Display("Delete Request")
				activities.DeleteRequest(database.DB)
			case 7:
				activities.Display("POST HTTP Request")
				activities.CreatePostRequest(database.DB)
			case 8:
				activities.Display("GET HTTP Request")
				activities.CreateGetRequest(database.DB)
			case 9:
				activities.Display("Delete HTTP Request")
				activities.CreatePostRequest(database.DB)
			}
		})
		return exitSignal
	})
}
