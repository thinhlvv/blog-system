package article

import "github.com/thinhlvv/blog-system/model"

// Article represents an article.
type Article struct {
	ID        int
	Title     string
	Content   string
	Author    string
	CreatedAt model.Date
	UpdatedAt model.Date
	DeletedAt model.Date
}
