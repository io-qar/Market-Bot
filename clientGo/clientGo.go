package clientGo

import (
	"database/sql"
	"fmt"
	"github.com/yanzay/tbot/v2"
)

//данный файл только для вноса данних юзера после регистрации
// вообще пока что это все примерно и временно

func ClientRegistration(m *tbot.Message, password string, db *sql.DB) {
	_, err := db.Exec("INSERT INTO user_table(login,password) values ($1,$2)", m.From.ID, m.Text)
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

func CheckCorrectPass(str string) (bool, string) {
	dubl := 0 // еннн5667
	thpair := str[0]
	for i := 1; i < len(str); i++ {
		thpair = str[i-1]
		if str[i] == thpair {

			dubl++
			thpair = str[i]
			fmt.Println(str[i])
		}
	}
	if dubl >= 2 {
		return false, "У вас тут символы подряд повторяются..."
	}
	if str == "" {
		return false, "Пароль пустой!"
	} else if len(str) < 6 {
		return false, "Пароль короткий("
	}

	return true, "Хороший пароль, лайк!\nЗаношу в бд..."
}
