package zaplog

import "testing"

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(SyncInfoFile("./info.log"))
	logger.Info("hello", "zap looger")
	logger.Infof("hello %s", "zap looger")
}
