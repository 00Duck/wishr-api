package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
	log.Print("Connected to database")
}

func (d *DB) AutoMigrate(model ...interface{}) error {
	err := d.db.AutoMigrate(model...)
	return err
}
