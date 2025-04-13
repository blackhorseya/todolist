package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Config 代表應用程式設定
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig 代表 HTTP 伺服器設定
type ServerConfig struct {
	Port int `yaml:"port"`
}

// DatabaseConfig 代表資料庫設定
type DatabaseConfig struct {
	ConnString string `yaml:"dsn"`
}

// LoadEnv 載入環境變數
func LoadEnv(envFile string) error {
	if envFile == "" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("載入環境變數檔案錯誤: %w", err)
	}
	return nil
}

// LoadConfig 從 YAML 檔案載入設定
func LoadConfig(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("讀取設定檔錯誤: %w", err)
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("解析設定檔錯誤: %w", err)
	}

	// 優先從環境變數讀取資料庫連線字串
	if databaseDSN := os.Getenv("DATABASE_DSN"); databaseDSN != "" {
		config.Database.ConnString = databaseDSN
	}

	return config, nil
}

// GetDSN 取得資料庫連線字串
func (c *DatabaseConfig) GetDSN() string {
	return c.ConnString
}
