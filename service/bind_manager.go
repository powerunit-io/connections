package service

import "fmt"

// BindManager -
type BindManager interface {
	Start(done chan bool) chan error
	Stop() error
	Validate() error
	Name() string
}

// Bind -
func (bs *BaseService) Bind(bind BindManager) error {
	bs.Debug("About to attach (bind: %s) to internal service bind management ...", bind.Name())

	if err := bind.Validate(); err != nil {
		bs.Error("Could not validate (bind: %s) due to (err: %s)", bind.Name(), err)
		return err
	}

	bs.Binds[bind.Name()] = bind
	return nil
}

func (bs *BaseService) ListAvailableBinds() []string {
	binds := []string{}

	for bind, _ := range bs.Binds {
		binds = append(binds, bind)
	}

	return binds
}

// GetBinds -
func (bs *BaseService) GetBinds() map[string]BindManager {
	return bs.Binds
}

// GetBind -
func (bs *BaseService) GetBind(b string) (BindManager, error) {
	if !bs.BindExists(b) {
		var bm BindManager
		return bm, fmt.Errorf("Could not retreive (bind: %s) as one is not attached.", b)
	}

	return bs.Binds[b], nil
}

// BindExists -
func (bs *BaseService) BindExists(b string) bool {
	for bind, _ := range bs.Binds {
		if bind == b {
			return true
		}
	}

	return false
}
