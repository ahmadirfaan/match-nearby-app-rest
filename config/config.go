package config

import (
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"

	"os"

	"github.com/joho/godotenv" // for .env file support (optional)
)

type Config struct {
	AppName                  string
	AppPort                  string
	AppTimeout               int
	LogLevel                 string
	Environment              string
	JWTSecret                string
	TokenTTL                 int
	RedisAddress             string
	DBUsername               string
	DBPassword               string
	DBHost                   string
	DBPort                   int
	DBName                   string
	DBMaxConnections         int
	DBMaxIdleConnections     int
	DBMaxLifetimeConnections int
	MaximalTotalSwipe        int
}

var cachedConfig *Config
var configOnce sync.Once

func Init() *Config {

	if cachedConfig != nil {
		return cachedConfig
	}

	// Load .env variables (optional)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	dbMaxConnections, err := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_CONNECTIONS: %v", err)
	}

	dbMaxIdleConnections, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_IDLE_CONNECTIONS: %v", err)
	}

	dbMaxLifetimeConnections, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_LIFETIME_CONNECTIONS: %v", err)
	}

	timeout, err := strconv.Atoi(os.Getenv("APP_TIMEOUT"))
	if err != nil {
		log.Warnf("Invalid APP_TIMEOUT: %v", err)
		timeout = 5
	}

	tokenTTL, err := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	if err != nil {
		log.Warnf("Invalid TOKEN_TTL: %v", err)
		tokenTTL = 604800
	}

	maximalTotalSwipe, err := strconv.Atoi(os.Getenv("MAXIMAL_TOTAL_SWIPE"))
	if err != nil {
		log.Warnf("Invalid MAXIMAL_TOTAL_SWIPE: %v", err)
		maximalTotalSwipe = 10
	}

	os.Getenv("MAXIMAL_TOTAL_SWIPE")

	config := &Config{
		AppPort:                  os.Getenv("APP_PORT"),
		AppName:                  os.Getenv("APP_NAME"),
		AppTimeout:               timeout,
		LogLevel:                 os.Getenv("LOG_LEVEL"),
		Environment:              os.Getenv("ENV_MODE"),
		JWTSecret:                os.Getenv("JWT_SECRET"),
		TokenTTL:                 tokenTTL,
		RedisAddress:             os.Getenv("REDIS_ADDRESS"),
		DBUsername:               os.Getenv("DB_USERNAME"),
		DBPassword:               os.Getenv("DB_PASSWORD"),
		DBHost:                   os.Getenv("DB_HOST"),
		DBPort:                   dbPort,
		DBName:                   os.Getenv("DB_NAME"),
		DBMaxConnections:         dbMaxConnections,
		DBMaxIdleConnections:     dbMaxIdleConnections,
		DBMaxLifetimeConnections: dbMaxLifetimeConnections,
		MaximalTotalSwipe:        maximalTotalSwipe,
	}

	configOnce.Do(func() {
		// Only load the config once
		cachedConfig = config
	})

	return config

}
