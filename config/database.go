package config

import (
	"fmt"
	"project-name/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB Context
var DB *gorm.DB

// Database Initialization
func Database() *gorm.DB {
	databaseUrl := LoadConfig().DatabaseURL

	host := LoadConfig().DatabaseHost
	user := LoadConfig().DatabaseUsername
	password := LoadConfig().DatabasePassword
	name := LoadConfig().DatabaseName
	port := LoadConfig().DatabasePort

	var err error
	dsn := databaseUrl
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	}
	fmt.Sprintln("dsn:", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if LoadConfig().EnableDatabaseAutomigration {
		err = DB.AutoMigrate(
			&models.User{},
		)
		if err != nil {
			fmt.Println(err)
			panic("Migration Failed")
		}
	}

	fmt.Println("Connected to Database:", LoadConfig().DatabaseName)

	return DB
}

func PathDb() string {
	return LoadConfig().PathDB
}
