package zaplog

import "testing"

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(WithSyncConsole(true), WithSyncWriteFile("./log/log.log"))
	logger.Info("hello %s", "zap looger")
}
