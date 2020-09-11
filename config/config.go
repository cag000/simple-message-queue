package config

import "os"

type DbConfig struct {
	DbUrl      string
	DbPort     string
	DbUserName string
	DbPassword string
	DbName     string
}

type QueueConfig struct {
	Address string
}

type User struct {
	Username string
	Password string
}
type Config struct {
	DbConfigClient DbConfig
	Mqueue         QueueConfig
	UserApp        User
}

func New() *Config {
	return &Config{
		DbConfigClient: DbConfig{
			DbName:     getEnv("DB_NUMBER", ""),
			DbPort:     getEnv("DB_PORT", ""),
			DbPassword: getEnv("DB_PASSWORD", ""),
			DbUrl:      getEnv("DB_URL", ""),
			DbUserName: getEnv("DB_USERNAME", ""),
		},
		Mqueue: QueueConfig{
			Address: getEnv("QUEUE_ADDRESS", ""),
		},
		UserApp: User{
			Username: getEnv("USER_NAME", ""),
			Password: getEnv("USER_PASSWORD", ""),
		},
	}
}

func getEnv(k, defVal string) string {
	if value, exists := os.LookupEnv(k); exists {
		return value
	}
	return defVal
}
