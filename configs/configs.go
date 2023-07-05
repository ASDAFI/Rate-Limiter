package configs

var Config Configuration

type ServerConfiguration struct {
	Host     string
	GRPCPort string
	HTTPPort string
}

type DatabaseConfiguration struct {
	Host                  string
	Port                  int
	User                  string
	Password              string
	DB                    string
	ConnectionMaxLifetime int
	MaxIdleConnections    int
	MaxOpenConnections    int
}

type Credential struct {
	TokenSecret string
}

type RateLimitConfiguration struct {
	Name              string
	RequestsPerSecond int
}

type Configuration struct {
	Server     ServerConfiguration
	Database   DatabaseConfiguration
	Credential Credential
	RateLimits []RateLimitConfiguration
}
