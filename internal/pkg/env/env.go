package env

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/spf13/viper"
)

type EnvStruct struct {
	ENV                 string
	LogLevel            string
	AppPort             string
	DBHost              string
	DBPort              string
	DBUser              string
	DBPass              string
	DBName              string
	JwtAccessSecretKey  []byte
	JwtAccessExpireTime time.Duration
}

var envObj *EnvStruct
var once sync.Once

func loadEnv() {
	fmt.Println("Loading environment configuration")

	envObj = &EnvStruct{}

	// Initialize viper for .env loading
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("[ENV] .env file not found, falling back to system environment variables")
	}

	envObj.ENV = getRequiredConfig("ENV")
	envObj.LogLevel = getRequiredConfig("LOG_LEVEL")
	envObj.AppPort = getRequiredConfig("APP_PORT")
	envObj.DBHost = getRequiredConfig("DB_HOST")
	envObj.DBPort = getRequiredConfig("DB_PORT")
	envObj.DBUser = getRequiredConfig("DB_USER")
	envObj.DBPass = getRequiredConfig("DB_PASS")
	envObj.DBName = getRequiredConfig("DB_NAME")
	envObj.JwtAccessSecretKey = []byte(getRequiredConfig("JWT_ACCESS_SECRET_KEY"))

	jwtAccessExpireTime, err := time.ParseDuration(getRequiredConfig("JWT_ACCESS_EXPIRE_TIME"))
	envObj.JwtAccessExpireTime = jwtAccessExpireTime

	if envObj.ENV != "development" && envObj.ENV != "staging" && envObj.ENV != "production" {
		log.GetLogger().Fatal("[ENV][loadEnv] ENV variable must be one of: development, staging, production")
	}

	log.GetLogger().Info("[ENV][loadEnv] Application is running on " + envObj.ENV + " mode")
}

func getRequiredConfig(key string) string {
	// Check if the key exists in .env or environment variables
	if value := viper.GetString(key); value != "" {
		return value
	}

	// Fallback to system environment variables
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	// If not found in both, log fatal error
	log.GetLogger().Fatal("[ENV][getRequiredConfig] Missing required environment variable: " + key)
	return ""
}

func SetEnv(env *EnvStruct) {
	envObj = env
}

func GetEnv() *EnvStruct {
	once.Do(loadEnv)
	return envObj
}
