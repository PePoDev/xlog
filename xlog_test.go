package xlog_test

import (
	"fmt"
	"testing"

	"github.com/pepodev/xlog"
)

func TestLog(t *testing.T) {
	fmt.Println(t.Name())
	xlog.Trace("Status is Trace")
	xlog.Debug("Status is Debug")
	xlog.Info("Status is Info")
	xlog.Warn("Status is Warning")
	xlog.Error("Status is Fatal")
	fmt.Println()
}

func TestLogFormatting(t *testing.T) {
	fmt.Println(t.Name())
	xlog.Tracef("%s %s", "Status is", "Trace")
	xlog.Debugf("%s %s", "Status is", "Debug")
	xlog.Debugf("%s %s", "Status is", "Debug")
	xlog.Infof("%s %s", "Status is", "Info")
	xlog.Warnf("%s %s", "Status is", "Warning")
	xlog.Errorf("%s %s", "Status is", "Fatal")
	fmt.Println()
}

func TestLogLine(t *testing.T) {
	fmt.Println(t.Name())
	xlog.Traceln("Status is", "Trace")
	xlog.Debugln("Status is", "Debug")
	xlog.Infoln("Status is", "Info")
	xlog.Warnln("Status is", "Warning")
	xlog.Errorln("Status is", "Fatal")
	fmt.Println()
}

func TestLogPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println(t.Name())
	xlog.Panicln("Status is", "Panic")
	fmt.Println()
}

func BenchmarkLogFormatting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xlog.Infof("%s %d", "Number is", i)
	}
}

func Example_log() {
	xlog.Trace("Trace")
	xlog.Debug("Debug")
	xlog.Info("Info")
	xlog.Warn("Warn")
	xlog.Error("Error")
	xlog.Panic("Panic")
	xlog.Fatal("Fatal")
	// Output:
	// Trace
	// Debug
	// Info
	// Warn
	// Error
	// Panic
	// Fatal
}
