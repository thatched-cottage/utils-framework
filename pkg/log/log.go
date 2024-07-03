package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

// Option 使用函数定义一个配置选项类型
type Option func(*ConfigOption)

// ConfigOption 持有可选参数
type ConfigOption struct {
	Out io.Writer
}

func WithFileOut(logFile string) Option {
	return func(cfg *ConfigOption) {
		out, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		cfg.Out = out
	}
}

func WithRotatingFileWriter(logFilePath, logFileName string) Option {
	return func(cfg *ConfigOption) {
		cfg.Out = NewRotatingFileWriter(logFilePath, logFileName)
	}
}

// Init 替换以将日志输出到文件
func Init(opts ...Option) error {
	cfg := &ConfigOption{
		Out: os.Stdout,
	}

	// 应用每个Option函数来设置配置
	for _, opt := range opts {
		opt(cfg)
	}
	logSingleton.Logger = log.New(cfg.Out, "Text:", log.LstdFlags)
	return nil
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
