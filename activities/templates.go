package activities

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Template struct {
	ID    uint // Primary key, GORM will use it automatically
	URL   string
	PORT  string
	HTTPS bool
}

func StartCreateTemplate(db *gorm.DB) {
	ClearScreen()
	Display("Create Template", "Leave empty to cancel")
	url, err := Input("URL")
	if err != nil {
		log.Printf("Input error: %v\n", err)
		return
	}
	port, err := Input("Port Number")
	if err != nil {
		log.Printf("Input error: %v\n", err)
		return
	}
	https, err := Input("Https (y or n)")
	if err != nil {
		log.Printf("Input error: %v\n", err)
		return
	}
	var htps bool
	if strings.EqualFold(https, "y") || strings.EqualFold(https, "yes") {
		htps = true
	}

	db.Create(&Template{
		URL:   url,
		PORT:  port,
		HTTPS: htps,
	})
}

func GetTemplates(db *gorm.DB) []Template {
	var templates []Template
	if err := db.Find(&templates).Error; err != nil {
		fmt.Println("Err: ", err)
	}
	return templates
}

func DisplayTemplates(data []Template) {
	for index, templates := range data {
		fmt.Printf("(%v) URL:%v\n\tPORT:%v\n\tHTTP:%v\n", index, templates.URL, templates.PORT, templates.HTTPS)
	}
}

func DeleteTemplate(db *gorm.DB) {
	ClearScreen()
	Display("Delete Template", "Leave empty to cancel")
	//Display list of templates
	templates := GetTemplates(db)
	//Select to delete template, one last option to cancel.
	selection, err := MakeSelection(templates)
	fmt.Println(selection)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if err := CreateDialogYesNo(fmt.Sprintf("Selected %v will be deleted from the database.",
		templates[selection]), func() error {
		if err = db.Where("id=?", &templates[selection].ID).Delete(&Template{}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
