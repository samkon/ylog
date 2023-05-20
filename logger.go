package ylog

import (
	"github.com/samkon/yerror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(msg string, fields ...zapcore.Field)
	Debug(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Panic(msg string, fields ...zapcore.Field)
	Warning(msg string, fields ...zapcore.Field)
	With(fields ...zapcore.Field) Logger
	Sync() error
	ConvertToZapLogger() *zap.Logger
}

type logger struct {
	logger *zap.Logger
}

func (l *logger) With(fields ...zapcore.Field) Logger {
	return &logger{l.logger.With(fields...)}
}

func (l *logger) ConvertToZapLogger() *zap.Logger {
	return l.logger
}

func (l *logger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *logger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zapcore.Field) {
	var flds []zapcore.Field

	for _, field := range fields {
		if field.Type == zapcore.ErrorType {
			if yerror.IsMerror(field.Interface.(error)) {
				m := yerror.AsMerror(field.Interface.(error))
				flds = append(flds, m.GetFields()...)
				flds = append(flds, zap.Any("error", m.GetMessage()))
				continue
			} else {
				flds = append(flds, zap.String("error", field.Interface.(error).Error()))
				continue
			}
		}
		flds = append(flds, field)
	}
	l.logger.WithOptions(zap.AddStacktrace(zapcore.DPanicLevel)).Error(msg, flds...)
}

func (l *logger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *logger) Panic(msg string, fields ...zapcore.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *logger) Warning(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}
