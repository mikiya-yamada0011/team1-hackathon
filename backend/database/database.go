package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB はデータベースに接続（リトライ機能付き）
func ConnectDB(dsn string) (*gorm.DB, error) {
	const maxRetries = 30
	const retryInterval = 1 * time.Second

	var db *gorm.DB
	var err error

	for i := 1; i <= maxRetries; i++ {
		slog.Info("データベース接続を試行中...", "attempt", i, "max", maxRetries)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			slog.Warn("データベース接続に失敗しました。リトライします...", "error", err, "attempt", i)
			time.Sleep(retryInterval)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			slog.Warn("sql.DBの取得に失敗しました。リトライします...", "error", err, "attempt", i)
			time.Sleep(retryInterval)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := sqlDB.PingContext(ctx); err != nil {
			cancel()
			_ = sqlDB.Close()
			slog.Warn("データベースへのPingに失敗しました。リトライします...", "error", err, "attempt", i)
			time.Sleep(retryInterval)
			continue
		}
		cancel()

		slog.Info("データベースに接続しました")
		return db, nil
	}

	slog.Error("最大リトライ回数に到達しました。データベース接続に失敗しました", "error", err)
	return nil, fmt.Errorf("データベース接続失敗（%d回リトライ後）: %w", maxRetries, err)
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

// RunMigrations はデータベースマイグレーションを実行（リトライ機能付き）
func RunMigrations(databaseURL string) error {
	const maxRetries = 30
	const retryInterval = 1 * time.Second

	for i := 1; i <= maxRetries; i++ {
		slog.Info("マイグレーションを試行中...", "attempt", i, "max", maxRetries)

		m, err := migrate.New(
			"file://db/migrations",
			databaseURL,
		)
		if err != nil {
			slog.Warn("マイグレーションの初期化に失敗しました。リトライします...", "error", err, "attempt", i)
			time.Sleep(retryInterval)
			continue
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			slog.Warn("マイグレーションに失敗しました。リトライします...", "error", err, "attempt", i)
			time.Sleep(retryInterval)
			continue
		}

		slog.Info("マイグレーションが正常に終了しました")
		return nil
	}

	slog.Error("最大リトライ回数に到達しました。マイグレーションに失敗しました")
	return fmt.Errorf("マイグレーション実行失敗（%d回リトライ後）", maxRetries)
}

// SeedDatabase は開発環境用のテストデータを挿入
func SeedDatabase(db *gorm.DB) error {
	slog.Info("シードデータの挿入を開始します")

	// seed.sqlファイルを読み込む
	sqlBytes, err := os.ReadFile("db/seed.sql")
	if err != nil {
		slog.Warn("シードデータファイルが見つかりません", "error", err)
		return nil // エラーにせず警告のみ
	}

	// SQLを実行
	if err := db.Exec(string(sqlBytes)).Error; err != nil {
		slog.Error("シードデータの挿入に失敗しました", "error", err)
		return fmt.Errorf("シードデータ挿入失敗: %w", err)
	}

	slog.Info("シードデータの挿入が完了しました")
	return nil
}
