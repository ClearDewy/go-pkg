/**
 * @ Author: ClearDewy
 * @ Desc:
 **/
package logrus

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type LogrusFormatter struct{}

func (f *LogrusFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	message := entry.Message

	var levelColor func(...interface{}) string
	switch entry.Level {
	case logrus.ErrorLevel:
		levelColor = color.New(color.FgRed).SprintFunc()
	case logrus.WarnLevel:
		levelColor = color.New(color.FgYellow).SprintFunc()
	case logrus.InfoLevel:
		levelColor = color.New(color.FgGreen).SprintFunc()
	case logrus.DebugLevel:
		levelColor = color.New(color.FgWhite).SprintFunc()
	default:
		levelColor = color.New(color.FgWhite).SprintFunc()
	}

	// Split the message into lines and add padding to all lines except the first
	lines := strings.Split(message, "\n")
	for i := 1; i < len(lines); i++ {
		lines[i] = fmt.Sprintf("%-7s\t%s\t%s", "", "", lines[i]) // Adding padding
	}
	// Join the lines back together
	message = strings.Join(lines, "\n")

	return []byte(fmt.Sprintf("%-7s\t%s\t%s\n", levelColor(level), timestamp, message)), nil
}

func New() *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(LogrusFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}

var (
	// std is the name of the standard logger in stdlib `log`
	std = New()
)

func StandardLogger() *logrus.Logger {
	return std
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	std.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	std.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func SetLevel(level logrus.Level) {
	std.SetLevel(level)
}

// GetLevel returns the standard logger level.
func GetLevel() logrus.Level {
	return std.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level logrus.Level) bool {
	return std.IsLevelEnabled(level)
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook logrus.Hook) {
	std.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return std.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *logrus.Entry {
	return std.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return std.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields logrus.Fields) *logrus.Entry {
	return std.WithFields(fields)
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return std.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	std.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	std.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	std.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn logrus.LogFunction) {
	std.TraceFn(fn)
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn logrus.LogFunction) {
	std.DebugFn(fn)
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn logrus.LogFunction) {
	std.PrintFn(fn)
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn logrus.LogFunction) {
	std.InfoFn(fn)
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn logrus.LogFunction) {
	std.WarnFn(fn)
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn logrus.LogFunction) {
	std.WarningFn(fn)
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn logrus.LogFunction) {
	std.ErrorFn(fn)
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn logrus.LogFunction) {
	std.PanicFn(fn)
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn logrus.LogFunction) {
	std.FatalFn(fn)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	std.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	std.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

var Skip = 3

func SetErrorSkip(skip int) {
	if skip < 1 {
		return
	}
	Skip = skip
}

func ErrorM(err error, msg string) {
	if err == nil {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		std.Errorf("%s\n.At %s:%d. \n%v", msg, file, line, err)
	} else {
		std.Errorf("%s\n.Can't get the location. \n%v", msg, err)
	}
	for i := 2; i <= Skip; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			std.Errorf("At %s:%d. \n%v", file, line, err)
		} else {
			break
		}
	}
}
