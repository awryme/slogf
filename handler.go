package slogf

import (
	"io"
	"log/slog"
)

// DefaultHandler returns a [slog.TextHandler] that skips level key
func DefaultHandler(output io.Writer) slog.Handler {
	opt := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == slog.LevelKey {
				return slog.Attr{}
			}
			return a
		},
	}
	return slog.NewTextHandler(output, opt)
}
