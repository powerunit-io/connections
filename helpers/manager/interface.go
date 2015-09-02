package manager

import (
	"github.com/powerunit-io/platform/helpers/helper"
	"github.com/powerunit-io/platform/logging"
)

// Manager -
type Manager interface {
	GetHelpers() map[string]helper.Helper
	AttachHelper(n string, h helper.Helper) error
	RemoveHelper(n string) error
	ListAvailableHelpers() []string
	GetHelper(n string) (helper.Helper, error)
	HelperExists(n string) bool
}

// NewManager -
func NewManager(log *logging.Logger) Manager {
	wm := func(m Manager) Manager {
		return m
	}

	return wm(&HelperManager{
		Logger:  log,
		Helpers: make(map[string]helper.Helper),
	})
}
