package database

import (
	"errors"
	"log"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"gorm.io/gorm"
)

func (d *DB) Authenticate(login *models.LoginModel) (*models.Session, error) {
	ERR_BAD_PW := errors.New("User Name or Password is incorrect")
	ERR_STH_BAD := errors.New("Something went wrong attempting to log you in. Please contact your administrator for help.")
	if login.UserName == "" || login.Password == "" {
		return nil, ERR_BAD_PW
	}
	user := &models.User{}
	res := d.db.Where(&models.User{UserName: login.UserName}).First(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, ERR_BAD_PW
	}
	ok := auth.CheckPasswordHash(login.Password, user.Password)
	if !ok {
		return nil, ERR_BAD_PW
	}
	session := &models.Session{}
	session.UserID = user.ID
	session.UserName = user.UserName
	session.FullName = user.FullName
	res = d.db.Create(&session)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil, ERR_STH_BAD
	}
	if res.RowsAffected != 1 {
		log.Println(res.Error.Error())
		return nil, ERR_STH_BAD
	}
	return session, nil
}

func (d *DB) Deauthenticate(sessionID string) error {
	session := &models.Session{
		ID: sessionID,
	}
	//Unscoped deletes permanently (no soft delete)
	res := d.db.Unscoped().Delete(&session)
	if res.RowsAffected != 1 {
		return errors.New("Could not deauthenticate user: no session found with ID " + sessionID)
	}
	return res.Error
}

func (d *DB) Register(user *models.User) error {
	res := d.db.Where(&models.User{UserName: user.UserName}).First(&models.User{})
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return errors.New("There was a problem attempting to register: " + res.Error.Error())
	}
	if res.RowsAffected == 1 {
		return errors.New("The name you have chosen is already in use")
	}
	pw, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pw
	res = d.db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errors.New("There was a problem creating your account. Please try again later.")
	}
	return nil
}

func (d *DB) CheckSession(sessionIDValue string) (*models.Session, error) {
	session := &models.Session{}
	res := d.db.Where(&models.Session{ID: sessionIDValue}).First(&session)
	if res.Error != nil {
		return nil, res.Error
	}
	return session, nil
}
