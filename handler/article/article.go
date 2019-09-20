package article

import "database/sql"

// New return article controller service.
func New(db *sql.DB) *Controller {
	repo := NewRepo(db)
	svc := NewService(repo)
	c := NewController(svc)
	return c
}
