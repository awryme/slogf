package slogf

import "log/slog"

const KeyError = "error"

func Error(err error) slog.Attr {
	if err == nil {
		return slog.Attr{}
	}
	return slog.String(KeyError, err.Error())
}

func Value(key string, valuer slog.LogValuer) slog.Attr {
	if valuer == nil {
		return slog.Attr{}
	}
	return slog.Attr{
		Key:   key,
		Value: valuer.LogValue(),
	}
}
