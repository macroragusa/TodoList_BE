package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"main/config"
)

// declare DB connector object
var DB *gorm.DB

func ConnectDataBase() {
	// open a db connection
	database, err := gorm.Open(
		"postgres",
		config.ConnectionString,
	)

	if err != nil {
		panic("failed to connect database\n" + err.Error())
	}

	// migrate the schema
	database.AutoMigrate(&Todos{})

	if config.DebugMode {
		// assign connection to the DB connector object with debug mode
		DB = database.LogMode(true)
	} else {
		DB = database
	}

}
