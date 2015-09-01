package connections

// Connection -
type Connection interface {
	Start(done chan bool) error
	Validate() error
	String() string
	Stop() error
}
