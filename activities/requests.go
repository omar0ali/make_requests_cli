package activities

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Request struct {
	ID         uint
	NAME       string
	PATH       string
	TemplateID uint     // Assuming TEMPLATEID is the foreign key field
	Template   Template `gorm:"foreignKey:TemplateID"`
	DATA       string   `gorm:"type:text[]"`
}

func StartCreateRequest(db *gorm.DB) {
	ClearScreen()
	Display("Create Request", "Leave empty to cancel")
	name, err := Input("Name (i.e GetProjects)")
	if err != nil {
		log.Printf("Input error: %v\n", err)
		return
	}
	path, err := Input("Path")
	if err != nil {
		log.Printf("Input error: %v\n", err)
		return
	}

	listOfTemplates := GetTemplates(db)

	selection, err := MakeSelection(listOfTemplates)

	if err != nil {
		fmt.Println("Selection Failed. ", err)
		return
	}
	includeData := false
	//Will this request include data to be sent, or updating existing information
	if err := CreateDialogYesNo("Will this request include data to be sent?", func() error {
		includeData = true
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	var data []string
	template := listOfTemplates[selection]
	if includeData {
		Display(fmt.Sprintf("Selected: %v", template),
			"(*) Fields are case sensitive.",
			"(*) Please ensure that it matches the backend.",
			"(*) Leave empty to stop.",
			"(*) Type 'CANCEL' to Cancel.")
		for counter := 0; ; counter++ {
			item, err := Input(fmt.Sprintf(" (%v) ENTER", counter))
			if strings.EqualFold("CANCEL", item) {
				Display("Operation Canceled.")
				return
			}
			if item == "" {
				break
			}
			if err != nil {
				Display(err.Error())
				break
			}
			data = append(data, item)
		}
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to json.")
		return
	}

	err = db.Create(&Request{
		NAME:       name,
		PATH:       path,
		TemplateID: template.ID,
		Template:   template,
		DATA:       string(jsonData),
	}).Error
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}
}

func GetRequests(db *gorm.DB) []Request {
	ClearScreen()
	var requests []Request
	if err := db.Find(&requests).Error; err != nil {
		fmt.Println("Err: ", err)
	}
	return requests
}

func DeleteRequest(db *gorm.DB) {
	ClearScreen()
	Display("Delete Request")
	requests := GetRequests(db)
	//Select to delete template, one last option to cancel.
	selection, err := MakeSelection(requests)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	//Delete
	if err := CreateDialogYesNo(fmt.Sprintf("Selected %v will be deleted from the database.",
		requests[selection]), func() error {
		if err = db.Where("id = ?", requests[selection].ID).Delete(&Request{}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
