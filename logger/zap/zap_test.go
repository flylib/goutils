package zaplog

import "testing"

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(
		SyncFile("./info.log"),
		SyncConsole(),
		//MinPrintLevel(DebugLevel),
		//JsonFormat(),
	).Sugar()

	for i := 0; i < 10; i++ {
		logger.Info("hello", i)
		logger.Debug("hello", i)
		logger.Infof("hello %d", i)
	}
}
