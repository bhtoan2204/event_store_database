package settings

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type SecurityConfig struct {
	JWTAccessSecret      string `mapstructure:"jwt_access_secret"`
	JWTRefreshSecret     string `mapstructure:"jwt_refresh_secret"`
	JWTAccessExpiration  string `mapstructure:"jwt_access_expiration"`
	JWTRefreshExpiration string `mapstructure:"jwt_refresh_expiration"`
	HMACSecret           string `mapstructure:"hmac_secret"`
}

type LogConfig struct {
	LogLevel   string `mapstructure:"log_level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type ConsulConfig struct {
	Address    string `mapstructure:"address"`
	Scheme     string `mapstructure:"scheme"`
	DataCenter string `mapstructure:"data_center"`
	Token      string `mapstructure:"token"`
}

type JaegerConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

type Service struct {
	PaymentServiceName string `mapstructure:"payment_service_name"`
	UserServiceName    string `mapstructure:"user_service_name"`
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
	Server         ServerConfig   `mapstructure:"server"`
	LogConfig      LogConfig      `mapstructure:"log"`
	SecurityConfig SecurityConfig `mapstructure:"security"`
	ConsulConfig   ConsulConfig   `mapstructure:"consul"`
	JaegerConfig   JaegerConfig   `mapstructure:"jaeger"`
	Service        Service        `mapstructure:"service"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
}
