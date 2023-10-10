package log

import "testing"

func TestLogger(t *testing.T) {
	newLogger := NewLogger(SyncFile("./log.log"))
	newLogger.Info("info")
	newLogger.Infof("info %d", 123)
	newLogger.Fatalf("info %d", 123)
}
