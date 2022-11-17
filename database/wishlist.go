package database

import (
	"errors"
	"strconv"

	"github.com/00Duck/wishr-api/models"
	"gorm.io/gorm"
)

func (d *DB) WishlistUpsert(wishlist *models.Wishlist) (string, error) {
	res := d.db.Where("Wishlist = ?", wishlist.ID).Delete(&models.WishlistItem{})
	if res.Error != nil {
		return "", res.Error
	}
	//new items on wishlist is the new list
	res = d.db.Save(&wishlist)
	return wishlist.ID, res.Error
}

func (d *DB) WishlistRetrieveOne(session *models.Session, id string) (*models.Wishlist, error) {
	wishlist := &models.Wishlist{}
	res := d.db.Debug().Model(&models.Wishlist{}).Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("wishlist_items.order ASC")
	}).Find(wishlist, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New("No record found")
	}
	wishlist.IsOwner = session.UserID == wishlist.Owner
	wishlist.OwnerFullName = d.getUserFullName(wishlist.Owner)
	if wishlist.IsOwner {
		return wishlist, res.Error
	}
	//If the list is no longer shared with this user, don't return it
	sharedUsers := []models.User{}
	err := d.db.Model(&models.Wishlist{ID: id}).Association("SharedWith").Find(&sharedUsers, &models.User{ID: session.UserID})
	if err != nil {
		return nil, err
	}
	if len(sharedUsers) == 0 {
		return nil, errors.New("No record found")
	}
	return wishlist, nil
}

func (d *DB) WishlistRetrieveAll(session *models.Session) ([]models.Wishlist, error) {
	wishlists := []models.Wishlist{}
	res := d.db.Where(&models.Wishlist{Owner: session.UserID}).Find(&wishlists)

	for i := 0; i < len(wishlists); i++ {
		wishlists[i].IsOwner = true
		wishlists[i].OwnerFullName = session.FullName
	}
	return wishlists, res.Error
}

func (d *DB) WishlistDelete(id string) (string, error) {
	res := d.db.Unscoped().Select("Items").Delete(&models.Wishlist{ID: id})
	if res.RowsAffected == 0 {
		return "", errors.New("No record found to delete")
	}
	return strconv.FormatInt(res.RowsAffected, 10) + " rows deleted", res.Error
}

func (d *DB) GetSharedWishlists(session *models.Session) ([]models.Wishlist, error) {
	wishlists := []models.Wishlist{}
	err := d.db.Model(&models.User{ID: session.UserID}).Association("SharedWishlists").Find(&wishlists)

	// Disable sharing and editing since these are shared lists
	for i := 0; i < len(wishlists); i++ {
		wishlists[i].IsOwner = false
		wishlists[i].OwnerFullName = d.getUserFullName(wishlists[i].Owner)
	}
	return wishlists, err
}

func (d *DB) ReserveWishlistItem(session *models.Session, wlItem *models.WishlistItem) error {
	testWlItem := &models.WishlistItem{}
	err := d.db.Find(&testWlItem, "id = ?", wlItem.ID).Error
	if err != nil {
		return err
	}
	if testWlItem.ReservedBy != "" {
		return errors.New("Item has already been reserved by someone else")
	}
	res := d.db.Find(&wlItem).Updates(&models.WishlistItem{ReservedBy: session.UserID, ReservedByFullName: session.FullName})
	if res.RowsAffected == 0 {
		return errors.New("No record found to update")
	}
	return res.Error
}

func (d *DB) UnreserveWishlistItem(session *models.Session, wlItem *models.WishlistItem) error {
	res := d.db.Model(&wlItem).Select("reserved_by", "reserved_by_full_name").Updates(&models.WishlistItem{ReservedBy: "", ReservedByFullName: ""})
	if res.RowsAffected == 0 {
		return errors.New("No record found to update")
	}
	return res.Error
}

func (d *DB) getUserFullName(userID string) string {
	if userID == "" {
		return ""
	}
	user := &models.User{ID: userID}
	err := d.db.First(&user).Error
	if err != nil {
		return ""
	}
	return user.FullName
}
