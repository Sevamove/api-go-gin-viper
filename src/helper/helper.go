package helper

import (
	"errors"
	"fmt"
	"log"
	"time"

	"api/src/db"

	"github.com/spf13/viper"
)

// Set config file using viper.
func SetConfig(filePath string, fileName string, configType string) {
	// Directory name where the config.yaml lives.
	viper.AddConfigPath(filePath)
	// File name of the config.yaml
	viper.SetConfigName(fileName)
	//OPTIONAL: indicating file extension of config.yaml
	viper.SetConfigType(configType)

	// Reading our config.yaml file.
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	// OPTIONAL: Set value of undefined keys from config file.
	//viper.SetDefault("<key1.key2>", "<value>")

	// OPTIONAL: Decode config file into Go Struct.
	//err := viper.Unmarshal(&<varname>)
	//if err != nil {
	//	fmt.Printf("Are the tags of the fields properly set?, %v", err)
	//}
}

// Allow viper to check for ENVIRONMENT VARIABLES
func SetEnv() {
	viper.AutomaticEnv()
}

// Return index of specified funder ID from db.Funders
func IndexOf(id uint64) (uint64, error) {
	for index, funder := range db.Funders {
		if funder.Id == id {
			return uint64(index), nil
		}
	}

	errorMessage := fmt.Sprintf("ID %v: funder not found.", id)
	return 0, errors.New(errorMessage)
}

// Returns "2022-02-17 18:21:41.64419 +0000 UTC" time format of current time.
func GetCurrentTime() string {
	loc, _ := time.LoadLocation("UTC")
	now := (time.Now().In(loc)).String()
	return now
}
