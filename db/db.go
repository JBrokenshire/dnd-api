package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Init() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		os.Getenv("DB_PASSWORD"),
		host,
		port,
		os.Getenv("DB_NAME"),
	)

	log.Printf("Connecting to %v on port %v with username %v", host, port, user)

	// Use mysql.Open() instead of passing a string
	db, err := gorm.Open(dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get generic database object sql.DB to set connection properties
	mysqlDB := db.DB()

	// Set the maximum lifetime for a connection
	mysqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func Paginate(c echo.Context) (int, int, func(db *gorm.DB) *gorm.DB) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 25
	}

	return page, pageSize, func(db *gorm.DB) *gorm.DB {
		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
