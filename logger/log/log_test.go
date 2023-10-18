package log

import "testing"

func TestLogger(t *testing.T) {
	//newLogger := NewLogger(SyncFile("./log.log"))
	newLogger := NewLogger(SyncConsole())
	newLogger.Info("info")
	newLogger.Infof("info %d", 123)
	newLogger.Errorf("info %d", 123)
}
