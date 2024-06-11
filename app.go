package main

import (
	"github.com/omar0ali/make_request_cli/database"
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
	//Run Application
	RunApp(OLD)
}
