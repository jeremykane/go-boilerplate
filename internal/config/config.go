package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	// Config struct to wrap config file
	Config struct {
		Server   Server
		Worker   Worker
		Database map[string]*Database
		SchedulerConfig
	}

	Server struct {
		APIPort       int32
		GlobalTimeout int32
		ApiLogLevel   string
	}

	Worker struct {
		Port          int32
		GlobalTimeout int32
		LogLevel      string
	}

	SchedulerConfig struct {
		PlaceholderSchedulerTime string
	}
)

var ConfigMap *viper.Viper
var config Config

func Load(configPath, name string) (*Config, error) {
	v, err := loadFromEnvAndConfigEnv(configPath, name)
	if err != nil {
		return nil, err
	}

	ConfigMap = v

	dbMasterUser := getStringOrPanic("DB_USER")
	dbMasterPassword := getStringOrPanic("DB_PASSWORD")
	dbMasterName := getStringOrPanic("DB_NAME")
	dbMasterHost := getStringOrPanic("DB_HOST")
	dbMasterPort := getStringOrPanic("DB_PORT")
	dbMasterSslMode := getStringOrPanic("DB_SSL_MODE")

	dbMasterUrl := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbMasterUser,
		dbMasterPassword,
		dbMasterName,
		dbMasterHost,
		dbMasterPort,
		dbMasterSslMode,
	)

	databaseGoMaster := &DatabaseConfig{
		Url: dbMasterUrl,
	}

	var databaseGoSlave *DatabaseConfig
	if len(ConfigMap.GetString("DB_SLAVE_DRIVER")) > 0 {
		databaseGoSlave = &DatabaseConfig{
			Url: getStringOrPanic("DB_SLAVE_URL"),
		}
	}

	db := map[string]*Database{
		DatabaseGo: {
			Master: databaseGoMaster,
			Slave:  databaseGoSlave,
		},
	}

	serverHttp := Server{
		APIPort:       int32(getIntOrDefault("SERVER_API_PORT", 8080)),
		GlobalTimeout: int32(getIntOrDefault("SERVER_GLOBAL_TIMEOUT", 1000)),
		ApiLogLevel:   getStringOrDefault("SERVER_API_LOG_LEVEL", "info"),
	}

	workerHttp := Worker{
		Port:          int32(getIntOrDefault("WORKER_API_PORT", 8090)),
		GlobalTimeout: int32(getIntOrDefault("WORKER_GLOBAL_TIMEOUT", 1000)),
		LogLevel:      getStringOrDefault("WORKER_LOG_LEVEL", "info"),
	}

	schedulerConfig := SchedulerConfig{
		PlaceholderSchedulerTime: getStringOrPanic("PLACEHOLDER_SCHEDULER_TIME"),
	}

	config = Config{
		Database:        db,
		Server:          serverHttp,
		Worker:          workerHttp,
		SchedulerConfig: schedulerConfig,
	}

	return &config, nil
}

// loadFromEnvOrConfigEnv load config from config.env file that will overriden
// by env vars if the same key-value pair inside config file can also be found
// in env vars.
// The purpose of this function is to prioritize env vars value
// rather than using values from config files if the key-value pair
// exist
func loadFromEnvAndConfigEnv(path, name string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType("env")
	v.AddConfigPath(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func getStringOrDefault(key string, defaultVal string) string {
	v := ConfigMap.GetString(key)
	if v == "" {
		return defaultVal
	}

	return v
}

func getBoolOrDefault(key string, def bool) bool {
	if !ConfigMap.IsSet(key) {
		return def
	}

	return ConfigMap.GetBool(key)
}

func getStringOrPanic(key string) string {
	v := ConfigMap.GetString(key)
	if v == "" {
		panic("No value found for key: " + key)
	}

	return v
}

func getIntOrDefault(key string, defaultVal int) int {
	v := ConfigMap.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

func getIntOrPanic(key string) int {
	v := ConfigMap.GetInt(key)
	if v == 0 {
		panic("No value found for key: " + key)
	}

	return v
}
