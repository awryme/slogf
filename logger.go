package slogf

import (
	"context"
	"log/slog"
	"slices"
	"time"
)

var backgroundCtx = context.Background()

type Logf func(msg string, attrs ...slog.Attr)

func New(handler slog.Handler) Logf {
	return func(msg string, attrs ...slog.Attr) {
		r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, 0)
		r.AddAttrs(attrs...)
		handler.Handle(backgroundCtx, r)
	}
}

func (logf Logf) With(attrs ...slog.Attr) Logf {
	return func(msg string, msgAttrs ...slog.Attr) {
		logf(msg, slices.Concat(attrs, msgAttrs)...)
	}
}
