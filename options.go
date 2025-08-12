package taos_native

type Config struct {
	Host      string
	Port      int
	UserName  string
	Password  string
	Database  string
	Precision string
}

type Option func(*Config)

func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func WithUserName(userName string) Option {
	return func(c *Config) {
		c.UserName = userName
	}
}

func WithPassword(password string) Option {
	return func(c *Config) {
		c.Password = password
	}
}

func WithDatabase(database string) Option {
	return func(c *Config) {
		c.Database = database
	}
}

func WithPrecision(precision string) Option {
	return func(c *Config) {
		c.Precision = precision
	}
}

func NewConfig(opts ...Option) Config {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}
