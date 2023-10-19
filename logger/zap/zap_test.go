package zaplog

import (
	"testing"
)

func TestZaplog(t *testing.T) {
	sugaredLogger := NewZapLogger(
		SyncFile("./info.log"),
		SyncConsole(),
		//MinPrintLevel(DebugLevel),
		//JsonFormat(),
	).Sugar()

	for i := 0; i < 10; i++ {
		sugaredLogger.Info("hello", i)
		sugaredLogger.Debug("hello", i)
		sugaredLogger.Infof("hello %d", i)
	}

}
