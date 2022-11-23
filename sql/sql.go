package sql

import (
	"github.com/yanzay/tbot/v2"

	"context"
)

func GetContext() (context.Context, context.CancelFunc) {
	if Db.Config().ConnectTimeout > 0 {
		return context.WithTimeout(context.Background(), Db.Config().ConnectTimeout)
	}
	return context.Background(), nil
}

func categoryProductShow(category string, m *tbot.Message) {
	println("asffqfqe")
}
