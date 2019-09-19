package article

import "database/sql"

// Repository is interface to interact to outside world.
type (
	// Repository is interface to interact with outside world.
	Repository interface {
		// GetArticleByID(types.ID)
	}

	repoImpl struct {
		*sql.DB
	}
)

// NewRepo returns repository implementation.
func NewRepo(db *sql.DB) Repository {
	return &repoImpl{DB: db}
}
