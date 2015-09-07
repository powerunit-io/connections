package manager

import (
	"github.com/powerunit-io/platform/connections/mysql"
	"github.com/powerunit-io/platform/devices/device"
	"github.com/powerunit-io/platform/logging"
)

// Manager -
type Manager interface {
	GetDevices() map[string]device.Device
	AttachDevice(dn string, w device.Device) error
	RemoveDevice(dn string) error
	ListAvailableDevices() []string
	GetDevice(dn string) (device.Device, error)
	DeviceExists(dn string) bool

	ImportDevices() error
}

// NewDeviceManager -
func NewDeviceManager(log *logging.Logger, db mysql.Manager) Manager {
	return Manager(&DeviceManager{
		Logger:  log,
		Db:      db,
		Devices: make(map[string]device.Device),
	})
}
