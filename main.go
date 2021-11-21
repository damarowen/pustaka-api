package main

import (
	//"encoding/json"

	//	"log"
	//	"log"
	"log"
	"pustaka-api/book"

	// "pustaka-api/DB"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {

	router := gin.Default()
	// err := DB.NewConnDb()


	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{  Logger: logger.Default.LogMode(logger.Info)	})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&book.Book{})

	log.Println("CONECTED TO DB")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.GetAllBookHandler)
	v1.GET("/books/:id", bookHandler.GetByIdHandler)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	//v1.GET("/query", bookHandler.QueryHandler)

	router.Run("127.0.0.1:9090")

}
