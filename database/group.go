package database

import (
	"errors"
	"strconv"

	"github.com/00Duck/wishr-api/models"
)

func (d *DB) GroupCreate(group *models.Group) (string, error) {
	res := d.db.Create(&group)
	return group.ID, res.Error
}

func (d *DB) GroupUpdate(group *models.Group) error {
	d.db.First(&models.Group{}, "id = ?", group.ID)
	res := d.db.Save(&group)
	return res.Error
}

// MAY NOT NEED
// func (d *DB) GroupRetrieveOne(id string) (*models.Group, error) {
// 	group := &models.Group{}
// 	res := d.db.First(group, "id = ?", id)
// 	if res.RowsAffected == 0 {
// 		return nil, errors.New("No record found")
// 	}
// 	return group, res.Error
// }

func (d *DB) GroupDelete(id string) (string, error) {
	res := d.db.Delete(&models.Group{ID: id})
	if res.RowsAffected == 0 {
		return "", errors.New("No record found to delete")
	}
	return strconv.FormatInt(res.RowsAffected, 10) + " rows deleted", res.Error
}

func (d *DB) GetGroupsForUser(userID string) ([]models.Group, error) {
	groups := []models.Group{}
	if userID == "" {
		return groups, errors.New("You must provide a user ID to get a group list")
	}
	err := d.db.Model(&models.User{ID: userID}).Association("Users").Find(&groups)
	return groups, err
}

// // Returns all groups as SearchGroup slice, minus those that are already shared and the current session group
// func (d *DB) GetSelectableGroupsToShareList(session *models.Session, wishlistID string) ([]models.SearchGroup, error) {
// 	selectableGroups := []models.SearchGroup{}
// 	sharedGroups, err := d.GetGroupsForWishlist(wishlistID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sharedGroupIDs := getIDsFromGroups(sharedGroups)
// 	sharedGroupIDs = append(sharedGroupIDs, session.GroupID)
// 	res := d.db.Model(&models.Group{}).Where("id NOT IN ?", sharedGroupIDs).Find(&selectableGroups)
// 	return selectableGroups, res.Error
// }

// func getIDsFromGroups(groups []models.SearchGroup) []string {
// 	ret := []string{}
// 	for _, k := range groups {
// 		ret = append(ret, k.ID)
// 	}
// 	return ret
// }

// func (d *DB) SetGroupsForWishlist(wishlistID string, groups []models.Group) error {
// 	//turn off hooks as to not generate a new ID (was causing an insert in group table)
// 	err := d.db.Session(&gorm.Session{
// 		SkipHooks: true,
// 	}).Model(&models.Wishlist{ID: wishlistID}).Association("SharedWith").Replace(&groups)
// 	return err
// }
