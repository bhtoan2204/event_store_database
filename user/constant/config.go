package constant

type ServerConfig struct {
	ServerMode    string `mapstructure:"server_mode"`
	ServerGinMode string `mapstructure:"server_gin_mode"`
	GRPCPort      int    `mapstructure:"grpc_port"`
	HTTPPort      int    `mapstructure:"http_port"`
	MetricsPort   int    `mapstructure:"metrics_port"`
}

type SecurityConfig struct {
	JWTAccessSecret      string `mapstructure:"jwt_access_secret"`
	JWTRefreshSecret     string `mapstructure:"jwt_refresh_secret"`
	JWTAccessExpiration  int    `mapstructure:"jwt_access_expiration"`
	JWTRefreshExpiration int    `mapstructure:"jwt_refresh_expiration"`
}

type PostgresConfig struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Database     string `mapstructure:"database"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Security SecurityConfig `mapstructure:"security"`
}
