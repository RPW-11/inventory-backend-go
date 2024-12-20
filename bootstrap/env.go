package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	CorsOrigin             string `mapstructure:"CORS_ORIGIN"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBSSLMode              string `mapstructure:"DB_SSLMODE"`
	DBTimeZone             string `mapstructure:"DB_TIMEZONE"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AWS_Region             string `mapstructure:"AWS_REGION"`
	S3Bucket               string `mapstructure:"S3_BUCKET"`
	AWSAccessKeyID         string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey     string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the .env file: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environmetn can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The app is running in development environment")
	}

	return &env
}
