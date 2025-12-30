package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yamada-mikiya/team1-hackathon/api"
	"github.com/yamada-mikiya/team1-hackathon/config"
	"github.com/yamada-mikiya/team1-hackathon/database"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	// 設定読み込み
	cfg, err := config.GetConfig()
	if err != nil {
		slog.Error("設定の読み込みに失敗しました", "error", err)
		return 1
	}

	// マイグレーション実行
	if err := database.RunMigrations(cfg.Database.GetDSN()); err != nil {
		slog.Error("マイグレーションに失敗しました", "error", err)
		return 1
	}

	// データベース接続
	db, err := database.ConnectDB(cfg.Database.GetDSN())
	if err != nil {
		slog.Error("データベース接続に失敗しました", "error", err)
		return 1
	}
	defer database.Close(db)

	router := api.SetupRouter(cfg)

	// サーバー設定
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	// サーバー起動
	errChan := make(chan error, 1)
	go func() {
		slog.Info("サーバーを起動しました", "port", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// シグナルハンドリング
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		slog.Info("終了シグナルを受信しました。シャットダウンを開始します...")
	case err := <-errChan:
		slog.Error("サーバーエラーが発生しました", "error", err)
		return 1
	}

	// グレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("サーバーのシャットダウンに失敗しました", "error", err)
		return 1
	}

	slog.Info("サーバーが正常に終了しました")
	return 0
}
