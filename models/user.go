package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Do not query this directly in a handler!
type User struct {
	gorm.Model
	ID                   string `gorm:"primaryKey"`
	UserName             string `gorm:"uniqueIndex"`
	FullName             string
	Password             string
	ImageURL             string
	RegistrationCode     string      `gorm:"-"` // Only used for registration
	SharedWishlists      []*Wishlist `gorm:"many2many:wishlist_share;"`
	Groups               []*Group    `gorm:"many2many:group_user;"`
	ResetToken           string
	ResetTokenExpiration time.Time
}

type SearchUser struct {
	ID       string
	UserName string
	FullName string
}

type ProfileUser struct {
	ID       string
	UserName string
	FullName string
	ImageURL string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
