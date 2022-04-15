package config

type Host struct {
	Port string `env:"HOST_PORT"`
	Key  string `env:"HOST_KEY_PATH"`  // Path to TLS key
	Cert string `env:"HOST_CERT_PATH"` // Path to TLS key
}

type Postgres struct {
	User     string `env:"POSTGRES_USER"`
	Pass     string `env:"POSTGRES_PASSWORD"`
	DbName   string `env:"POSTGRES_DATABASE"`
	IP       string `env:"POSTGRES_IP"`
	Port     string `env:"POSTGRES_PORT"`
	Protocol string `env:"POSTGRES_PROTOCOL"`
}

type Memcached struct {
	IP   string `env:"MEMCACHED_IP"`
	Port string `env:"MEMCACHED_PORT"`
}

type Config struct {
	Memcached Memcached
	Postgres  Postgres
	Host      Host
}
