package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func ParseConfig() {
	fmt.Printf("Config File = %s\n", Config)
	viper.SetConfigFile(Config)
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error-%s", err.Error())
		os.Exit(2)
	}
	fmt.Printf("Parser Done\n")
	User = viper.GetString("USER")
	App.UserName = viper.GetString("USER")
	App.Id = viper.GetInt("ID")
	if App.Id <= 0 {
		log.Printf("Invalid value ID - %v", App.Id)
		os.Exit(2)
	}
}
