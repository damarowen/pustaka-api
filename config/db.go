package config

import (
	"fmt"
	"log"
	"os"
	"pustaka-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConn struct {
	DbSQL  *gorm.DB

	// other db...
	// DbMasterFDBR *gorm.DB
	// DbSlaveFDBR  *gorm.DB
	// DbMasterForum *gorm.DB
	// DbSlaveForum  *gorm.DB
	// DbGiveaway *gorm.DB
	// DbAdmin    *gorm.DB

}



func ConnectDatabase() (data *DbConn, err error){

	dbSource := &DbConn{}

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	dbSource.DbSQL, err  = gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	dbSource.DbSQL.AutoMigrate(&models.Book{})
	log.Println("CONECTED TO DB")


	return dbSource, err
}


//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}