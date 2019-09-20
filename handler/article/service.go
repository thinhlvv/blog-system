package article

type (
	// Service is interface to interact with outside world.
	Service interface {
		CreateArticle(createReq) (*Article, error)
		GetAllArticles() ([]Article, error)
		GetArticleByID(id int) (*Article, error)
	}

	serviceImpl struct {
		repo Repository
	}
)

// NewService returns service implementation.
func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

// CreateArticle ...
func (s *serviceImpl) CreateArticle(req createReq) (*Article, error) {
	var art Article
	art.Title = req.Title
	art.Content = req.Content
	art.Author = req.Author

	id, err := s.repo.CreateArticle(art)
	if err != nil {
		return nil, err
	}
	art.ID = id

	return &art, nil
}

// GetAllArticles ...
func (s *serviceImpl) GetAllArticles() ([]Article, error) {
	return s.repo.GetAllArticles()
}

// GetArticleByID ...
func (s *serviceImpl) GetArticleByID(id int) (*Article, error) {
	return s.repo.GetArticleByID(id)
}
