package zaplogger

import "testing"

func TestZapLogger(t *testing.T) {
	logger := NewZapLogger(
		WithSyncFile("./log.log"))
	logger.Infof("test %s", "hello")
	logger.Info("sugar info wwww")
}
