package slog

import (
	"sync/atomic"
	"time"
)

func New() *Logger {
	l := NewLogger(
		OptPart(
			PartLevel(),
			PartDateTime(time.RFC3339),
			PartCaller(true),
			PartMessage(),
		),
	)
	return l
}

func NewDailyLogger(folder, prefix string) *Logger {
	l := NewLogger(
		OptOutput(
			NewDailyWriter(folder, prefix),
		),
		OptPart(
			PartLevel(),
			PartDateTime(time.RFC3339),
			PartCaller(true),
			PartMessage(),
		),
	)

	return l
}

var defaultLogger atomic.Value

func init() {
	SetDefault(New())
}

func Default() *Logger {
	return defaultLogger.Load().(*Logger)
}

func SetDefault(l *Logger) *Logger {
	defaultLogger.Store(l)
	return Default()
}

func Debug(i ...any)                 { Default().Log(LevelDebug, "", i) }
func Debugf(format string, i ...any) { Default().Log(LevelDebug, format, i) }
func Info(i ...any)                  { Default().Log(LevelInfo, "", i) }
func Infof(format string, i ...any)  { Default().Log(LevelInfo, format, i) }
func Warn(i ...any)                  { Default().Log(LevelWarn, "", i) }
func Warnf(format string, i ...any)  { Default().Log(LevelWarn, format, i) }
func Err(i ...any)                   { Default().Log(LevelError, "", i) }
func Errf(format string, i ...any)   { Default().Log(LevelError, format, i) }
func With(parts ...any) ILogger      { return Default().With(parts...) }
