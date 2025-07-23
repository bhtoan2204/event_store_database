package config

import (
	"event_sourcing_user/constant"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() (*constant.Config, error) {
	var config constant.Config

	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnv(v)

	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %w", err))
	}

	return &config, nil
}

func bindEnv(v *viper.Viper) {
	// Server mappings
	v.BindEnv("server.server_mode", "SERVER_MODE")
	v.BindEnv("server.server_gin_mode", "SERVER_GIN_MODE")
	v.BindEnv("server.service_name", "SERVICE_NAME")

	// Security mappings
	v.BindEnv("security.jwt_access_secret", "SECURITY_JWT_ACCESS_SECRET")
	v.BindEnv("security.jwt_refresh_secret", "SECURITY_JWT_REFRESH_SECRET")
	v.BindEnv("security.jwt_access_expiration", "SECURITY_JWT_ACCESS_EXPIRATION")
	v.BindEnv("security.jwt_refresh_expiration", "SECURITY_JWT_REFRESH_EXPIRATION")

	// Log mappings
	v.BindEnv("log.log_level", "LOG_LOG_LEVEL")
	v.BindEnv("log.max_size", "LOG_MAX_SIZE")
	v.BindEnv("log.max_backups", "LOG_MAX_BACKUPS")
	v.BindEnv("log.max_age", "LOG_MAX_AGE")
	v.BindEnv("log.compress", "LOG_COMPRESS")

	// Postgres mappings
	v.BindEnv("postgres.username", "POSTGRES_USERNAME")
	v.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	v.BindEnv("postgres.host", "POSTGRES_HOST")
	v.BindEnv("postgres.port", "POSTGRES_PORT")
	v.BindEnv("postgres.database", "POSTGRES_DATABASE")
	v.BindEnv("postgres.max_idle_conns", "POSTGRES_MAX_IDLE_CONNS")
	v.BindEnv("postgres.max_open_conns", "POSTGRES_MAX_OPEN_CONNS")
	v.BindEnv("postgres.max_lifetime", "POSTGRES_MAX_LIFETIME")

	// Consul mappings
	v.BindEnv("consul.address", "CONSUL_ADDRESS")
	v.BindEnv("consul.scheme", "CONSUL_SCHEME")
	v.BindEnv("consul.data_center", "CONSUL_DATA_CENTER")
	v.BindEnv("consul.token", "CONSUL_TOKEN")
}
