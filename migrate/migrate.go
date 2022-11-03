package main

import (
	"log"

	"github.com/00Duck/wishr-api/database"
	"github.com/00Duck/wishr-api/models"
)

func main() {
	d := &database.DB{}
	d.Connect()

	migrate(d)
	log.Println("Finished AutoMigration")
}

func migrate(d *database.DB) error {
	err := d.AutoMigrate(&models.User{}, &models.Wishlist{}, &models.WishlistItem{}, &models.Session{})
	return err
}
