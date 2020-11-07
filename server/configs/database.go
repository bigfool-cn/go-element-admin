package configs

import "github.com/spf13/viper"

type Database struct {
	Database string
	Dbtype   string
	Host     string
	Password string
	Port     int
	Username string
}

func InitDatabase(cfg *viper.Viper) *Database  {
	return &Database{
		Database:      cfg.GetString("database"),
		Dbtype:        cfg.GetString("dbtype"),
		Host:          cfg.GetString("host"),
		Port:          cfg.GetInt("port"),
		Password:      cfg.GetString("password"),
		Username:      cfg.GetString("username"),
	}
}

var DatabaseConfig = new(Database)