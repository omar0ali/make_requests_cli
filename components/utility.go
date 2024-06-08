package components

import (
	"strings"

	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func GoToPage(page string, app *tview.Application, pages *tview.Pages, db *gorm.DB) {
	//Construct pages
	homePage := HomePage(app, pages, db)
	templatePage := CreateTemplatePage(app, pages, db)
	requestPage := CreateRequestPage(app, pages, db)
	postHttpPage := PostHTTPRequest(app, pages, db)
	//Add all pages
	pages.AddPage("HomePage", homePage, true, true)
	pages.AddPage("TemplatePage", templatePage, true, false)
	pages.AddPage("RequestPage", requestPage, true, false)
	pages.AddPage("PostPage", postHttpPage, true, false)
	if page != "" {
		pages.SwitchToPage(page)
	}
}

func CenterTextVertically(text string, height int) string {
	lines := strings.Split(text, "\n")
	lineCount := len(lines)
	paddingLines := (height - lineCount) / 2

	padding := strings.Repeat("\n", paddingLines)
	return padding + text + padding
}
