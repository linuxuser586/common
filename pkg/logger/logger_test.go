package logger

import "testing"

func TestLogger(t *testing.T) {
	z := Zap()
	z.Info("testing")
}
