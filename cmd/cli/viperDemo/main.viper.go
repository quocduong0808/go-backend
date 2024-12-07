package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User   string `mapstructure:"user"`
		Pass   string `mapstructure:"pass"`
		Host   string `mapstructure:"host"`
		DbName string `mapstructure:"dbName"`
		Schema string `mapstructure:"schema"`
	} `mapstructure:"databases"`

	Security struct {
		Jwt struct {
			Key string `mapstructure:"key"`
		}
	} `mapstructure:"security"`
}

func main() {
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
	var configOb Config
	parseErr := config.Unmarshal(&configOb)
	if parseErr != nil {
		panic(fmt.Errorf("error when parse config to struct %w", parseErr))
	}

	fmt.Println("Port::", configOb.Server.Port)
	for _, db := range configOb.Databases {
		fmt.Println("user::", db.User)
		fmt.Println("user::", db.Pass)
		fmt.Println("user::", db.Host)
		fmt.Println("user::", db.DbName)
		fmt.Println("user::", db.Schema)
	}
	fmt.Println("JWT::", configOb.Security.Jwt.Key)
}
