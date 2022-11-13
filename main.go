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
	state  string
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// 1014223178:AAFdXeKePaDixf9pK42lK3Co6W9vJQCHHnE
// 5612522930:AAH3NoXrFB0_c0dpHUINJ3yhCkvjWPJ_3Gs ///////

func main() {
	//	sql.ConnectToDB()
	state = "START"
	//ConnectToDB()
	err := godotenv.Load(".env")
	CheckError(err)
	bot = tbot.New(os.Getenv("TOKEN"))
	client = bot.Client()
	bot.HandleMessage("/start", startHandler)
	bot.HandleMessage(".+", stateHandler)
	err = bot.Start()
	log.Fatal(err)

}

func badMessage(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "а? не понимаю....")
	//	tbot.OptReplyKeyboardRemoveSelective()
}

func stateHandler(m *tbot.Message) {
	switch state {
	case "START":
		if m.Text == "РЕГИСТРАЦИЯ" {
			registrationHandler(m)
		} else if m.Text == "ВХОД" {
			loginHandler(m)
		}
	case "LOGIN":
		checkPassHandler(m)
	case "REG":
		sendPasswHandler(m)
	case "CLIENT_INTERFACE":
		switch m.Text {
		case "Выход":
			state = "START"
			client.SendMessage(m.Chat.ID, "Выход из аккаунта", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
		case "Сменить роль":
			client.SendMessage(m.Chat.ID, "Смена роли. Теперь вы продавец.", tbot.OptReplyKeyboardMarkup(makeButtons("seller_interface")))
			state = "SELLER_INTERFACE"
		}
	case "SELLER_INTERFACE":
		switch m.Text {
		case "Выход":
			state = "START"
			client.SendMessage(m.Chat.ID, "Выход из аккаунта", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
		case "Сменить роль":
			client.SendMessage(m.Chat.ID, "Смена роли. Теперь вы покупатель.", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
			state = "CLIENT_INTERFACE"
		}
	default:
		client.SendMessage(m.Chat.ID, "а? Не понимаю...")
	}
}

func loginHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Введите пароль:", tbot.OptReplyKeyboardRemove)
	state = "LOGIN"
}

func registrationHandler(m *tbot.Message) {
	fmt.Println(m.Text)
	client.SendMessage(m.Chat.ID, "Для регистрации, боту необходим ваш пароль. Длина пароля должна быть от шести символов и больше.\nВведите пароль", tbot.OptReplyKeyboardRemove)
	state = "REG"
}

func checkPassHandler(m *tbot.Message) {
	pass := m.Text
	pass_from_bd := pass
	// some chec bd log func
	if pass == pass_from_bd {
		client.SendMessage(m.Chat.ID, "Пароль верный!")
		state = "CLIENT_INTERFACE"
		customerInterfaceHandler(m)
	} else {
		client.SendMessage(m.Chat.ID, "Неправильный пароль")

	}
}

func sendPasswHandler(m *tbot.Message) {
	pass := m.Text
	check, msg := clientGo.CheckCorrectPass(pass)
	if check == false {
		msg = msg + "\nПридумайте получше:"
		client.SendMessage(m.Chat.ID, msg)
	} else {
		fmt.Println("Ну и где переход")
		client.SendMessage(m.Chat.ID, msg)
		customerInterfaceHandler(m)
		//clientGo.ClientRegistration(m,pass,db)
	}
}

func customerInterfaceHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Да да вы покупатель, а текст этого сообщения не доделан", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
	state = "CLIENT_INTERFACE"
}

func startHandler(m *tbot.Message) {
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
	button6 := tbot.KeyboardButton{
		Text: "Избранное",
	}
	button7 := tbot.KeyboardButton{
		Text: "Сменить роль",
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
				[]tbot.KeyboardButton{button3, button4, button5, button6, button7},
			},
		}
	case "seller_interface":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button4, button7},
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
