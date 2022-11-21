package customer

import (
	"Market-Bot/models"
	"Market-Bot/sql"
	"context"
	"fmt"
	"github.com/yanzay/tbot/v2"
	"strconv"
)

//{(strconv.Itoa(c.Number)+". "+c.Name+"("+strconv.Itoa(c.Ads_score)+")"),false}
func GetCategory() []models.Chosen_Category {
	var c models.Category
	categorySl := make([]models.Chosen_Category, 0)
	rows, err := sql.Db.Query(context.Background(), `select * from category`)
	for rows.Next() {
		err = rows.Scan(&c.Number, &c.Name, &c.Ads_score)
		categorySl = append(categorySl, models.Chosen_Category{
			Category_info: (strconv.Itoa(c.Number) + ". " + c.Name + "(" + strconv.Itoa(c.Ads_score) + ")"),
			Chosen:        c.Name,
		})
	}
	models.CheckError(err)
	return categorySl
}

func ClientShowCategory(m *tbot.Message, client *tbot.Client, bot *tbot.Server) {
	categorySl := GetCategory()
	fmt.Println(categorySl)
	bot.HandleCallback(callbackCategoryHandler)
	_, err := client.SendMessage(m.Chat.ID, "Категории товаров:", tbot.OptInlineKeyboardMarkup(makeCategoryButtons(categorySl)))
	models.CheckError(err)
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

func makeCategoryButtons(varible []models.Chosen_Category) *tbot.InlineKeyboardMarkup {
	buttons := make([]tbot.InlineKeyboardButton, len(varible))
	for i, v := range varible {
		buttons[i].Text = v.Category_info
		buttons[i].CallbackData = v.Chosen

	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			buttons,
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
