package database

import (
	"log"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"gorm.io/gorm"
)

/*
	All functions are only intended to be used for the command line interface.
*/

// PasswordResetCLI CLI ONLY FUNCTION!
func (d *DB) PasswordResetCLI(userName string, password string) string {
	user := &models.User{}
	BADMSG := "There was a problem completing this request - please try again later."
	res := d.db.Where(&models.User{UserName: userName}).First(&user)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return "Invalid user - please try again."
	}
	if res.Error != nil {
		log.Println("PasswordResetCLI error: " + res.Error.Error())
		return BADMSG
	}

	pw, err := auth.HashPassword(password)
	if err != nil {
		log.Println("PasswordResetCLI error: " + res.Error.Error())
		return BADMSG
	}
	user.Password = pw

	d.db.Save(&user)

	return "Password updated."
}

// ListObjects CLI ONLY FUNCTION!
func (d *DB) ListObjects(objectType string) string {
	BADMSG := "There was a problem completing this request - please try again later."
	switch objectType {
	case "users":
		users := []models.User{}
		res := d.db.Find(&users)
		if res.Error != nil {
			return BADMSG
		}
		message := ""
		for _, k := range users {
			message += k.UserName + ":" + k.ID + ":" + k.FullName + "\n"
		}
		return message
	case "wishlists":
		wishlists := []models.Wishlist{}
		res := d.db.Find(&wishlists)
		if res.Error != nil {
			return BADMSG
		}
		message := ""
		for _, k := range wishlists {
			message += k.Name + ":" + k.ID + ":" + k.OwnerFullName + "\n"
		}
		return message
	default:
		return "Invalid list type selected."
	}
}
