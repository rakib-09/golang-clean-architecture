package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"time"
)

type AppConfig struct {
	Name         string
	Port         string
	AppKeyHeader string
	AppKey       string
	ImgBaseUrl   string
}

type DBConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	Schema          string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
}

type RedisConfig struct {
	Host               string
	Port               string
	Pass               string
	DB                 int
	MandatoryPrefix    string
	PartnerTokenPrefix string
	StoreDetailsPrefix string
	CacheMenuTTL       int // in second
}

type AsynqConfig struct {
	RedisAddr           string
	DB                  int
	Pass                string
	Concurrency         int
	Queue               string
	Retention           time.Duration // in hours
	RetryCount          int
	Delay               time.Duration // in seconds
	SyncOrderRetryDelay time.Duration // in seconds
	FullMenuRetryDelay  time.Duration // in seconds
}

type Config struct {
	App   *AppConfig
	Redis *RedisConfig
	Asynq *AsynqConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func Redis() *RedisConfig {
	return config.Redis
}

func Asynq() *AsynqConfig {
	return config.Asynq
}

func LoadConfig() {
	setDefaultConfig()

	_ = viper.BindEnv("CONSUL_URL")
	_ = viper.BindEnv("CONSUL_PATH")

	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")

	if consulURL != "" && consulPath != "" {
		_ = viper.AddRemoteProvider("consul", consulURL, consulPath)

		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()

		if err != nil {
			//logger.Error("%s named \"%s\"\n", err.Error(), consulPath)
			panic(err)
		}

		config = Config{}

		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}

		if r, err := json.MarshalIndent(&config, "", "  "); err == nil {
			fmt.Println(string(r))
		}
	} else {
		//logger.Info("CONSUL_URL or CONSUL_PATH missing! Serving with default config...")
	}

}

func setDefaultConfig() {
	config.App = &AppConfig{
		Name:         "github.com/rakib-09/golang-clean-architecture",
		Port:         "8080",
		AppKeyHeader: "github.com/rakib-09/golang-clean-architecture-app-key",
		AppKey:       "fb60e941d7c48db4810d2a4282732acf",
	}

	config.Redis = &RedisConfig{
		Host:            "redis",
		Port:            "6379",
		Pass:            "",
		DB:              6,
		MandatoryPrefix: "github.com/rakib-09/golang-clean-architecture_",
	}

	config.Asynq = &AsynqConfig{
		RedisAddr:   "redis:6379",
		DB:          15,
		Pass:        "",
		Concurrency: 10,
		Queue:       "github.com/rakib-09/golang-clean-architecture",
		Retention:   0,
		RetryCount:  0,
		Delay:       10,
	}
}
