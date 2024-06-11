package components

import (
	"strings"

	"github.com/rivo/tview"
	"gorm.io/gorm"
)

const (
	HOME_PAGE     = "HomePage"
	TEMPLATE_PAGE = "TemplatePage"
	REQUEST_PAGE  = "RequestPage"
	POST_PAGE     = "PostPage"
	GET_PAGE      = "GetPage"
	DELETE_PAGE   = "DeletePage"
	UPDATE_PAGE   = "UpdatePage"
)

func GoToPage(page string, app *tview.Application, pages *tview.Pages, db *gorm.DB, refresh bool) {
	if refresh {
		refreshPage(app, pages, db)
	}
	pages.SwitchToPage(page)
}

func refreshPage(app *tview.Application, pages *tview.Pages, db *gorm.DB) {
	//Construct pages
	homePage := HomePage(app, pages, db)
	templatePage := CreateTemplatePage(app, pages, db)
	requestPage := CreateRequestPage(app, pages, db)
	postHttpPage := PostHTTPRequest(app, pages, db)
	getHttpPage := GetHTTPRequest(app, pages, db)
	deleteHttpPage := DeleteHTTPRequest(app, pages, db)
	updateHttpPage := UpdateHTTPRequest(app, pages, db)
	//Add all pages
	pages.AddPage(HOME_PAGE, homePage, true, true)
	pages.AddPage(TEMPLATE_PAGE, templatePage, true, false)
	pages.AddPage(REQUEST_PAGE, requestPage, true, false)
	pages.AddPage(POST_PAGE, postHttpPage, true, false)
	pages.AddPage(GET_PAGE, getHttpPage, true, false)
	pages.AddPage(DELETE_PAGE, deleteHttpPage, true, false)
	pages.AddPage(UPDATE_PAGE, updateHttpPage, true, false)
}

func CenterTextVertically(text string, height int) string {
	lines := strings.Split(text, "\n")
	lineCount := len(lines)
	paddingLines := (height - lineCount) / 2
	padding := strings.Repeat("\n", paddingLines)
	return padding + text + padding
}
