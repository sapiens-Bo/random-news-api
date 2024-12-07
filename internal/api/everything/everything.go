package everything

var (
	URL    = "https://newsapi.org/v2/everything?q=bitcoin"
	Method = "GET"
)

type Response struct {
	Status   string `json:"status"`
	Total    string `json:"totalResult"`
	Articles []Articles
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
