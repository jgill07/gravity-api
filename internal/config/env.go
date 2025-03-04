package config

import (
	"os"
	"strconv"

	"github.com/jgill07/gravity-api/internal/log"
	"go.uber.org/zap"
)

func getEnv(key string) string {
	return os.Getenv(key)
}

func getIntEnv(key string, fallback int) int {
	val := getEnv(key)
	if val == "" {
		return fallback
	}

	cvt, err := strconv.Atoi(val)
	if err != nil {
		log.WithFields(zap.String(key, val)).Debug("Error converting env to int")
		return fallback
	}
	return cvt
}
