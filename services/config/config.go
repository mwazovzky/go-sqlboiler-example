package config

import (
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() Config {
	return Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func (cfg Config) FormatDSN() string {
	c := mysql.NewConfig()
	c.Net = "tcp"
	c.Addr = net.JoinHostPort(cfg.Host, cfg.Port)
	c.User = cfg.User
	c.Passwd = cfg.Password
	c.DBName = cfg.Database
	c.ParseTime = true
	c.MultiStatements = true

	return c.FormatDSN()
}
