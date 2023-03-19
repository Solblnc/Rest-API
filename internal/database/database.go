package database

import (
	"fmt"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func NewDataBase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbDatabaseName := os.Getenv("DB_DATABASE_NAME")
	dbPort := os.Getenv("DB_PORT")

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbDatabaseName
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUserName, dbDatabaseName, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	db.DB().Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
