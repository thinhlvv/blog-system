package article

// Article represents an article.
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
