package wishr

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(env *Env) {
	conn_str := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD")
	conn_db := os.Getenv("DB_DATABASE")
	conn_port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		conn_port = "3306"
	}
	var err error
	env.db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn_str + "@tcp(127.0.0.1:" + conn_port + ")/" + conn_db + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                                  // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                 // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                 // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                 // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
}
