package Infrastructure

import "github.com/spf13/viper"

func ReadConfig() error {
	viper.SetConfigFile(".env")

	return viper.ReadInConfig()
}
