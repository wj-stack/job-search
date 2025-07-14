package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config 定义项目配置结构
type Config struct {
	Database struct {
		Driver string `yaml:"driver"`
		Source string `yaml:"source"`
	} `yaml:"database"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

// LoadConfig 从文件加载配置
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
