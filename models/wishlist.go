package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Items       []WishlistItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Wishlist;"`
	ItemCount   int
	Owner       string
	CreatedAt   time.Time `gorm:"<-:create"` // allow read and create
}

type WishlistItem struct {
	ID       uint `gorm:"primaryKey;autoIncrement;"`
	Wishlist string
	Name     string
	URL      string
	Notes    string
	Price    string
	Quantity int
}

func (u *Wishlist) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = strings.Replace(uid.String(), "-", "", -1)
	return nil
}