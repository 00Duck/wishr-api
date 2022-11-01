package database

import (
	"errors"
	"strconv"

	"github.com/00Duck/wishr-api/models"
	"gorm.io/gorm"
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

func (d *DB) GetUsersForWishlist(wishlistID string) ([]models.SearchUser, error) {
	users := []models.SearchUser{}
	if wishlistID == "" {
		return users, errors.New("You must provide a wishlist ID to get a shared user list")
	}
	err := d.db.Model(&models.Wishlist{ID: wishlistID}).Association("SharedWith").Find(&users)
	return users, err
}

// Returns all users as SearchUser slice, minus those that are already shared and the current session user
func (d *DB) GetSelectableUsersToShareList(session *models.Session, wishlistID string) ([]models.SearchUser, error) {
	selectableUsers := []models.SearchUser{}
	sharedUsers, err := d.GetUsersForWishlist(wishlistID)
	if err != nil {
		return nil, err
	}
	sharedUserIDs := getIDsFromUsers(sharedUsers)
	sharedUserIDs = append(sharedUserIDs, session.UserID)
	res := d.db.Model(&models.User{}).Where("id NOT IN ?", sharedUserIDs).Find(&selectableUsers)
	return selectableUsers, res.Error
}

func getIDsFromUsers(users []models.SearchUser) []string {
	ret := []string{}
	for _, k := range users {
		ret = append(ret, k.ID)
	}
	return ret
}

func (d *DB) SetUsersForWishlist(wishlistID string, users []models.User) error {
	//turn off hooks as to not generate a new ID (was causing an insert in user table)
	err := d.db.Session(&gorm.Session{
		SkipHooks: true,
	}).Model(&models.Wishlist{ID: wishlistID}).Association("SharedWith").Replace(&users)
	return err
}
