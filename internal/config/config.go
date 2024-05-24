package config

import "time"

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"default" env-required:"true"`
	Storage    string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}
