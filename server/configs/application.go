package configs

import "github.com/spf13/viper"

type Application struct {
	IsInit        bool
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	LogPath       string
	Env           string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		IsInit:        cfg.GetBool("isInit"),
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		Name:          cfg.GetString("name"),
		JwtSecret:     cfg.GetString("jwtSecret"),
		LogPath:       cfg.GetString("logPath"),
		Env:           cfg.GetString("env"),
	}
}

var ApplicationConfig = new(Application)
