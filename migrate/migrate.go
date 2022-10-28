package main

import (
	"log"

	"github.com/00Duck/wishr-api/database"
	"github.com/00Duck/wishr-api/models"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	d := &database.DB{}
	d.Connect()

	migrateUser(d)
	log.Println("Finished AutoMigration")
}

func migrateUser(d *database.DB) error {
	err := d.AutoMigrate(&models.User{})
	return err
}
