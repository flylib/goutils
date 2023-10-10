package zaplog

import "testing"

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(SyncInfoFile("./log/info.log"))
	logger.Info("hello", "zap looger")
	logger.Infof("hello %s", "zap looger")
}
