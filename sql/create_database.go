package sql

import (
	"Market-Bot/models"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"

	"context"
	"os"
)

func ConnectToDB() {
	err := godotenv.Load(".env")
	models.CheckError(err)

	var dbURL string = os.Getenv("URL")

	db, err := pgx.Connect(context.Background(), dbURL)
	models.CheckError(err)
	defer db.Close(context.Background())
	err = db.Ping(context.Background())
	models.CheckError(err)
	println("Successfully connected to database!")
}
