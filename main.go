package main

import (
    "context"
    jwt "github.com/platinum139/jwt-manager"
    rds "github.com/platinum139/jwt-manager/redis"
    "log"
    "os"
)

func main() {
    ctx := context.Background()
    logger := log.New(os.Stdout, "[main]", log.Ldate)

    appConfig := NewAppConfig()
    if err := appConfig.Load(".", ".env"); err != nil {
        logger.Printf("cannot load app config: %s\n", err)
        return
    }

    jwtConfig := jwt.JwtConfig{
        SecretKey:       appConfig.JwtSecretKey,
        AccessTokenMin:  appConfig.AccessTokenMin,
        RefreshTokenMin: appConfig.RefreshTokenMin,
    }
    redisConfig := rds.RedisConfig{
        Host:     appConfig.RedisHost,
        Port:     appConfig.RedisPort,
        Password: appConfig.RedisPassword,
    }
    config := jwt.Config{
        Jwt:   jwtConfig,
        Redis: redisConfig,
    }

    userId := "platinum139"
    jwtManager := jwt.NewJwtManager(ctx, logger, config)

    accessToken, err := jwtManager.GenerateAccessToken(userId)
    if err != nil {
        log.Printf("unable to generate access token: %s\n", err)
    } else {
        log.Printf("access token is generated successfully: %s\n", accessToken)
    }

    refreshToken, err := jwtManager.GenerateRefreshToken()
    if err != nil {
        log.Printf("unable to generate refresh token: %s\n", err)
    } else {
        log.Printf("refresh token is generated successfully: %s\n", refreshToken)
    }

    err = jwtManager.SaveRefreshToken(userId, refreshToken)
    if err != nil {
        log.Printf("unable to save refresh token: %s\n", err)
    } else {
        log.Printf("refresh token is saved successfully: %s\n", refreshToken)
    }

    gotUserId, err := jwtManager.ValidateAccessToken(accessToken)
    if err != nil {
        log.Printf("unable to validate access token: %s\n", err)
    } else {
        log.Printf("access token is valid. userId=%s\n", gotUserId)
    }

    isValid, err := jwtManager.ValidateRefreshToken(userId, refreshToken)
    if err != nil {
        log.Printf("unable to validate refresh token: %s\n", err)
    } else {
        log.Printf("refresh token is valid: %t\n", isValid)
    }

    err = jwtManager.DeleteRefreshToken(userId, refreshToken)
    if err != nil {
        log.Printf("unable to delete refresh token: %s\n", err)
    } else {
        log.Printf("refresh token is deleted successfully")
    }
}
