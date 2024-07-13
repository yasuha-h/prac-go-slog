package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/yasuha-h/logger"
)

func main() {
	// カスタムロガーを作成
	h := logger.NewCustomLogHandler(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
	))
	// カスタムロガーをセット
	slog.SetDefault(slog.New(h))

	// トレースIDをUUIDで生成
	traceId := uuid.New().String()
	// トレースIDをコンテキストにセット
	ctx := logger.WithTraceId(context.Background(), traceId)

	// ログ出力 トレースIDをコンテキスト経由で共通で出力できているか確認
	slog.InfoContext(ctx, "Messageeeeee")
	slog.InfoContext(ctx, "Messageeeeee")
	slog.InfoContext(ctx, "Messageeeeee")
	// コンテキストを使わないとトレースIDが出力されない
	slog.Info("Messageeeeee")
}
