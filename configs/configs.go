package configs

var Config Configuration

type ServerConfiguration struct {
	Host     string `mapstructure:"host"`
	GRPCPort string `mapstructure:"grpcport"`
	HTTPPort string `mapstructure:"httpport"`
}

type DatabaseConfiguration struct {
	Host                  string `mapstructure:"host"`
	Port                  int    `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	DB                    string `mapstructure:"db"`
	ConnectionMaxLifetime int    `mapstructure:"connectionMaxLifetime"`
	MaxIdleConnections    int    `mapstructure:"MaxIdleConnections"`
	MaxOpenConnections    int    `mapstructure:"MaxOpenConnections"`
}

type CacheConfiguration struct {
	Client   string `mapstructure:"client"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type Credential struct {
	TokenSecret string `mapstructure:"tokensecret"`
}

type RateLimitConfiguration struct {
	RPCName           string `mapstructure:"rpcName"`
	RequestsPerSecond int    `mapstructure:"requestsPerSecond"`
}

type Configuration struct {
	Server     ServerConfiguration      `mapstructure:"server"`
	Database   DatabaseConfiguration    `mapstructure:"database"`
	Cache      CacheConfiguration       `mapstructure:"cache"`
	Credential Credential               `mapstructure:"credential"`
	RateLimits []RateLimitConfiguration `mapstructure:"ratelimit"`
}
