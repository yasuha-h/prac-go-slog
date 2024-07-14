package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/yasuha-h/logger"
)

func main() {
	// トレースIDをUUIDで生成
	traceId := uuid.New().String()
	// トレースIDをコンテキストにセット
	ctx := logger.WithTraceId(context.Background(), traceId)

	// accessLog
	accessLog, err := os.OpenFile("./logs/access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// アプリケーション用のカスタムロガーを作成
	accessLogHandler := logger.NewCustomLogHandler(
		slog.NewJSONHandler(accessLog, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	slog.New(accessLogHandler).DebugContext(ctx, "access log")

	// appLog
	appLog, err := os.OpenFile("./logs/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// アプリケーション用のカスタムロガーを作成
	appLogHandler := logger.NewCustomLogHandler(
		slog.NewJSONHandler(
			appLog,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		))
	// カスタムロガーをセット
	slog.SetDefault(slog.New(appLogHandler))

	// ログ出力 トレースIDをコンテキスト経由で共通で出力できているか確認
	slog.InfoContext(ctx, "app1")
	slog.InfoContext(ctx, "app2")
	slog.InfoContext(ctx, "app3")
	// コンテキストを使わないとトレースIDが出力されない
	slog.Info("app4")

	// accessLogとLogLevelが異なるため出力されない
	slog.Debug("Not outputed.")
}
