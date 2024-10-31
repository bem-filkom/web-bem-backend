package env

import (
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type EnvStruct struct {
	ENV                  string        `mapstructure:"ENV"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	AppPort              string        `mapstructure:"APP_PORT"`
	DBHost               string        `mapstructure:"DB_HOST"`
	DBPort               string        `mapstructure:"DB_PORT"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBPass               string        `mapstructure:"DB_PASS"`
	DBName               string        `mapstructure:"DB_NAME"`
	RedisAddr            string        `mapstructure:"REDIS_ADDR"`
	RedisPassword        string        `mapstructure:"REDIS_PASS"`
	JwtAccessSecretKey   []byte        `mapstructure:"-"`
	JwtAccessExpireTime  time.Duration `mapstructure:"JWT_ACCESS_EXPIRE_TIME"`
	JwtRefreshSecretKey  []byte        `mapstructure:"-"`
	JwtRefreshExpireTime time.Duration `mapstructure:"JWT_REFRESH_EXPIRE_TIME"`
}

var envObj *EnvStruct
var once sync.Once

func loadEnv() {
	fmt.Println("Creating env instance")

	envObj = &EnvStruct{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Fatal("[ENV][loadEnv] failed to read config file")
	}

	if err := viper.Unmarshal(envObj); err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Fatal("[ENV][loadEnv] failed to unmarshal to struct")
	}

	envObj.JwtAccessSecretKey = []byte(viper.Get("JWT_ACCESS_SECRET_KEY").(string))
	envObj.JwtRefreshSecretKey = []byte(viper.Get("JWT_REFRESH_SECRET_KEY").(string))

	if envObj.ENV != "development" && envObj.ENV != "staging" && envObj.ENV != "production" {
		log.GetLogger().Fatal("[ENV][loadEnv] ENV variable is undefined")
	}

	log.GetLogger().Info("[ENV][loadEnv] Application is running on " + envObj.ENV + " mode")
}

func SetEnv(env *EnvStruct) {
	envObj = env
}

func GetEnv() *EnvStruct {
	once.Do(loadEnv)
	return envObj
}
