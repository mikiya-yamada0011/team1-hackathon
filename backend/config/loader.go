package config

import (
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

// 設定を読み込むインターフェース
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
	var config *Config

	// YAMLファイルが存在すれば読み込む（オプション）
	if _, err := os.Stat(path); err == nil {
		yamlConfig, err := l.load(path)
		if err == nil {
			config = yamlConfig
			slog.Debug("YAMLファイルを読み込みました", "path", path)
		} else {
			slog.Warn("YAMLファイルの読み込みに失敗しました", "error", err, "path", path)
			config = &Config{}
		}
	} else {
		slog.Info("YAMLファイルが見つかりません。環境変数で設定します", "path", path)
		config = &Config{}
	}

	// 環境変数で上書き（環境変数が設定されている項目のみ）
	if err := env.Parse(config); err != nil {
		slog.Error("環境変数のパースに失敗しました", "error", err)
		return nil, err
	}

	slog.Debug("設定の読み込みが完了しました")

	return config, nil
}
