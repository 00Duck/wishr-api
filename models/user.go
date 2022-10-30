package models

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	UserName        string `gorm:"uniqueIndex"`
	FullName        string
	Password        string
	SharedWishlists []*Wishlist `gorm:"many2many:wishlist_share;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
