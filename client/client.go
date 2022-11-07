package client

import (
	"database/sql"
	"fmt"
	"github.com/yanzay/tbot/v2"
)

//данный файл только для вноса данних юзера после регистрации
// вообще пока что это все примерно и временно

func ClientRegistration(m *tbot.Message, password string, db *sql.DB) {
	_, err := db.Exec("INSERT INTO user_table(login,password) values ($1,$2)", m.From.ID, password)
	if err != nil {
		panic(err)
	}
}

func ClientLogin(m *tbot.Message, password string, db *sql.DB) bool {
	pass_from_db := ""
	if err := db.QueryRow("SELECT (password) from user_table where login = $1",
		m.From.ID).Scan(&pass_from_db); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Нет такого логина")
		}
		if pass_from_db == "" {
			fmt.Println("Пароль пустой")
			return false
		}
	}
	if password == pass_from_db {
		fmt.Println("Зашел удачно")
		return true
	} else {
		fmt.Println("Пароли не совпадают")
		return false
	}
}
