package config

import (
	"event_sourcing_gateway/package/settings"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitLoadConfig() (*settings.Config, error) {
	if fallbackErr := godotenv.Load(".env"); fallbackErr != nil {
		panic(fmt.Errorf("error loading .env files: %w", fallbackErr))
	}

	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnv(v)

	var config settings.Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %w", err))
	}

	return &config, nil
}

func bindEnv(v *viper.Viper) {
	// Set up mappings for environment variables to configuration structure
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("server.mode", "SERVER_MODE")

	// Log mappings
	v.BindEnv("log.log_level", "LOG_LOG_LEVEL")
	v.BindEnv("log.max_size", "LOG_MAX_SIZE")
	v.BindEnv("log.max_backups", "LOG_MAX_BACKUPS")
	v.BindEnv("log.max_age", "LOG_MAX_AGE")
	v.BindEnv("log.compress", "LOG_COMPRESS")

	// Consul mappings
	v.BindEnv("consul.address", "CONSUL_ADDRESS")
	v.BindEnv("consul.scheme", "CONSUL_SCHEME")
	v.BindEnv("consul.data_center", "CONSUL_DATA_CENTER")
	v.BindEnv("consul.token", "CONSUL_TOKEN")

	// Security mappings
	v.BindEnv("security.jwt_access_secret", "SECURITY_JWT_ACCESS_SECRET")
	v.BindEnv("security.jwt_refresh_secret", "SECURITY_JWT_REFRESH_SECRET")
	v.BindEnv("security.jwt_access_expiration", "SECURITY_JWT_ACCESS_EXPIRATION")
	v.BindEnv("security.jwt_refresh_expiration", "SECURITY_JWT_REFRESH_EXPIRATION")
	v.BindEnv("security.hmac_secret", "SECURITY_HMAC_SECRET")

	// Jaeger mappings
	v.BindEnv("jaeger.endpoint", "JAEGER_ENDPOINT")

	// Service mappings
	v.BindEnv("service.payment_service_name", "PAYMENT_SERVICE_NAME")
	v.BindEnv("service.user_service_name", "USER_SERVICE_NAME")

	// Redis mappings
	v.BindEnv("redis.connection_url", "REDIS_CONNECTION_URL")
	v.BindEnv("redis.db", "REDIS_DB")
	v.BindEnv("redis.pool_size", "REDIS_POOL_SIZE")
	v.BindEnv("redis.dial_timeout_seconds", "REDIS_DIAL_TIMEOUT_SECONDS")
	v.BindEnv("redis.read_timeout_seconds", "REDIS_READ_TIMEOUT_SECONDS")
	v.BindEnv("redis.write_timeout_seconds", "REDIS_WRITE_TIMEOUT_SECONDS")
	v.BindEnv("redis.idle_timeout_seconds", "REDIS_IDLE_TIMEOUT_SECONDS")
	v.BindEnv("redis.max_idle_conn_number", "REDIS_MAX_IDLE_CONN_NUMBER")
	v.BindEnv("redis.max_active_conn_number", "REDIS_MAX_ACTIVE_CONN_NUMBER")
}
