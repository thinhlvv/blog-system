package article

type (
	// Service is interface to interact with outside world.
	Service interface {
		// GetByID
		// GetArticles
		// CreateArticle
	}

	serviceImpl struct {
		repo Repository
	}
)

// NewService returns service implementation.
func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}
