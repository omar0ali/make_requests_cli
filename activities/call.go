package activities

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func createRequest(db *gorm.DB, method string) {
	//List all requests
	requests := GetRequests(db)
	if len(requests) < 1 {
		fmt.Println("There is no requests.")
		return
	}
	choice, err := MakeSelection(requests)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ClearScreen()
	//Select one request
	selectedRequest := requests[choice]
	db.First(&selectedRequest.Template, "id = ?", selectedRequest.TemplateID)
	//Prepare the request
	requestData := selectedRequest.DATA
	// Parse JSON data into a slice of strings
	var headerData []string
	if err := json.Unmarshal([]byte(requestData), &headerData); err != nil {
		fmt.Println("Error:", err)
		return
	}
	contentData := make(map[string]interface{})
	for _, element := range headerData {
		content, err := Input(fmt.Sprintf("Enter %v", element))
		if err != nil {
			fmt.Println(err)
			return
		}
		contentData[element] = content
	}
	jsonData, err := json.Marshal(contentData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	Display(string(jsonData))
	//Call request.
	body, err := HttpRequest(method, ConstructURL(selectedRequest), jsonData)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	Display(ConstructURL(selectedRequest))
	Display(string(body))
}

func CreatePostRequest(db *gorm.DB) {
	Display("POST HTTP Request")
	createRequest(db, POST)
}

func CreateUpdateRequest(db *gorm.DB) {
	Display("UPDATE HTTP Request")
	createRequest(db, UPDATE)
}

func CreateDeleteRequest(db *gorm.DB) {
	ClearScreen()
	Display("DELETE HTTP Request")
	requests := GetRequests(db)
	if len(requests) < 1 {
		fmt.Println("There is no requests.")
		return
	}
	choice, err := MakeSelection(requests)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ClearScreen()

	//Select one request and get Template
	selectedRequest := requests[choice]
	db.First(&selectedRequest.Template, "id = ?", selectedRequest.TemplateID)

	body, err := HttpRequest(DELETE, ConstructURL(selectedRequest), nil)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	Display(ConstructURL(selectedRequest))
	Display(string(body))
}

func CreateGetRequest(db *gorm.DB) {
	ClearScreen()
	Display("GET HTTP Request")
	//List all requests
	requests := GetRequests(db)
	if len(requests) < 1 {
		fmt.Println("There is no requests.")
		return
	}
	choice, err := MakeSelection(requests)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ClearScreen()
	//Select one request and get Template
	selectedRequest := requests[choice]
	db.First(&selectedRequest.Template, "id = ?", selectedRequest.TemplateID)
	//Call request.
	body, err := HttpRequest(GET, ConstructURL(selectedRequest), nil)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	Display(ConstructURL(selectedRequest))
	Display(string(body))
}

func ConstructURL(req Request) string {
	PORT := req.Template.PORT
	URL := req.Template.URL
	HTTPValue := func() string {
		if req.Template.HTTPS {
			return "https"
		}
		return "http"
	}()
	PATH := req.PATH
	return fmt.Sprintf("%v://%v:%v/%v", HTTPValue, URL, PORT, PATH)
}
