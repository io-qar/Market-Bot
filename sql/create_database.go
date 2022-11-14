package sql

import (
	"database/sql"
	"fmt"
)

func ConnectToDB() {
	connStr := "user=postgres dbname=tg_bot password=1111 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	defer db.Close()
	err = db.Ping()
	CheckError(err)
	fmt.Printf("\nSuccessfully connected to database!\n")
}
