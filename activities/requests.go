package activities

import (
	"encoding/json"
	"fmt"
	"log"

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
	name, err := Input("Name (i.e GetProjects)")
	if err != nil {
		log.Fatalf("Input error: %v\n", err)
	}
	path, err := Input("Path")
	if err != nil {
		log.Fatalf("Input error: %v\n", err)
	}

	listOfTemplates := GetTemplates(db)

	selection, err := MakeSelection(listOfTemplates)

	if err != nil {
		fmt.Println("Selection Failed. ", err)
		return
	}

	template := listOfTemplates[selection]
	Display(fmt.Sprintf("You have selected %v", template))
	fmt.Println("*Filed names are case sensitive; please ensure that it matches the backend.")
	var data []string
	for counter := 0; ; counter++ {
		item, err := Input(fmt.Sprintf("%v: Leave empty to stop\nEnter Filed Name", counter))
		if err != nil {
			fmt.Println("Error while entering data.")
			break
		}
		if item == "" {
			break
		}
		data = append(data, item)
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
	var requests []Request
	if err := db.Find(&requests).Error; err != nil {
		fmt.Println("Err: ", err)
	}
	return requests
}

func DeleteRequest(db *gorm.DB) {
	//Display list of requests
	requests := GetRequests(db)
	//Select to delete template, one last option to cancel.
	selection, err := MakeSelection(requests)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	//Delete
	if err := CreateDialogYesNo(fmt.Sprintf("Are you sure?\nSelected %v will be deleted from the database.",
		requests[selection]), func() error {
		if err = db.Where("id = ?", requests[selection].ID).Delete(&Request{}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
