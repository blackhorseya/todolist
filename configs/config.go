package configs

import (
	"fmt"
	"os"

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
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
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

	return config, nil
}

// DSN 取得資料庫連線字串
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
