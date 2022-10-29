package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	UserName string
	UserID   string
	FullName string
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	s.ID = uid.String()
	return nil
}
