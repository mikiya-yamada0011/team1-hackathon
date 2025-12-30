package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB はデータベースに接続
func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("データベース接続に失敗しました", "error", err)
		return nil, fmt.Errorf("データベース接続失敗: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("sql.DBの取得に失敗しました", "error", err)
		return nil, fmt.Errorf("sql.DB取得失敗: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		slog.Error("データベースへのPingに失敗しました", "error", err)
		return nil, fmt.Errorf("データベースPing失敗: %w", err)
	}

	slog.Info("データベースに接続しました")

	return db, nil
}

// Close はデータベース接続をクローズ
func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("sql.DBの取得に失敗しました", "error", err)
		return fmt.Errorf("sql.DB取得失敗: %w", err)
	}

	slog.Info("データベース接続をクローズしています...")
	if err := sqlDB.Close(); err != nil {
		slog.Error("データベース接続のクローズに失敗しました", "error", err)
		return fmt.Errorf("データベース接続クローズ失敗: %w", err)
	}

	slog.Info("データベース接続をクローズしました")
	return nil
}

// RunMigrations はデータベースマイグレーションを実行
func RunMigrations(databaseURL string) error {
	slog.Info("マイグレーションを開始します")

	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)
	if err != nil {
		slog.Error("マイグレーションの初期化に失敗しました", "error", err)
		return fmt.Errorf("マイグレーション初期化失敗: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("マイグレーションに失敗しました", "error", err)
		return fmt.Errorf("マイグレーション実行失敗: %w", err)
	}

	slog.Info("マイグレーションが正常に終了しました")
	return nil
}
