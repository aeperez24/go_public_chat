package configuration

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfiguration() {
	log.Printf("reading configuration")
	viper.AddConfigPath("../")

	viper.SetConfigFile("../configFiles.json")

	viper.AutomaticEnv()
	error := viper.ReadInConfig()

	if error != nil {
		panic(error)
	}
	value := viper.Get("port")
	ReadQuerysFile()
	log.Printf("el puerto leido de viper es %v", value)

}

func ReadQuerysFile() {
	log.Printf("reading queries")

	viper.SetConfigFile("../queryFiles.json")
	viper.AddConfigPath("../")
	viper.SetConfigType("json")
	error := viper.MergeInConfig()

	if error != nil {
		panic(error)
	}

}
