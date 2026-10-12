package core

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type apiKeys struct {
	abacatePayAPIKey string
}

type EnvConfig struct {
	apiKeys apiKeys
}

var envConfigData EnvConfig

func GetEnvConfig() EnvConfig {
	return envConfigData
}

func LoadEnvConfig() error {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return err
	}

	keys, err := loadAPIKeysConfig()
	if err != nil {
		return err
	}

	envConfigData = EnvConfig{
		apiKeys: keys,
	}

	return nil
}

// MARK: - Load env

func loadAPIKeysConfig() (apiKeys, error) {
	abacatePayAPIKey := os.Getenv("ABACATEPAY_API_KEY")
	if abacatePayAPIKey == "" {
		return apiKeys{}, errors.New("ABACATEPAY_API_KEY is required.")
	}

	return apiKeys{
		abacatePayAPIKey: abacatePayAPIKey,
	}, nil
}

// MARK: - Read envs

func (ec EnvConfig) AbacatePayAPIKey() string {
	return ec.apiKeys.abacatePayAPIKey
}
