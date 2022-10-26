package database

import (
	"github.com/00Duck/wishr-api/models"
)

func (d *DB) Authenticate(auth *models.AuthenticateModel) string {
	return "Received " + auth.User + " Pass: " + auth.Password
}
