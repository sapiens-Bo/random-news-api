package everything

import (
	"encoding/json"
	"fmt"
)

var (
	URL    = "https://newsapi.org/v2/everything?q=bitcoin&language=ru"
	Method = "GET"
)

type Response struct {
	Status   string     `json:"status"`
	Total    int        `json:"totalResults"`
	Articles []Articles `json:"articles"`
}

type Articles struct {
	Source struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLTOImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

func NewResponse(body []byte) (*Response, error) {
	const op = "api.everything.NewResponse"
	var Response Response
	err := json.Unmarshal(body, &Response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Response, nil
}

func (a *Articles) FormatShow() string {
	return fmt.Sprintf(`
    id: %s
    name: %s
    Author: %s
    Title: %s
    Description: %s
    url: %s
    url to image: %s
    Published at: %s
    Content: %s
    `, a.Source.Id, a.Source.Name,
		a.Author, a.Title, a.Description,
		a.URL, a.URLTOImage, a.PublishedAt, a.Content)
}
