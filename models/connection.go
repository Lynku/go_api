package models

import (
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var IT string

func ConnectDataBase() {

	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	IT = "TextInput, NumberInput, PhoneInput, EmailInput, UrlInput, DateInput, ColorInput, TextareInput, EditorInput, DropdownInput, RadioInput, CheckboxInput, Repeater, ImageInput"

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.Migrator().CreateTable(&Settings{}, &Users{}, &Tokens{}, &PageType{}, &Page{}, &PageTypeOptions{}, &PageCf{}, &Media{})

	// Seed
	result := database.First(&Settings{})
	if uint(result.RowsAffected) == 0 {
		seed(database)
		testData(database)
	}

	DB = database
}

func seed(db *gorm.DB) {

	db.Save(&Settings{
		Name:  "email",
		Value: "admin@web.com",
	})
	admin := db.Save(&Users{
		Name:     "root",
		Email:    "admin@web.com",
		Password: "123",
		Status:   0,
	})

	static := db.Save(&PageType{Name: "Static"})

	homePage := Page{
		PageTypeId: uint(static.RowsAffected),
		Author:     uint(admin.RowsAffected),
		Title:      "Home",
		Content:    "Hello World!",
		Status:     1,
	}
	db.Save(&homePage)
}

func testData(dbn *gorm.DB) {
	dbn.Create(&PageType{Name: "Full"})

	filedsArr := strings.Split(IT, ",")
	testFields := []PageTypeOptions{}
	for _, field := range filedsArr {
		testFields = append(testFields,
			PageTypeOptions{
				PageTypeId: 2,
				Type:       strings.TrimSpace(field),
				Name:       strings.TrimSpace(field),
			})
	}
	dbn.Save(&testFields)

	fullPage := Page{
		PageTypeId: 2,
		Author:     0,
		Title:      "Full",
		Content:    "Full...",
		Status:     1,
	}
	dbn.Save(&fullPage)
}
