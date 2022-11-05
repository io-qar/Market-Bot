package main

import (
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
	"github.com/joho/godotenv"
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

func main() {
	err := godotenv.Load(".env")
	CheckError(err)

	bot = tbot.New(os.Getenv("TOKEN"))
	client = bot.Client()

	bot.HandleMessage("/start", startHandler)

	err = bot.Start()
	log.Fatal(err)
}

func startHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Привет! Данный бот предназначен для покупки и продажи товара.\nУ каждого пользователя есть аккаунт для покупки и продажи, смена роли осуществляется через кнопку в меню кнопок.\n\nВозможности покупателя:\n\t- Возможность просматривать товары;\n\t- Добавлять товары в корзину;\n\t- Подтверждение покупки в корзине.\n\nВозможности продавца:\n\t- Создание объявлений с товарами;\n\t- Просмотр своих объявлений.")

	//sendUserInfoToBD(m)
}

//func makeButtons(ups, downs int) *tbot.InlineKeyboardMarkup {
//	button1 := tbot.InlineKeyboardButton{
//		Text:         fmt.Sprintf("РЕГИСТРАЦИЯ %d", ups),
//		CallbackData: "up",
//	}
//	button2 := tbot.InlineKeyboardButton{
//		Text:         fmt.Sprintf("👎 %d", downs),
//		CallbackData: "down",
//	}
//	return &tbot.InlineKeyboardMarkup{
//		InlineKeyboard: [][]tbot.InlineKeyboardButton{
//			[]tbot.InlineKeyboardButton{button1, button2},
//		},
//	}
//}
