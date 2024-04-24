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
	url, err := Input("URL")
	if err != nil {
		log.Fatalf("Input error: %v\n", err)
	}
	port, err := Input("Port Number")
	if err != nil {
		log.Fatalf("Input error: %v\n", err)
	}
	https, err := Input("Https (y or n)")
	if err != nil {
		log.Fatalf("Input error: %v\n", err)
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

func DeleteTemplate(db *gorm.DB) {
	//Display list of templates
	templates := GetTemplates(db)
	//Select to delete template, one last option to cancel.
	selection, err := MakeSelection(templates)
	fmt.Println(selection)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if err := CreateDialogYesNo(fmt.Sprintf("Are you sure?\nSelected %v will be deleted from the database.", templates[selection]), func() error {
		if err = db.Where("id=?", &templates[selection].ID).Delete(&Template{}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
