package main

import (
	"Market-Bot/clientGo"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
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
	err = bot.Start()
	log.Fatal(err)

}
func badMessage(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "а? не понимаю....")
	//	tbot.OptReplyKeyboardRemoveSelective()
}

func registrationHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Для регистрации, боту необходим ваш пароль. Длина пароля должна быть от шести символов и больше.\nПравильный ввод:\npass:your_password", tbot.OptReplyKeyboardRemove)
	bot.HandleMessage(".+", sendPasswHandler)

}

func sendPasswHandler(m *tbot.Message) {
	pass := m.Text
	check, msg := clientGo.CheckCorrectPass(pass)

	if check == false {
		msg = msg + "\nПридумайте получше:"
		client.SendMessage(m.Chat.ID, msg)
		//bot.HandleMessage(".+", sendPasswHandler)

	} else {
		fmt.Println("Ну и где переход")
		client.SendMessage(m.Chat.ID, msg)
		customerInterfaceHandler(m)

		return
		//	clientGo.ClientRegistration(m,pass,db)
	}
}

func customerInterfaceHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Да да вы покупатель, а текст этого сообщения не доделан", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
	//bot.HandleMessage(".+", customerInterfaceHandler)
}

func startHandler(m *tbot.Message) {
	bot.HandleMessage("РЕГИСТРАЦИЯ", registrationHandler)
	keyb := makeButtons("reg")
	fmt.Println(keyb.Keyboard[0])
	keyb.OneTimeKeyboard = true
	client.SendMessage(m.Chat.ID, "Привет! Данный бот предназначен для покупки и продажи товара.\nУ каждого пользователя есть аккаунт для покупки и продажи, смена роли осуществляется через кнопку в меню кнопок.\n\nВозможности покупателя:\n\t- Возможность просматривать товары;\n\t- Добавлять товары в корзину;\n\t- Подтверждение покупки в корзине.\n\nВозможности продавца:\n\t- Создание объявлений с товарами;\n\t- Просмотр своих объявлений.", tbot.OptReplyKeyboardMarkup(keyb))
}

func makeButtons(state string) *tbot.ReplyKeyboardMarkup {
	button1 := tbot.KeyboardButton{
		Text: "РЕГИСТРАЦИЯ",
	}
	button2 := tbot.KeyboardButton{
		Text: "ВХОД",
	}
	button3 := tbot.KeyboardButton{
		Text: "Категории товаров",
	}
	button4 := tbot.KeyboardButton{
		Text: "Выход",
	}
	button5 := tbot.KeyboardButton{
		Text: "Корзина",
	}
	switch state {
	case "reg":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button1, button2},
			},
		}
	case "customer_interface":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button3, button4, button5},
			},
		}
	default:
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{},
			},
		}
	}
}
