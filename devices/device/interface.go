package device

// Device -
type Device interface {
	Start(done chan bool) error
	Stop() error
	Validate() error
	DeviceName() string
}
