package main

//go get github.com/georgysavva/scany/v2
import (
	"Market-Bot/clientGo/customer"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
	"io"
	"net/http"

	"Market-Bot/clientGo"
	"Market-Bot/models"
	"Market-Bot/sql"

	"fmt"
	"log"
	"os"
)

var (
	bot    *tbot.Server
	client *tbot.Client
	state  = map[string]string{}
)

type user struct {
	username string
	state    string
}

func main() {
	sql.ConnectToDB()
	sql.CreateDataBase()

	err := godotenv.Load(".env")
	models.CheckError(err)
	bot = tbot.New(os.Getenv("TOKEN"))
	client = bot.Client()

	bot.HandleMessage("/start", startHandler)
	bot.HandleMessage(".+", stateHandler)
	bot.HandleMessage("", func(m *tbot.Message) {
		if m.Document != nil {
			doc, err := client.GetFile(m.Document.FileID)
			if err != nil {
				log.Println(err)
				return
			}
			url := client.FileURL(doc)
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()
			err = os.Chdir("./imgs")
			out, err := os.Create(m.Document.FileName)
			if err != nil {
				log.Println(err)
				return
			}
			defer out.Close()
			io.Copy(out, resp.Body)
			fmt.Println(out, resp)
		}
	})
	err = bot.Start()
	log.Fatal(err)
}

func stateHandler(m *tbot.Message) {
	switch state[m.From.Username] {
	case "START":
		switch m.Text {
		case "РЕГИСТРАЦИЯ":
			registrationHandler(m)
		case "ВХОД":
			check, _ := clientGo.LoginCheck(m)
			if !check {
				loginHandler(m)
			} else {
				client.SendMessage(m.Chat.ID, "Мне тут по секрету сказали, что вас нету в базе...", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
				state[m.From.Username] = "START"
			}
		default:
			client.SendMessage(m.Chat.ID, "а? Не понимаю...")
		}
	case "LOGIN":
		checkPassHandler(m)
	case "REG":
		sendPasswHandler(m)
	case "CHANGE_PASS":
		changePasswHandler(m)
	case "DELETE_ACC":
		switch m.Text {
		case "Yes":
			clientGo.DeleteAcc(m)
			client.SendMessage(m.Chat.ID, "Аккаунт удалён!", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
			client.SendSticker(m.Chat.ID, "CAACAgIAAxkBAAEGcoRjdS_QjsmJSpkAAbp4s8mYIA6XkqUAAkcUAALd3YBKLVp8NfRstaIrBA")
			state[m.From.Username] = "START"
		case "NO":
			client.SendMessage(m.Chat.ID, "Ну.. Ладно..", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
			state[m.From.Username] = "CLIENT_INTERFACE"
		}
	case "CLIENT_CART":
		switch m.Text {
		case "Удалить товары":
			customer.ClientDeleteAllProductsFromCart(m, client, bot)
		case "Назад":
			state[m.From.Username] = "CLIENT_INTERFACE"
			client.SendMessage(m.Chat.ID, "Обратно к интерфейсу...", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
		case "Купить товары":
			customer.ClientBuyAllProduct(m, client, bot)
		}
	case "CLIENT_FAVOR":
		switch m.Text {
		case "Удалить товары":
			customer.ClientDeleteAllProductsFromFavor(m, client, bot)
		case "Назад":
			state[m.From.Username] = "CLIENT_INTERFACE"
			client.SendMessage(m.Chat.ID, "Обратно к интерфейсу...", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
		}
	case "CLIENT_SETTINGS":
		switch m.Text {
		case "Выход":
			state[m.From.Username] = "START"
			client.SendMessage(m.Chat.ID, "Выход из аккаунта", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
			client.SendSticker(m.Chat.ID, "CAACAgIAAxkBAAEGaxFjclj9sC5c8pPkx1YpjaH0l9BHtQACARUAAn3WYUhaT836O4P01isE")
		case "Сменить роль":
			client.SendMessage(m.Chat.ID, "Смена роли. Теперь вы продавец.", tbot.OptReplyKeyboardMarkup(makeButtons("seller_interface")))
			state[m.From.Username] = "SELLER_INTERFACE"
		case "Изменить пароль":
			client.SendMessage(m.Chat.ID, "Введите пароль", tbot.OptReplyKeyboardRemove)
			state[m.From.Username] = "CHANGE_PASS"
		case "Удалить аккаунт":
			client.SendMessage(m.Chat.ID, "Вы действительно хотите удалить аккаунт?", tbot.OptReplyKeyboardMarkup(makeButtons("delete_acc")))
			state[m.From.Username] = "DELETE_ACC"
		case "Назад":
			client.SendMessage(m.Chat.ID, "Обратно к интерфейсу...", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
			state[m.From.Username] = "CLIENT_INTERFACE"
		}
	case "CLIENT_INTERFACE":
		switch m.Text {
		case "Купленные товары":
			customer.ClientShowOrderProduct(m, client, bot)
		case "Категории товаров":
			client.SendMessage(m.Chat.ID, "Ну тут кароч будут категории в виде кнопок, еще в каждой категории указывается количество существующих объявлений")
			customer.ClientShowCategory(m, client, bot)
		case "Корзина":
			client.SendMessage(m.Chat.ID, "Ваши товары:", tbot.OptReplyKeyboardMarkup(makeButtons("customer_shopping_cart")))
			sql.ClientShowCart(m, client, bot)
			state[m.From.Username] = "CLIENT_CART"
		case "Настройки":
			state[m.From.Username] = "CLIENT_SETTINGS"
			client.SendMessage(m.Chat.ID, "Настройки", tbot.OptReplyKeyboardMarkup(makeButtons("customer_settings")))
		case "Избранное":
			client.SendMessage(m.Chat.ID, "Ваши избранные товары:", tbot.OptReplyKeyboardMarkup(makeButtons("customer_favor_cart")))
			sql.ClientShowFavor(m, client, bot)
			state[m.From.Username] = "CLIENT_FAVOR"
		}
	case "SELLER_INTERFACE":
		switch m.Text {
		case "Выход":
			state[m.From.Username] = "START"
			client.SendMessage(m.Chat.ID, "Выход из аккаунта", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
			client.SendSticker(m.Chat.ID, "CAACAgIAAxkBAAEGaxFjclj9sC5c8pPkx1YpjaH0l9BHtQACARUAAn3WYUhaT836O4P01isE")
		case "Сменить роль":
			client.SendMessage(m.Chat.ID, "Смена роли. Теперь вы покупатель.", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
			state[m.From.Username] = "CLIENT_INTERFACE"
		}
	default:
		client.SendMessage(m.Chat.ID, "а? Не понимаю...")
	}
}

func loginHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Введите пароль:", tbot.OptReplyKeyboardRemove)
	state[m.From.Username] = "LOGIN"
}

func registrationHandler(m *tbot.Message) {
	fmt.Println(m.Text)
	check, _ := clientGo.LoginCheck(m)
	if check {
		client.SendMessage(m.Chat.ID, "Для регистрации, боту необходим ваш пароль. Длина пароля должна быть от шести символов и больше.\nВведите пароль", tbot.OptReplyKeyboardRemove)
		state[m.From.Username] = "REG"
	} else {
		client.SendMessage(m.Chat.ID, "С таким логином уже регистрировались", tbot.OptReplyKeyboardMarkup(makeButtons("reg")))
		state[m.From.Username] = "START"
	}
}

func checkPassHandler(m *tbot.Message) {
	pass := m.Text
	check := clientGo.ClientLogin(m, pass)
	if check {
		client.SendMessage(m.Chat.ID, "Пароль верный!")
		state[m.From.Username] = "CLIENT_INTERFACE"
		customer.CallBackDataOn(client, bot)
		customerInterfaceHandler(m)
	} else {
		client.SendMessage(m.Chat.ID, "Неправильный пароль")
	}
}

func sendPasswHandler(m *tbot.Message) {
	pass := m.Text
	check, msg := clientGo.CheckCorrectPass(pass)
	if !check {
		msg += "\nПридумайте получше:"
		client.SendMessage(m.Chat.ID, msg)
	} else {
		println("Ну и где переход")
		client.SendMessage(m.Chat.ID, msg)
		clientGo.ClientRegistration(m)
		customerInterfaceHandler(m)
	}
}

func changePasswHandler(m *tbot.Message) {
	pass := m.Text
	check, msg := clientGo.CheckCorrectPass(pass)
	if !check {
		msg += "\nПридумайте получше:"
		client.SendMessage(m.Chat.ID, msg)
	} else {
		fmt.Println("Ну и где переход")
		clientGo.ChangePassword(m)
		client.SendMessage(m.Chat.ID, "Пароль изменен", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
		state[m.From.Username] = "CLIENT_INTERFACE"
	}
}

func customerInterfaceHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Да да вы покупатель, а текст этого сообщения не доделан", tbot.OptReplyKeyboardMarkup(makeButtons("customer_interface")))
	client.SendSticker(m.Chat.ID, "CAACAgIAAxkBAAEGaxNjcllRH0z0TqnjUA5zl5Otm0tkvwACwhUAAlAdSUhTlP1Qw1XqOCsE")
	state[m.From.Username] = "CLIENT_INTERFACE"
}

func startHandler(m *tbot.Message) {
	state[m.From.Username] = "START"
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
	button8 := tbot.KeyboardButton{
		Text: "Купить товары",
	}
	button9 := tbot.KeyboardButton{
		Text: "Удалить товары",
	}
	button10 := tbot.KeyboardButton{
		Text: "Назад",
	}
	button11 := tbot.KeyboardButton{
		Text: "Настройки",
	}
	button12 := tbot.KeyboardButton{
		Text: "Изменить пароль",
	}
	button13 := tbot.KeyboardButton{
		Text: "Удалить аккаунт",
	}
	button14 := tbot.KeyboardButton{
		Text: "Yes",
	}
	button15 := tbot.KeyboardButton{
		Text: "NO",
	}
	button16 := tbot.KeyboardButton{
		Text: "Купленные товары",
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
				[]tbot.KeyboardButton{button3, button5},
				[]tbot.KeyboardButton{button6, button11},
				[]tbot.KeyboardButton{button16},
			},
		}
	case "customer_favor_cart":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button9, button10},
			},
		}
	case "seller_interface":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button4, button7},
			},
		}
	case "customer_shopping_cart":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button8, button9, button10},
			},
		}
	case "customer_settings":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button10, button7},
				[]tbot.KeyboardButton{button4, button12},
				[]tbot.KeyboardButton{button13},
			},
		}
	case "delete_acc":
		return &tbot.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard: [][]tbot.KeyboardButton{
				[]tbot.KeyboardButton{button14, button15},
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
