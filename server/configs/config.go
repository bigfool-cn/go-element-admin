package configs

import (
	"github.com/spf13/viper"
	"log"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper

func init()  {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("./configs") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {
		log.Println(err) // Handle errors reading the config file
	}

	cfgDatabase = viper.Sub("database")
	if cfgDatabase == nil {
		panic("config not found database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("application")
	if cfgApplication == nil {
		panic("config not found application")
	}
	ApplicationConfig = InitApplication(cfgApplication)
}
