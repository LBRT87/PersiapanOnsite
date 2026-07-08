package config

import (
	"fmt"
	"os"
	"strconv"
)

type CourseConfig struct {
	Port   int64
	DbPort int64
	DbHost string
	DbName string
	DbUser string
	DbPass string
}

func Load() (*CourseConfig, error) {
	port, err := strconv.Atoi(lookupOrFallbackEnv("PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing : %v", err)
	}

	dbPort, err := strconv.Atoi(lookupOrFallbackEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing : %v", err)
	}

	return &CourseConfig{
		Port: int64(port),
		DbPort: int64(dbPort),
		DbName: lookupOrFallbackEnv("DB_NAME", "esdemy-course"),
		DbHost: lookupOrFallbackEnv("DB_HOST", "pgpool"),
		DbPass: lookupOrFallbackEnv("DB_PASS", "postgres"),
		DbUser: lookupOrFallbackEnv("DB_USER", "postgres"),
	}, nil
}

func lookupOrFallbackEnv(target, fallback string) string {
	val, ok := os.LookupEnv(target)

	if ok {
		return val
	}

	return fallback
}