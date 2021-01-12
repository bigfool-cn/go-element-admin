package configs

import "github.com/spf13/viper"

type Application struct {
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	LogPath       string
  Debug         bool
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		Name:          cfg.GetString("name"),
		JwtSecret:     cfg.GetString("jwtSecret"),
		LogPath:       cfg.GetString("logPath"),
    Debug:         cfg.GetBool("debug"),
	}
}

var ApplicationConfig = new(Application)
