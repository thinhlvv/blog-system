package article

import (
	"database/sql"
	"fmt"
)

// Repository is interface to interact to outside world.
type (
	// Repository is interface to interact with outside world.
	Repository interface {
		CreateArticle(art Article) (int, error)
		GetArticleByID(id string) (*Article, error)
		GetAllArticles() ([]Article, error)
	}

	repoImpl struct {
		db *sql.DB
	}
)

// NewRepo returns repository implementation.
func NewRepo(db *sql.DB) Repository {
	return &repoImpl{db: db}
}

func (repo *repoImpl) CreateArticle(art Article) (int, error) {
	stmt := fmt.Sprintf(`
		INSERT INTO article (title, content, author)
		VALUES (?, ?, ?)
	`)

	res, err := repo.db.Exec(stmt, art.Title, art.Content, art.Author)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

// GetArticleByID ...
func (repo *repoImpl) GetArticleByID(id string) (*Article, error) {
	query := repo.db.QueryRow(`
		SELECT 
			id,
			title,
			content,
			author
		FROM 
			article 
		WHERE
			id = ? 
		ORDER BY id
	`, id)

	var art Article
	err := query.Scan(
		&art.ID,
		&art.Title,
		&art.Content,
		&art.Author,
	)

	return &art, err
}

func (repo *repoImpl) GetAllArticles() ([]Article, error) {
	stmt := `
		SELECT 
			id,
			title,
			content,
			author
		FROM 
			article
	`
	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Article
	for rows.Next() {
		var art Article
		err = rows.Scan(
			&art.ID,
			&art.Title,
			&art.Content,
			&art.Author,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, art)
	}
	err = rows.Err()

	return result, err

}
