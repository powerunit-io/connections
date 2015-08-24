package platform

import (
	"testing"

	"github.com/powerunit-io/platform/logging"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Sirupsen/logrus"
)

// TestLoggingManager - Just basic test to ensure that logging loger returns right
// logging context. In addition to that we'll check few additional methods such
// as context
func TestLoggingManager(t *testing.T) {
	logger := logging.New()

	Convey("Logging Manager Pointer Check", t, func() {
		So(*logger, ShouldHaveSameTypeAs, logging.Logger{})
	})

	Convey("We Expect Logrus Entry", t, func() {
		context := logger.GetContextLogger(map[string]interface{}{
			"common": "Wola",
		})

		So(context, ShouldHaveSameTypeAs, &logrus.Entry{})
	})

}
