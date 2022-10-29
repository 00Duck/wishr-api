package database

import (
	"errors"
	"strconv"

	"github.com/00Duck/wishr-api/models"
)

func (d *DB) WishlistCreate(wishlist *models.Wishlist) (string, error) {
	res := d.db.Create(&wishlist)
	return wishlist.ID, res.Error
}

func (d *DB) WishlistUpdate(wishlist *models.Wishlist) error {
	//Find wishlist and error if it doesn't exist
	res := d.db.First(&models.Wishlist{}, "id = ?", wishlist.ID)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("No record found to update")
	}
	//delete all associated items
	res = d.db.Where("Wishlist = ?", wishlist.ID).Delete(&models.WishlistItem{})
	if res.Error != nil {
		return res.Error
	}
	//new items on wishlist is the new list
	res = d.db.Save(&wishlist)
	return res.Error
}

func (d *DB) WishlistRetrieveOne(id string) (*models.Wishlist, error) {
	wishlist := &models.Wishlist{}
	res := d.db.Model(&models.Wishlist{}).Preload("Items").Find(wishlist, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New("No record found")
	}
	return wishlist, res.Error
}

func (d *DB) WishlistRetrieveAll() ([]models.Wishlist, error) {
	wishlists := []models.Wishlist{}
	res := d.db.Find(&wishlists)
	return wishlists, res.Error
}

func (d *DB) WishlistDelete(id string) (string, error) {
	res := d.db.Select("WishlistItem").Delete(&models.Wishlist{ID: id})
	if res.RowsAffected == 0 {
		return "", errors.New("No record found to delete")
	}
	return strconv.FormatInt(res.RowsAffected, 10) + " rows deleted", res.Error
}
