package manager

import (
	"fmt"

	"github.com/powerunit-io/platform/helpers/helper"
	"github.com/powerunit-io/platform/logging"
)

// HelperManager -
type HelperManager struct {
	*logging.Logger
	Helpers map[string]helper.Helper
}

// GetHelpers -
func (hm *HelperManager) GetHelpers() map[string]helper.Helper {
	return hm.Helpers
}

// AttachHelper -
func (hm *HelperManager) AttachHelper(n string, h helper.Helper) error {

	if hm.HelperExists(n) {
		return fmt.Errorf("Could not attach (helper: %s) as one is already attached!", h)
	}

	// Make sure that all is validated full before we go anywhere with it ...
	if err := h.Validate(); err != nil {
		return err
	}

	hm.Helpers[n] = h

	return nil
}

// RemoveHelper - Will check if helper exists (returns error) and if does will unset it from helpers collection
func (hm *HelperManager) RemoveHelper(n string) error {
	if !hm.HelperExists(n) {
		return fmt.Errorf("Could not remove (helper: %s) as one does not exist in store...", n)
	}

	delete(hm.Helpers, n)
	return nil
}

// ListAvailableHelpers - Will fetch list of all available helpers
func (hm *HelperManager) ListAvailableHelpers() []string {
	helpers := []string{}

	for helper, _ := range hm.Helpers {
		helpers = append(helpers, helper)
	}

	return helpers
}

// GetHelper -
func (hm *HelperManager) GetHelper(n string) (helper.Helper, error) {
	if !hm.HelperExists(n) {
		var h helper.Helper
		return h, fmt.Errorf("Could not retreive (helper: %s) as one does not exist in store...", n)
	}

	return hm.Helpers[n], nil
}

// HelperExists -
func (hm *HelperManager) HelperExists(n string) bool {
	for helper, _ := range hm.Helpers {
		if helper == n {
			return true
		}
	}

	return false
}
