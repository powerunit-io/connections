package connections

// Connection -
type Connection interface {
	Start() error
	IsConnected() bool
	Stop() error
}
