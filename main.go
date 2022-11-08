package main

import (
	"Market-Bot/clientGo"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
	"strings"
)

var (
	bot    *tbot.Server
	client *tbot.Client
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// 1014223178:AAFdXeKePaDixf9pK42lK3Co6W9vJQCHHnE
// 5612522930:AAH3NoXrFB0_c0dpHUINJ3yhCkvjWPJ_3Gs ///////

//var db *sql.DB
//
//func ConnectToDB() {
//	var err error
//	connStr := "user=postgres dbname=tg_bot password=1111 host=localhost sslmode=disable"
//	db, err = sql.Open("postgres", connStr)
//	CheckError(err)
//	CheckError(err)
//	defer db.Close()
//	err = db.Ping()
//	CheckError(err)
//	fmt.Printf("\nSuccessfully connected to database!\n")
//}

func main() {

	//ConnectToDB()
	err := godotenv.Load(".env")
	CheckError(err)

	bot = tbot.New(os.Getenv("TOKEN"))
	client = bot.Client()

	bot.HandleMessage("/start", startHandler)
	bot.HandleMessage("РЕГИСТРАЦИЯ", registrationHandler)

	err = bot.Start()
	log.Fatal(err)

}
func badMessage(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "а? не понимаю....", tbot.OptReplyKeyboardMarkup(makeButtonsReg()))
}

func registrationHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Для регистрации, боту необходим ваш пароль. Длина пароля должна быть от шести символов и больше.\nПример правильного ввода\npass:your_password", tbot.OptReplyKeyboardRemove)
	bot.HandleMessage("pass:.+", sendPasswHandler)
	bot.HandleMessage(".+", badMessage)

}

func sendPasswHandler(m *tbot.Message) {
	//fmt.Println(m.Text)
	pass := strings.TrimPrefix(m.Text, "pass:")
	pass = strings.TrimSpace(pass)
	check, msg := clientGo.CheckCorrectPass(pass)

	if check == false {
		msg = msg + "\nПридумайте получше:"
		client.SendMessage(m.Chat.ID, msg)
		bot.HandleMessage(".+", sendPasswHandler)
	} else {
		client.SendMessage(m.Chat.ID, msg)
		//	clientGo.ClientRegistration(m,pass,db)
	}

}

func startHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Привет! Данный бот предназначен для покупки и продажи товара.\nУ каждого пользователя есть аккаунт для покупки и продажи, смена роли осуществляется через кнопку в меню кнопок.\n\nВозможности покупателя:\n\t- Возможность просматривать товары;\n\t- Добавлять товары в корзину;\n\t- Подтверждение покупки в корзине.\n\nВозможности продавца:\n\t- Создание объявлений с товарами;\n\t- Просмотр своих объявлений.", tbot.OptReplyKeyboardMarkup(makeButtonsReg()))
}

func makeButtonsReg() *tbot.ReplyKeyboardMarkup {
	button1 := tbot.KeyboardButton{
		Text: "РЕГИСТРАЦИЯ",
	}

	return &tbot.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tbot.KeyboardButton{
			[]tbot.KeyboardButton{button1},
		},
	}
}
