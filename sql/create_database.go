package sql

import (
	"Market-Bot/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"io"
	"os"
)

var Db *pgx.Conn

func ConnectToDB() {
	err := godotenv.Load(".env")
	models.CheckError(err)

	var dbURL string = os.Getenv("URL")

	Db, err = pgx.Connect(context.Background(), dbURL)
	models.CheckError(err)
	//	defer db.Close(context.Background())
	err = Db.Ping(context.Background())
	models.CheckError(err)
	println("Successfully connected to database!")
}

func CreateDataBase() {
	var baseExist bool
	row := Db.QueryRow(
		context.Background(),
		"SELECT EXISTS (SELECT FROM pg_database WHERE datname = 'market_bot')")
	err := row.Scan(&baseExist)
	models.CheckError(err)
	fmt.Println(baseExist)
	if baseExist == true {
		var dbURL string = os.Getenv("URLMARKET")
		Db, err = pgx.Connect(context.Background(), dbURL)
		models.CheckError(err)
		fmt.Println("Successfully connected to MARKET database!")
		return
	}
	sqlString := ""
	file, err := os.Open("./sql/database.sql")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		sqlString += string(data[:n])
	}
	defer file.Close()
	_, err = Db.Exec(context.Background(), "create database market_bot;")
	_, err = Db.Exec(context.Background(), sqlString)
	models.CheckError(err)
	fmt.Println("Successfully created tables")

}
