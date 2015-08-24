package service

// Service -
type Service interface {
	Start() error
	Stop() error
}

// NewService - Will basically be used to instanciate service over generic
// interface that all services needs to met.
func NewService(service Service) Service {
	return service
}
