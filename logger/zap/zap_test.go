package zaplog

import "testing"

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(SyncFile("./info.log"))

	for i := 0; i < 10024; i++ {
		logger.Info("hello", "zap looger")
		logger.Infof("hello %s", "zap looger")
	}

}
