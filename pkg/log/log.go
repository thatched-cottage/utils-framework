package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func Init() {
	logSingleton.Logger = log.New(os.Stdout, "Test:", log.LstdFlags)
}

var logSingleton ServiceProviderLogger

type ServiceProviderLogger struct {
	*log.Logger
}

func defaultLogger() *log.Logger {
	if nil == logSingleton.Logger {
		Init()
	}
	return logSingleton.Logger
}

func NewLogger(logger *log.Logger) {
	logSingleton.Logger = logger
}

func logHeader(format string) string {
	pc, _, lineNo, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("func:%s:%d [%s]",
		runtime.FuncForPC(pc).Name(), lineNo, format)
}

func logHeaderB(args ...any) string {
	pc, _, lineNo, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("func:%s:%d [%s]",
		runtime.FuncForPC(pc).Name(), lineNo, args)
}

func Tracef(format string, args ...any) {
	defaultLogger().Printf(logHeader("[Trace]:["+format+"]"), args...)
}

func Debug(format string) {
	defaultLogger().Printf(logHeaderB("[Debug]:[%s]"), format)
}

func Debugf(format string, args ...any) {
	defaultLogger().Printf(logHeader("[Debug]:["+format+"]"), args...)
}

func Infof(format string, args ...any) {
	defaultLogger().Printf(logHeader("[Info]:["+format+"]"), args...)
}

func Warnf(format string, args ...any) {
	defaultLogger().Printf(logHeader("[Warn]:["+format+"]"), args...)
}

func Errorf(format string, args ...any) {
	defaultLogger().Printf(logHeader("[Error]:"+format+"]"), args...)
}

func Fatalf(format string, args ...any) {
	defaultLogger().Fatalf(logHeader("[Fatal]:"+format+"]"), args...)
}

func Info(args ...any) {
	defaultLogger().Print(logHeader("[Info]:[" + fmt.Sprint(args...) + "]"))
}

func Error(args ...any) {
	defaultLogger().Print(logHeader("[Error]:[" + fmt.Sprint(args...) + "]"))
}
