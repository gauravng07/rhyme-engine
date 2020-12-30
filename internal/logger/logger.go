package logger

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"rhyme-engine/internal"
	"rhyme-engine/internal/config"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logger *logrus.Logger

type Level uint32

const CorrelationId = "correlation-id"

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func init() {
	logger = logrus.New()
	formatter := logrus.JSONFormatter{}
	logger.SetFormatter(&formatter)
	logger.SetLevel(GetLogLevel(viper.GetString(config.LogLevel)))
	logger.SetOutput(os.Stdout)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Fatalf(format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Infof(format, args...)
}

func Info(ctx context.Context, msg string) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Info(msg)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	formattedError := escapeString(format, args...)
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Debug(formattedError)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Warnf(format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	formattedError := escapeString(format, args...)
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Error(formattedError)
}

func escapeString(format string, args ...interface{}) string {
	errorMessage := fmt.Sprintf(format, args...)
	re := regexp.MustCompile(`(\n)|(\r\n)`)
	formattedError := re.ReplaceAllString(errorMessage, "\\n ")
	return formattedError
}

type customFormatter struct{}

func (c customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	logMessage := ""
	correlationId := entry.Data[CorrelationId]

	if correlationId != nil {
		logMessage = fmt.Sprintf("[%s] [%s] %s", strings.ToUpper(entry.Level.String()), correlationId, entry.Message)
	} else {
		logMessage = fmt.Sprintf("[%s] [] %s", strings.ToUpper(entry.Level.String()), entry.Message)
	}

	b.WriteString(logMessage)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func GetLogLevel(level string) logrus.Level {
	switch level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
