package sql

import (
	"Market-Bot/models"

	"github.com/joho/godotenv"

	"database/sql"
	"fmt"
	"os"
)

func ConnectToDB() {
	err := godotenv.Load("../.env")
	models.CheckError(err)

	var (
		dbUser string = os.Getenv("USER")
		dbName string = os.Getenv("DBNAME")
		dbPassword string = os.Getenv("PASSWORD")
		dbHost string = os.Getenv("HOST")
		dbMode string = os.Getenv("MODE")
	)

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", dbUser, dbName, dbPassword, dbHost, dbMode)
	db, err := sql.Open("postgres", connStr)
	models.CheckError(err)
	defer db.Close()
	err = db.Ping()
	models.CheckError(err)
	fmt.Printf("\nSuccessfully connected to database!\n")
}
