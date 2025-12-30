package config

import (
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

// 設定を読み込むインターフェース
// このインターフェースを実装することで、異なる読み込み方法（YAML、JSON、DBなど）を切り替えられる
type ConfigLoader interface {
	LoadWithEnv(path string) (*Config, error)
}

type YAMLConfigLoader struct{}

func NewYAMLConfigLoader() ConfigLoader {
	return &YAMLConfigLoader{}
}

// YAMLファイルのみから設定を読み込む
func (l *YAMLConfigLoader) load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

// 環境変数による上書きを含めて設定を読み込む
func (l *YAMLConfigLoader) LoadWithEnv(path string) (*Config, error) {
	config := &Config{}

	// YAMLファイルが存在すれば読み込む
	if _, err := os.Stat(path); err == nil {
		config, err = l.load(path)
		if err == nil {
			slog.Debug("YAMLファイルを読み込みました", "path", path)
		} else {
			slog.Warn("YAMLファイルの読み込みに失敗しましたが、環境変数で設定します", "error", err, "path", path)
			config = &Config{}
		}
	} else {
		slog.Info("YAMLファイルが見つかりません。環境変数とデフォルト値で設定します", "path", path)
	}

	// 環境変数で設定・上書き
	if err := env.Parse(config); err != nil {
		slog.Error("環境変数のパースに失敗しました", "error", err)
		return nil, err
	}

	slog.Debug("環境変数による設定の上書きを完了しました")

	return config, nil
}
