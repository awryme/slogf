package slogf

import (
	"context"
	"log/slog"
	"slices"
	"time"
)

var backgroundCtx = context.Background()

// Logf is a func based logger
type Logf func(msg string, attrs ...slog.Attr)

// New creates a new [Logf]
func New(handler slog.Handler) Logf {
	return func(msg string, attrs ...slog.Attr) {
		r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, 0)
		r.AddAttrs(attrs...)
		handler.Handle(backgroundCtx, r)
	}
}

// With creates a new [Logf] with wrapped attrs
func (logf Logf) With(attrs ...slog.Attr) Logf {
	return func(msg string, msgAttrs ...slog.Attr) {
		logf(msg, slices.Concat(attrs, msgAttrs)...)
	}
}
