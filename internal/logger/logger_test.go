package logger_test

import (
	"github.com/sirupsen/logrus"
	"rhyme-engine/internal/logger"
	"testing"
)

func Test_getLogLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want logrus.Level
	}{
		{"should return trace level for trace", args{"trace"}, logrus.TraceLevel},
		{"should return info level for info", args{"info"}, logrus.InfoLevel},
		{"should return debug level for debug", args{"debug"}, logrus.DebugLevel},
		{"should return error level for error", args{"error"}, logrus.ErrorLevel},
		{"should return fatal level for fatal", args{"fatal"}, logrus.FatalLevel},
		{"should return panic level for panic", args{"panic"}, logrus.PanicLevel},
		{"should return warn level for warn", args{"warning"}, logrus.WarnLevel},
		{"should return info level for invalid value", args{"abcd"}, logrus.InfoLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logger.GetLogLevel(tt.args.level); got != tt.want {
				t.Errorf("getLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
