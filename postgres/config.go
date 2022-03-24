package postgres

import "fmt"

type Config struct {
	User     string
	Password string
	Port     string
	Database string
	Host     string
	Ssl      string
	Timezone string
}

func GetDSN(config Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.Database, config.Port, config.Ssl, config.Timezone)
}
