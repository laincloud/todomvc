package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// StaticFileDirectory 表示静态文件目录
const StaticFileDirectory = "/lain/app/dist"

// Config 表示配置
type Config struct {
	MySQL MySQLConfig
}

// New 返回初始化后的 Config
func New(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var c Config
	if err = json.NewDecoder(file).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// MySQLConfig 表示 MySQL 配置
type MySQLConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

// DataSourceName 返回 MySQL 的连接字符串
func (m MySQLConfig) DataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}
