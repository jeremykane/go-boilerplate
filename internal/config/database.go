package config

const DatabaseGo = "db"

type Database struct {
	Master *DatabaseConfig
	Slave  *DatabaseConfig
}

type DatabaseConfig struct {
	Url   string
	Retry int
}
