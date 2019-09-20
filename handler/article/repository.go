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
		GetArticleByID(id int) (*Article, error)
		GetArticles() ([]Article, error)
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
		VALUES (%s, %s, %s)
	`, art.Title, art.Content, art.Author)

	res, err := repo.db.Exec(stmt)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

// GetArticleByID ...
func (repo *repoImpl) GetArticleByID(id int) (*Article, error) {
	query := repo.db.QueryRow(`
		SELECT 
			id,
			title,
			content,
			author,
			created_at,
			updated_at
		WHERE
			id = ? AND 
			deleted_at = NULL
		ORDER BY id
		LIMIT 1
	`)

	var art Article
	err := query.Scan(
		&art.ID,
		&art.Title,
		&art.Content,
		&art.Author,
		&art.CreatedAt,
		&art.UpdatedAt,
	)

	return &art, err
}

func (repo *repoImpl) GetArticles() ([]Article, error) {
	stmt := `
		SELECT 
			id,
			title,
			content,
			author,
			created_at,
			updated_at
		WHERE
			deleted_at = NULL
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
			&art.CreatedAt,
			&art.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, art)
	}
	err = rows.Err()

	return result, err

}
