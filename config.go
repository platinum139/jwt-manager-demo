package main

import (
    "github.com/spf13/viper"
    "sync"
)

type Config struct {
    JwtSecretKey    string `mapstructure:"JWT_SECRET_KEY"`
    AccessTokenMin  int    `mapstructure:"ACCESS_TOKEN_MIN"`
    RefreshTokenMin int    `mapstructure:"REFRESH_TOKEN_MIN"`
    RedisHost       string `mapstructure:"REDIS_HOST"`
    RedisPort       string `mapstructure:"REDIS_PORT"`
    RedisPassword   string `mapstructure:"REDIS_PASSWORD"`
}

func (config *Config) Load(path, name string) (err error) {
    once := sync.Once{}
    once.Do(func() {
        viper.AddConfigPath(path)
        viper.SetConfigName(name)
        viper.SetConfigType("env")
        err = viper.ReadInConfig()
    })
    if err != nil {
        viper.AutomaticEnv()
        viper.BindEnv("DB_USERNAME")
        viper.BindEnv("DB_PASSWORD")
        viper.BindEnv("DB_HOST")
        viper.BindEnv("DB_PORT")
        viper.BindEnv("DB_NAME")
        return viper.Unmarshal(config)
    }
    return viper.Unmarshal(config)
}

func NewAppConfig() *Config {
    return &Config{}
}
