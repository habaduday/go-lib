package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

func Load() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadInto[T any](cfg *T) (*T, error) {
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadFromEnvFile(path string) (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.UpdateEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadFromEnvFileInto[T any](path string, target *T) (*T, error) {
	if err := cleanenv.ReadConfig(path, target); err != nil {
		return nil, err
	}

	if err := cleanenv.UpdateEnv(target); err != nil {
		return nil, err
	}

	return target, nil
}
