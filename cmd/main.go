package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sapiens-Bo/random-news-api/config"
	"github.com/sapiens-Bo/random-news-api/internal/api/everything"
	"github.com/sapiens-Bo/random-news-api/internal/bot/telegram"
)

func main() {
	cfg := config.MustLoad()

	payload := []byte(`{}`)

	client := &http.Client{}

	req, err := http.NewRequest(everything.Method, everything.URL, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("X-Api-Key", cfg.API)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ошибка чтения тела ответа: %v", err)
	}

	response, err := everything.NewResponse(body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Status: %s\nTotal:%v\nFirst article: %s\n", response.Status, response.Total, response.Articles[0].FormatShow())
}

func StartBot() {
	cfg := config.MustLoad()
	tgClient := telegram.New(cfg.TgBotHost, cfg.TgAPI)
}
