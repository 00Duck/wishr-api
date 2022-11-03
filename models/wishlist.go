package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	Name          string
	Items         []WishlistItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Wishlist;"`
	ItemCount     int
	Owner         string
	OwnerFullName string
	CreatedAt     time.Time `gorm:"<-:create"` // allow read and create
	SharedWith    []*User   `gorm:"many2many:wishlist_share;"`
	IsOwner       bool      `gorm:"-"` //handled in app layer
}

type WishlistItem struct {
	ID                 uint `gorm:"primaryKey;autoIncrement;"`
	Wishlist           string
	Name               string
	URL                string `gorm:"size:4000"`
	Notes              string `gorm:"size:4000"`
	Price              string
	PersonalItem       bool
	Quantity           int
	ReservedBy         string
	ReservedByFullName string
}

func (u *Wishlist) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
