package logger

import (
	"context"
	"log/slog"
)

type customLogHandler struct {
	sh slog.Handler
}

type TraceIdKey string

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, TraceIdKey("traceId"), traceId)
}

func NewCustomLogHandler(sh slog.Handler) *customLogHandler {
	return &customLogHandler{sh: sh}
}

func (h *customLogHandler) Handle(ctx context.Context, rc slog.Record) error {
	// トレースIDがコンテキストにセットされている場合はログに追加
	traceId, ok := ctx.Value(TraceIdKey("traceId")).(string)
	if ok {
		rc.Add(string(TraceIdKey("traceId")), traceId)
	}

	return h.sh.Handle(ctx, rc)
}

func (h *customLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.sh.Enabled(ctx, level)
}

func (h *customLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &customLogHandler{h.sh.WithAttrs(attrs)}
}

func (h *customLogHandler) WithGroup(name string) slog.Handler {
	return &customLogHandler{h.sh.WithGroup(name)}
}
