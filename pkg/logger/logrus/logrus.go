package logrus

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/jeremykane/go-boilerplate/pkg/logger"
	"github.com/sirupsen/logrus"
)

type (
	logClient struct {
		lgs           logrus.Logger
		defaultFields logger.Fields
	}

	LogrusParam struct {
		Level string
	}
)

func getLevelLog(level string) logrus.Level {
	switch level {
	case logger.InfoLevel:
		return logrus.InfoLevel
	case logger.WarnLevel:
		return logrus.WarnLevel
	case logger.DebugLevel:
		return logrus.DebugLevel
	case logger.ErrorLevel:
		return logrus.ErrorLevel
	case logger.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func NewLogrus(param LogrusParam) logger.Logger {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(getFormatterJSON())

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(getLevelLog(param.Level))
	logrus.SetReportCaller(true)
	return &logClient{
		lgs: logrus.Logger{},
	}
}

// getFormatterJSON to generate JSON formatter object used by logrus
func getFormatterJSON() logrus.Formatter {
	return &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			_, file, line, ok := runtime.Caller(9)
			if !ok {
				file = "<???>"
				line = 1
			}
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	}
}

func (l *logClient) Tracef(ctx context.Context, format string, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Tracef(format, args...)
}

func (l *logClient) Debugf(ctx context.Context, format string, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Debugf(format, args...)
}

func (l *logClient) Infof(ctx context.Context, format string, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Infof(format, args...)
}

func (l *logClient) Warningf(ctx context.Context, format string, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Warningf(format, args...)
}

func (l *logClient) Errorf(ctx context.Context, format string, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Errorf(format, args...)
}

func (l *logClient) Fatalf(ctx context.Context, format string, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Fatalf(format, args...)
}

func (l *logClient) Trace(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Trace(args...)
}
func (l *logClient) Debug(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Debug(args...)
}

func (l *logClient) Info(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Info(args...)
}

func (l *logClient) Warning(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Warning(args...)
}

func (l *logClient) Error(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Error(args...)
}

func (l *logClient) Fatal(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Fatal(args...)
}

func (l *logClient) Traceln(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Traceln(args...)
}

func (l *logClient) Debugln(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Debugln(args...)
}

func (l *logClient) Infoln(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Infoln(args...)
}

func (l *logClient) Warningln(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Warningln(args...)
}

func (l *logClient) Errorln(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Errorln(args...)
}

func (l *logClient) Fatalln(ctx context.Context, args ...interface{}) {
	fields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, make(map[string]interface{})))
	logrus.WithFields(logrus.Fields(fields)).Fatalln(args...)
}

func (l *logClient) TraceWithFields(ctx context.Context, msg string, fields logger.Fields) {
	newFields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, fields))
	logrus.WithContext(ctx).WithFields(logrus.Fields(newFields)).Traceln(msg)
}

func (l *logClient) DebugWithFields(ctx context.Context, msg string, fields logger.Fields) {
	newFields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, fields))
	logrus.WithContext(ctx).WithFields(logrus.Fields(newFields)).Debugln(msg)
}

func (l *logClient) InfoWithFields(ctx context.Context, msg string, fields logger.Fields) {
	newFields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, fields))
	logrus.WithContext(ctx).WithFields(logrus.Fields(newFields)).Infoln(msg)
}

func (l *logClient) WarningWithFields(ctx context.Context, msg string, fields logger.Fields) {
	newFields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, fields))
	logrus.WithContext(ctx).WithFields(logrus.Fields(newFields)).Warningln(msg)
}

func (l *logClient) ErrorWithFields(ctx context.Context, msg string, fields logger.Fields) {
	newFields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, fields))
	logrus.WithContext(ctx).WithFields(logrus.Fields(newFields)).Errorln(msg)
}

func (l *logClient) FatalWithFields(ctx context.Context, msg string, fields logger.Fields) {
	newFields := MergeFields(l.defaultFields, MergeContextWithFields(ctx, fields))
	logrus.WithContext(ctx).WithFields(logrus.Fields(newFields)).Fatalln(msg)
}
