package configs

import (
	"fmt"
	"os"
	"strconv"

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

// GetDSN 取得資料庫連線字串
func (c *DatabaseConfig) GetDSN() string {
	return c.ConnString
}

// LoadEnv 載入環境變數並回傳設定
func LoadEnv(configPath string) (*Config, error) {
	// 載入環境變數
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("載入環境變數檔案錯誤: %w", err)
	}

	var config *Config

	// 如果有提供設定檔路徑，就讀取設定檔
	if configPath != "" {
		buf, err := os.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("讀取設定檔錯誤: %w", err)
		}

		config = &Config{}
		if err = yaml.Unmarshal(buf, config); err != nil {
			return nil, fmt.Errorf("解析設定檔錯誤: %w", err)
		}
	} else {
		// 如果沒有提供設定檔，就建立預設設定
		config = &Config{
			Server: ServerConfig{
				Port: 8080, // 預設連接埠
			},
		}
	}

	// 優先使用環境變數覆蓋設定
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Server.Port = p
		}
	}

	if databaseDSN := os.Getenv("DATABASE_DSN"); databaseDSN != "" {
		config.Database.ConnString = databaseDSN
	}

	return config, nil
}
