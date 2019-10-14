package enviroment
import (
	"github.com/spf13/viper"
	
)

func ReadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile("./enviroment/env.json")
	v.AddConfigPath("./src/enviroment")
	viper.AddConfigPath("$HOME/enviroment")
	v.SetConfigName("env")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	return v, err
}