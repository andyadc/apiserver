package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logger "github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func openDB(username, password, url, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		url,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		logger.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setUp(db)

	return db
}

func setUp(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.url"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}
