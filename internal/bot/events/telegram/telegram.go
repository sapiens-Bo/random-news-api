package telegram

import "github.com/sapiens-Bo/random-news-api/internal/bot/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

func New(client *telegram.Client)
