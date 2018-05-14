package logger

import (
	"testing"
)

func TestWriters(t *testing.T) {
	L := NewLogger()
	L.AllStdout()

	t.Run("Trace", func(t *testing.T) {
		L.Trace("Testing trace output")
	})
	t.Run("Info", func(t *testing.T) {
		L.Info("Testing info output")
	})
	t.Run("Warn", func(t *testing.T) {
		L.Warn("Testing warn output")
	})
	t.Run("Error", func(t *testing.T) {
		L.Error("Testing error output")
	})
	t.Run("Fatal", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Fatal didn't panic")
			}
		}()
		L.Fatal("Testing fatal output")
	})
}
