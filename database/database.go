package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/00Duck/wishr-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	db *gorm.DB
}

func (d *DB) Connect() {
	conn_str := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD")
	conn_db := os.Getenv("DB_DATABASE")
	conn_port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		conn_port = "3306"
	}
	conn_host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		conn_host = "127.0.0.1"
	}
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conn_str, conn_host, conn_port, conn_db)
	var err error
	d.db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // turn this off - we'll do our own logging
	})
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
	log.Println("Connected to database")
	err = d.AutoMigrate(&models.User{}, &models.Wishlist{}, &models.WishlistItem{}, &models.Session{}, &models.Group{})
	if err != nil {
		log.Fatal("Failed to AutoMigrate database.")
	}

	go d.scheduleDeleteOldSessions()
}

func (d *DB) AutoMigrate(model ...interface{}) error {
	err := d.db.AutoMigrate(model...)
	return err
}

func (d *DB) scheduleDeleteOldSessions() {
	for {
		time.Sleep(time.Minute * 15)
		daysAgo := time.Now().Add(-(time.Hour * 24 * 30)) // days
		err := d.db.Unscoped().Model(&models.Session{}).Where("created_at < ?", daysAgo).Delete(&models.Session{}).Error
		if err != nil {
			log.Println("Error deleting stale sessions: " + err.Error())
		}
	}

}
