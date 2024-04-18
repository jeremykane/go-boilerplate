package logger

import "context"

const (
	// Debug has verbose message
	DebugLevel = "debug"
	// Info is default log level
	InfoLevel = "info"
	// Warn is for logging messages about possible issues
	WarnLevel = "warn"
	// Error is for logging errors
	ErrorLevel = "error"
	// Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	FatalLevel = "fatal"
)

var log Logger

type Fields map[string]interface{}

// Logger is our contract for the logger
type Logger interface {
	Tracef(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})

	Trace(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})

	Traceln(ctx context.Context, args ...interface{})
	Debugln(ctx context.Context, args ...interface{})
	Infoln(ctx context.Context, args ...interface{})
	Warningln(ctx context.Context, args ...interface{})
	Errorln(ctx context.Context, args ...interface{})
	Fatalln(ctx context.Context, args ...interface{})

	TraceWithFields(ctx context.Context, msg string, fields Fields)
	DebugWithFields(ctx context.Context, msg string, fields Fields)
	InfoWithFields(ctx context.Context, msg string, fields Fields)
	WarningWithFields(ctx context.Context, msg string, fields Fields)
	ErrorWithFields(ctx context.Context, msg string, fields Fields)
	FatalWithFields(ctx context.Context, msg string, fields Fields)
}

// NewLogger returns an instance of logger
func NewLogger(l Logger) {
	log = l
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	log.Tracef(ctx, format, args...)
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	log.Infof(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	log.Warningf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	log.Fatalf(ctx, format, args...)
}

func Trace(ctx context.Context, args ...interface{}) {
	log.Trace(ctx, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	log.Debug(ctx, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	log.Info(ctx, args...)
}

func Warning(ctx context.Context, args ...interface{}) {
	log.Warning(ctx, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	log.Error(ctx, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	log.Fatal(ctx, args...)
}

func Traceln(ctx context.Context, args ...interface{}) {
	log.Traceln(ctx, args...)
}

func Debugln(ctx context.Context, args ...interface{}) {
	log.Debugln(ctx, args...)
}

func Infoln(ctx context.Context, args ...interface{}) {
	log.Infoln(ctx, args...)
}

func Warningln(ctx context.Context, args ...interface{}) {
	log.Warningln(ctx, args...)
}

func Errorln(ctx context.Context, args ...interface{}) {
	log.Errorln(ctx, args...)
}

func TraceWithFields(ctx context.Context, msg string, fields Fields) {
	log.TraceWithFields(ctx, msg, fields)
}

func DebugWithFields(ctx context.Context, msg string, fields Fields) {
	log.DebugWithFields(ctx, msg, fields)
}

func InfoWithFields(ctx context.Context, msg string, fields Fields) {
	log.InfoWithFields(ctx, msg, fields)
}

func WarningWithFields(ctx context.Context, msg string, fields Fields) {
	log.WarningWithFields(ctx, msg, fields)
}

func ErrorWithFields(ctx context.Context, msg string, fields Fields) {
	log.ErrorWithFields(ctx, msg, fields)
}

func FatalWithFields(ctx context.Context, msg string, fields Fields) {
	log.FatalWithFields(ctx, msg, fields)
}
