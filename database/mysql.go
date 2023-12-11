package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectMysqlDB(ctx context.Context) {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("MYSQL_USERNAME"),
		viper.GetString("MYSQL_PASSWORD"),
		viper.GetString("MYSQL_HOST"),
		viper.GetString("MYSQL_PORT"),
		viper.GetString("MYSQL_DATABASE"))

	fmt.Println(conn)
	dbConn, err := gorm.Open(mysql.Open(conn))
	if err != nil {
		panic("Mysql failed to open/connect to the database")
	}
	db = dbConn
}

func Close() error {
	fmt.Println("mysql close function called")
	db, err := db.DB()
	if err != nil {
		return err
	}
	err = db.Close()

	return err
}

func GetConnection() *gorm.DB {
	return db
}
