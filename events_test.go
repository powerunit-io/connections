package platform

import (
	"testing"

	"github.com/powerunit-io/platform/events"
	. "github.com/smartystreets/goconvey/convey"
)

type TestMessage struct {
	duplicate bool
	qos       byte
	retained  bool
	topic     string
	messageID uint16
	payload   []byte
}

func (m *TestMessage) Duplicate() bool {
	return m.duplicate
}

func (m *TestMessage) Qos() byte {
	return m.qos
}

func (m *TestMessage) Retained() bool {
	return m.retained
}

func (m *TestMessage) Topic() string {
	return m.topic
}

func (m *TestMessage) MessageID() uint16 {
	return m.messageID
}

func (m *TestMessage) Payload() []byte {
	return m.payload
}

var (
	TestMsgBedroomDhtSensor = `{"floor": "middle", "room": "bedroom", "plc": "bedroom-main-controller", "device": "dht-sensor", "device_type": "sensor", "data": { "celsius": 20, "fahrenheit": 80, "heat_index_celsius": 29.20, "heat_index_fahrenheit": 82.33, "humidity_percentage": 40 }}`
)

// TestNewEventCreation - Just basic test to ensure that logging loger returns right
// logging context. In addition to that we'll check few additional methods such
// as context
func TestNewEventCreation(t *testing.T) {

	msg := TestMessage{
		false, byte(0), true, "powerunit-io-bridge", 01, []byte(TestMsgBedroomDhtSensor),
	}

	Convey("Test If Proper Type", t, func() {
		e, err := events.NewEvent(&msg)
		So(err, ShouldHaveSameTypeAs, nil)
		So(e, ShouldHaveSameTypeAs, events.Event{})
	})

}
