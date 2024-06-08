# slogf

## Key points
- logger is a simple `func(msg string, attrs ...slog.Attr)`
- minimalistic set of exports
- `slog.TextHandler` (aka [logfmt](https://brandur.org/logfmt)) by default
- no levels by default, errors are passed as just another key=value to logger
    - `slogf.Error` is exported for convenience

## If you want:

### Levels
Pass a custom `slog.Handler` to `slogf.New`, default one simply ignores them with `ReplaceAttr` in options

Put levels in logger invocation, or wrap them in separate loggers

### Verbosity
Create another `slogf.Logf` for verbose logs, control verbosity level on your own