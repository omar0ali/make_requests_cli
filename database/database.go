package database

import (
	"fmt"

	"github.com/omar0ali/make_request_cli/activities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting database")
		panic(err.Error())
	}

	DB = db
	DB.AutoMigrate(&activities.Template{}, &activities.Request{})
}
