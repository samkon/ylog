package ylog

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lf zap.Config

var Log, _ = New(MConfig("info")) // If not initialized use MConfig("info") with info level

func Init(config ...zap.Config) {
	if len(config) > 0 {
		lf = config[0]
	} else {
		lf = MConfig("info")
	}
	Log, _ = New(lf)
}

func Clear() {
	lf = zap.Config{}
	Log = nil
}

func New(config ...zap.Config) (Logger, error) {
	if len(config) > 0 {
		lg, err := config[0].Build(zap.AddCallerSkip(1))
		if err != nil {
			return nil, err
		}
		return &logger{logger: lg}, nil
	}
	lg, err := lf.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &logger{logger: lg}, nil
}

func NewNop() Logger {
	return &logger{logger: zap.NewNop()}
}

func getLogLevel(level string) zapcore.Level {
	switch levelFromConfig := strings.TrimSpace(level); {
	case strings.EqualFold(levelFromConfig, "debug"):
		return zapcore.DebugLevel
	case strings.EqualFold(levelFromConfig, "error"):
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func MConfig(logLevel string) zap.Config {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	loggerConfig.Level = zap.NewAtomicLevelAt(getLogLevel(logLevel))
	loggerConfig.EncoderConfig.CallerKey = "caller"

	return loggerConfig
}
