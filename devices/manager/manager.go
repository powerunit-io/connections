package manager

import (
	"fmt"

	"github.com/powerunit-io/platform/connections/mysql"
	"github.com/powerunit-io/platform/devices/device"
	"github.com/powerunit-io/platform/logging"
)

// DeviceManager -
type DeviceManager struct {
	*logging.Logger
	Db      mysql.Manager
	Devices map[string]device.Device
}

// GetDevices - Will return list of all available devices
func (dm *DeviceManager) GetDevices() map[string]device.Device {
	return dm.Devices
}

// AttachDevice -
func (dm *DeviceManager) AttachDevice(dn string, d device.Device) error {

	if dm.DeviceExists(dn) {
		return fmt.Errorf("Could not attach (device: %s) as one is already attached!", dn)
	}

	// Make sure that all is validated full before we go anywhere with it ...
	if err := d.Validate(); err != nil {
		return err
	}

	dm.Devices[dn] = d

	return nil
}

// RemoveDevice - Remove device from storage. Additionally, execute stop
func (dm *DeviceManager) RemoveDevice(dn string) error {
	if !dm.DeviceExists(dn) {
		return fmt.Errorf("Could not remove (device: %s) as one does not exist in store...", dn)
	}

	// Stop device before it's removed from internal storage ...
	if err := dm.Devices[dn].Stop(); err != nil {
		return err
	}

	delete(dm.Devices, dn)
	return nil
}

// ListAvailableDevices -
func (dm *DeviceManager) ListAvailableDevices() []string {
	devices := []string{}

	for device, _ := range dm.Devices {
		devices = append(devices, device)
	}

	return devices
}

// GetDevice -
func (dm *DeviceManager) GetDevice(dn string) (device.Device, error) {
	if !dm.DeviceExists(dn) {
		var d device.Device
		return d, fmt.Errorf("Could not retreive (device: %s) as one does not exist in store...", dn)
	}

	return dm.Devices[dn], nil
}

// DeviceExists - Return whenever device is set or not
func (dm *DeviceManager) DeviceExists(dn string) bool {
	for device, _ := range dm.Devices {
		if device == dn {
			return true
		}
	}

	return false
}
