package config

import (
	"fmt"
	"log"
	"mulfinance/pkg/limit/repository/entities"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RUN Automigrate.........")

	dbConn.AutoMigrate(
		&entities.Limit{},
		&entities.Debitur{},
		&entities.DebiturLimit{},
		&entities.TransactionLimit{},
		&entities.Transaction{},
	)

	fmt.Println("Automigrate Successfully")

	return dbConn
}
