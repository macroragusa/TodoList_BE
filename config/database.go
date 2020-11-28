package config

import "fmt"

type DataBaseConnConf struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SslMode  string
	TimeZone string
}

func dbConfig() string {
	dbConnConf := DataBaseConnConf{
		"localhost",
		"5432",
		"todo",
		"postgres",
		"1234",
		"disable",
		"Europe/Rome",
	}

	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s TimeZone=%s",
		dbConnConf.Host,
		dbConnConf.Port,
		dbConnConf.DBName,
		dbConnConf.User,
		dbConnConf.Password,
		dbConnConf.SslMode,
		dbConnConf.TimeZone,
	)
}
