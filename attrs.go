package slogf

import "log/slog"

// KeyError sets default key for [Error] func
var KeyError = "error"

// Error returns [slog.Attr] with error string as value and [KeyError] as key
// returns empty attr (which is skipped by slog) on nil error
func Error(err error) slog.Attr {
	if err == nil {
		return slog.Attr{}
	}
	return slog.String(KeyError, err.Error())
}

// Value is a helper function for [slog.LogValuer]
func Value(key string, valuer slog.LogValuer) slog.Attr {
	if valuer == nil {
		return slog.Attr{}
	}
	return slog.Attr{
		Key:   key,
		Value: valuer.LogValue(),
	}
}
