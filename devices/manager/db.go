package manager

import (
	"fmt"
	"strings"

	"github.com/powerunit-io/devices"
)

// ImportDevices -
func (dm *DeviceManager) ImportDevices() error {

	dm.Info("Starting device manager import ...")

	sql := `select d.id, b.id, b.name, f.id, f.name, r.id, r.name, r.room_type, d.name, d.type, d.device, d.protocol
	FROM devices AS d LEFT JOIN building AS b ON(b.id = d.building_id)
	LEFT JOIN floors AS f ON (f.id = d.floor_id) LEFT JOIN rooms AS r ON (r.id = d.room_id)
	WHERE d.status AND b.status AND r.status AND f.status ORDER by d.date_created ASC`

	dbconn := dm.Db.Conn()
	rows, err := dbconn.Query(sql)

	if err != nil {
		dm.Error("Could not select devices from database due to (err: %s)", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var deviceId, buildingId, floorId, roomId int64
		var buildingName, floorName, roomName, roomType, deviceName string
		var deviceType, device, deviceProtocol string

		if err := rows.Scan(
			&deviceId, &buildingId, &buildingName, &floorId, &floorName, &roomId,
			&roomName, &roomType, &deviceName, &deviceType, &device, &deviceProtocol,
		); err != nil {
			dm.Error("Could not select database row due to (err: %s)", err)
			return err
		}

		dm.Debug("Searching if (device: %s) is available ...", strings.ToLower(device))

		dn := strings.ToLower(device)

		if !devices.IsDeviceAvailable(dn) {
			return fmt.Error("Could not find (device: %s) ...", dn)
		}

		_, err := dm.AttachDevice(dn, devices.CreateDeviceByName(dn, map[string]interface{}{
			"device_id": deviceId,
		}))

		if err != nil {
			return fmt.Error("Could not attach (device: %s) due to (err: %s)", dn, err)
		}
	}

	return nil
}
