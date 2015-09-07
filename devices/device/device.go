package device

// Base -
type Base struct {
	Name string `json:"device_name"`
	Type string `json:"device_type"`
}

// Start -
func (d *Base) Start(done chan bool) error {
	return nil
}

// Stop -
func (d *Base) Stop() error {
	return nil
}

// Validate -
func (d *Base) Validate() error {
	return nil
}

// DeviceName -
func (d *Base) DeviceName() string {
	return d.Name
}
