package main

import (
	"github.com/yanzay/tbot/v2"
	"log"
)

const token = "5612522930:AAH3NoXrFB0_c0dpHUINJ3yhCkvjWPJ_3Gs"

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
	bot = tbot.New(token)
	client = bot.Client()

	bot.HandleMessage("/start", startHandler)

	err := bot.Start()
	log.Fatal(err)
}

func startHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Привет! Данный бот предназначен для покупки и продажи товара.\nУ каждого пользователя есть аккаунт для покупки и продажи, смена роли осуществляется через кнопку в меню кнопок.\n\nВозможности покупателя:\nВозможность просматривать товары.\nДобавлять товары в корзину\nПодтверждение покупки в корзине\n\nВозможности продавца:\nСоздание объявлений с товарами\nПросмотр своих объявлений\n")

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
