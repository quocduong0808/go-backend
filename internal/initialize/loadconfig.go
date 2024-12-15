package initialize

import (
	"fmt"
	"go/go-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	config := viper.New()
	config.AddConfigPath("./config/")
	config.SetConfigName("dev")
	config.SetConfigType("yaml")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot read configuration %w", err))
	}

	fmt.Println("Port: ", config.GetInt("server.port"))

	//parse to config object
	parseErr := config.Unmarshal(&global.Config)
	if parseErr != nil {
		panic(fmt.Errorf("error when parse config to struct %w", parseErr))
	}
}
