package models

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex"`
	Description string
	GroupOwner  string
	Users       []*User `gorm:"many2many:group_user;"`
}

func (g *Group) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	g.ID = strings.Replace(uid.String(), "-", "", -1)
	return nil
}
