package article

// New return article controller service.
func New() *Controller {
	repo := NewRepo(nil)
	svc := NewService(repo)
	c := NewController(svc)
	return c
}
