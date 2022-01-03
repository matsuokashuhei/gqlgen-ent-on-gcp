package database

import (
	"fmt"

	"github.com/matsuokashuhei/landin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var err error
	// https://gorm.io/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config("MYSQL_USER"), config.Config("MYSQL_PASSWORD"),config.Config("MYSQL_HOST"), config.Config("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	return db, nil
}
