package database

import (
	"errors"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
	All functions are only intended to be used for the command line interface.
*/

var cliGenErrMsg = "There was a problem completing this request - please try again later."

func (d *DB) getUserForCLI(userName string, user *models.User) error {
	res := d.db.Where(&models.User{UserName: userName}).First(&user)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return errors.New("record not found")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// PasswordResetCLI CLI ONLY FUNCTION!
func (d *DB) PasswordResetCLI(userName string, password string) string {
	user := models.User{}
	err := d.getUserForCLI(userName, &user)
	if err != nil {
		if err.Error() == "record not found" {
			return "Invalid user - please try again."
		}
		return cliGenErrMsg
	}

	pw, err := auth.HashPassword(password)
	if err != nil {
		log.Println("PasswordResetCLI error: " + err.Error())
		return cliGenErrMsg
	}
	user.Password = pw

	d.db.Save(&user)

	return "Password updated."
}

// GeneratePasswordResetRequest CLI ONLY FUNCTION!
func (d *DB) GeneratePasswordResetRequest(userName string) string {
	user := models.User{}
	err := d.getUserForCLI(userName, &user)
	if err != nil {
		if err.Error() == "record not found" {
			return "Invalid user - please try again."
		}
		return cliGenErrMsg
	}
	expiration := time.Now().Add(time.Hour * 24) // 1 day from now

	user.ResetToken = uuid.New().String()
	user.ResetTokenExpiration = &expiration

	d.db.Save(&user)
	webHost := os.Getenv("HOST_NAME")
	if webHost == "" {
		webHost = "https://<host_name>/"
	}
	retURL, err := url.JoinPath(webHost, "/passwordreset/"+user.ResetToken)
	if err != nil {
		log.Println("GeneratePasswordResetRequest error: " + err.Error())
		return cliGenErrMsg
	}

	return "Reset token generated. Instruct user to visit the following" +
		" link to reset their password (expires 24h).\n\n" + retURL
}

// ListObjects CLI ONLY FUNCTION!
func (d *DB) ListObjects(objectType string) string {
	switch objectType {
	case "users":
		users := []models.User{}
		res := d.db.Find(&users)
		if res.Error != nil {
			return cliGenErrMsg
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
			return cliGenErrMsg
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
