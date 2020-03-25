package service

import (
	"aigram-backend/config"
	"aigram-backend/entities"
	"fmt"

	"github.com/jinzhu/gorm"
	// import dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// InitializeDatabase connects and automigrates the db
func InitializeDatabase(config *config.AppConfiguration) *gorm.DB {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name))
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database %s", err.Error()))
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Post{})

	Seed()
	return db
}

// Seed creates nessesary prefilled info
func Seed() {

}
