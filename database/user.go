package database

import (
	"errors"
	"strconv"

	"github.com/00Duck/wishr-api/models"
)

func (d *DB) UserCreate(user *models.User) (string, error) {
	res := d.db.Create(&user)
	return user.ID, res.Error
}

func (d *DB) UserUpdate(user *models.User) error {
	d.db.First(&models.User{}, "id = ?", user.ID)
	res := d.db.Save(&user)
	return res.Error
}

func (d *DB) UserRetrieveOne(id string) (*models.User, error) {
	user := &models.User{}
	res := d.db.First(user, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New("No record found")
	}
	return user, res.Error
}

func (d *DB) UserRetrieveAll() ([]models.User, error) {
	users := []models.User{}
	res := d.db.Find(&users)
	return users, res.Error
}

func (d *DB) UserDelete(id string) (string, error) {
	res := d.db.Delete(&models.User{ID: id})
	if res.RowsAffected == 0 {
		return "", errors.New("No record found to delete")
	}
	return strconv.FormatInt(res.RowsAffected, 10) + " rows deleted", res.Error
}

func (d *DB) GetUsersForWishlist(wishlistID string) ([]models.User, error) {
	users := []models.User{}
	err := d.db.Model(&models.Wishlist{}).Where(&models.Wishlist{ID: wishlistID}).Association("SharedWishlists").Find(&users)
	return users, err
}
