package constant

type ServerConfig struct {
	ServerMode    string `mapstructure:"server_mode"`
	ServerGinMode string `mapstructure:"server_gin_mode"`
	ServiceName   string `mapstructure:"service_name"`
	// GRPCPort      int    `mapstructure:"grpc_port"`
	// HTTPPort      int    `mapstructure:"http_port"`
	// MetricsPort   int    `mapstructure:"metrics_port"`
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

type ConsulConfig struct {
	Address    string `mapstructure:"address"`
	Scheme     string `mapstructure:"scheme"`
	DataCenter string `mapstructure:"data_center"`
	Token      string `mapstructure:"token"`
}

type LogConfig struct {
	LogLevel   string `mapstructure:"log_level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type EventStoreConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RedisConfig struct {
	ConnectionURL       string `mapstructure:"connection_url"`
	Password            string `mapstructure:"password"`
	DB                  int    `mapstructure:"db"`
	PoolSize            int    `mapstructure:"pool_size"`
	DialTimeoutSeconds  int    `mapstructure:"dial_timeout_seconds"`
	ReadTimeoutSeconds  int    `mapstructure:"read_timeout_seconds"`
	WriteTimeoutSeconds int    `mapstructure:"write_timeout_seconds"`
	IdleTimeoutSeconds  int    `mapstructure:"idle_timeout_seconds"`
	MaxIdleConn         int    `mapstructure:"max_idle_conn_number"`
	MaxActiveConn       int    `mapstructure:"max_active_conn_number"`
}

type Config struct {
	Server       ServerConfig     `mapstructure:"server"`
	Postgres     PostgresConfig   `mapstructure:"postgres"`
	ConsulConfig ConsulConfig     `mapstructure:"consul"`
	EventStore   EventStoreConfig `mapstructure:"event_store"`
	Redis        RedisConfig      `mapstructure:"redis"`
}
