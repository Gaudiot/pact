package core

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type apiKeys struct {
	abacatePayAPIKey string
	resendAPIKey     string
	clerkSecretKey   string
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

	resendAPIKey := os.Getenv("RESEND_API_KEY")
	if resendAPIKey == "" {
		return apiKeys{}, errors.New("RESEND_API_KEY is required.")
	}

	clerkSecretKey := os.Getenv("RESEND_API_KEY")
	if clerkSecretKey == "" {
		return apiKeys{}, errors.New("RESEND_API_KEY is required.")
	}

	return apiKeys{
		abacatePayAPIKey: abacatePayAPIKey,
		resendAPIKey:     resendAPIKey,
		clerkSecretKey:   clerkSecretKey,
	}, nil
}

// MARK: - Read envs

func (ec EnvConfig) AbacatePayAPIKey() string {
	return ec.apiKeys.abacatePayAPIKey
}

func (ec EnvConfig) ResendAPIKey() string {
	return ec.apiKeys.resendAPIKey
}

func (ec EnvConfig) ClerkSecretKey() string {
	return ec.apiKeys.clerkSecretKey
}
