package ylog

import (
	"testing"

	"github.com/samkon/yerror"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWhenInitLoggerThenLoggerConfigShouldFill(t *testing.T) {
	Init(MConfig("debug"))

	loggerAfterInit := Log

	assert.NotEmpty(t, loggerAfterInit)

	Clear()
}

func TestNewLoggerWhenLoggerConfigPropertyEmptyAndNotInitLoggerThenShouldReturnError(t *testing.T) {
	logger, err := New()

	assert.Error(t, err)
	assert.Empty(t, logger)
}

func TestNewLoggerWhenLoggerConfigPropertyEmptyAndInitLoggerThenShouldReturnLogger(t *testing.T) {
	Init(MConfig("debug"))
	logger, err := New()

	assert.NoError(t, err)
	assert.NotEmpty(t, logger)

	Clear()
}

func TestNewLoggerWithCustomConfigPropertyThenShouldReturnLogger(t *testing.T) {
	logger, err := New(MConfig("debug"))

	assert.NoError(t, err)
	assert.NotEmpty(t, logger)
}

func TestNewLoggerWithCustomMissingConfigPropertyThenShouldReturnError(t *testing.T) {
	logger, err := New(zap.Config{})

	assert.Error(t, err)
	assert.Empty(t, logger)
}

func TestWhenInitLoggerWithEmptyConfigPropertyThenShouldReturnLogger(t *testing.T) {
	Init()

	loggerAfterInit := Log

	assert.NotEmpty(t, loggerAfterInit)

	Clear()
}

func TestGivenErrIsWrappedWhenErrorCalledThenShouldPrintErrorMessage(t *testing.T) {
	Init()
	merr := yerror.New("this is an error", yerror.Code(500), zap.String("key", "value"))
	wrapped := yerror.Wrap(merr, zap.Int("key2", 1))
	Log.Error("testing error", zap.Error(wrapped))
	Clear()
}
