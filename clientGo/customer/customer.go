package customer

import (
	"github.com/yanzay/tbot/v2"

	"Market-Bot/models"
	"Market-Bot/sql"

	"context"
	"fmt"
	"strconv"
)

//{(strconv.Itoa(c.Number)+". "+c.Name+"("+strconv.Itoa(c.Ads_score)+")"),false}
func GetCategory() []models.Chosen_Category {
	var c models.Category
	categorySl := make([]models.Chosen_Category, 0)
	rows, err := sql.Db.Query(context.Background(), `select * from category`)
	for rows.Next() {
		err = rows.Scan(&c.Number, &c.Name, &c.Ads_score)
		categorySl = append(categorySl, models.Chosen_Category {
			Category_info: (strconv.Itoa(c.Number) + ". " + c.Name + "(" + strconv.Itoa(c.Ads_score) + ")"),
			Chosen:        c.Name,
		})
	}
	if err != nil {
		panic(err)
	}
	return categorySl
}

func ClientShowCategory(m *tbot.Message, client *tbot.Client, bot *tbot.Server) {
	categorySl := GetCategory()
	fmt.Println(categorySl)
	bot.HandleCallback(callbackCategoryHandler)
	_, err := client.SendMessage(m.Chat.ID, "Категории товаров:", tbot.OptInlineKeyboardMarkup(makeCategoryButtons(categorySl)))
	if err != nil {
		panic(err)
	}
	//ChosenID := fmt.Sprintf("%s:%d", m.Chat.ID, msg.MessageID)

}

func callbackCategoryHandler(cq *tbot.CallbackQuery) {
	switch cq.Data {
	case "Одежда и обувь":
		//client.SendMessage(m.Chat.ID, "ОК..", tbot.OptInlineKeyboardMarkup(deleteCategoryButtons()))
		//categoryProductShow(cq.Data, m)

		//client.EditMessageReplyMarkup(cq.Message.Chat.ID, cq.Message.MessageID, tbot.OptInlineKeyboardMarkup(deleteCategoryButtons()))
	case "Аксессуары к одежде":
	}
}

func makeCategoryButtons(variable []models.Chosen_Category) *tbot.InlineKeyboardMarkup {
	buttons := make([]tbot.InlineKeyboardButton, len(variable))
	for i, v := range variable {
		buttons[i].Text = v.Category_info
		buttons[i].CallbackData = v.Chosen
	}
	return &tbot.InlineKeyboardMarkup {
		InlineKeyboard: [][]tbot.InlineKeyboardButton {
			[]tbot.InlineKeyboardButton{buttons[0], buttons[1]},
			[]tbot.InlineKeyboardButton{buttons[2], buttons[3]},
			[]tbot.InlineKeyboardButton{buttons[4], buttons[5]},
			[]tbot.InlineKeyboardButton{buttons[6], buttons[7]},
			[]tbot.InlineKeyboardButton{buttons[8], buttons[9]},
		},
	}
}

func deleteCategoryButtons() *tbot.InlineKeyboardMarkup {
	button1 := tbot.InlineKeyboardButton{
		Text:         fmt.Sprintf(""),
		CallbackData: "",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{button1},
		},
	}
}
