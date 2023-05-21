package zaplogger

import "testing"

func TestZapLogger(t *testing.T) {
	logger := NewZapLogger(
		WithSyncFile("./log.log"))
	logger.Info("sugar info wwww")
	logger.Warn("sugar info wwww")

	logger.Sugar().Infof("%s", "tst")
	logger.Sugar().Warnf("%s", "tst")
}
