package service

// Service -
type Service interface {
	Start() error
	Stop() error

	Setup() error
	SyncDb() error

	Name() string
}
