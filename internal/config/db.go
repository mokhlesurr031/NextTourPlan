package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Database holds the database configuration
type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

var (
	db Database
)

// DB returns the default database configuration
func DB() *Database {
	return &db
}

// LoadDB loads database configuration
func loadDB() {
	//viper.AddConfigPath("conf")
	viper.SetConfigFile("config.yml")
	er := viper.ReadInConfig()
	if er != nil {
		log.Println(er)
	}

	currentDB := viper.GetString("current_db.RUNNING")

	fmt.Println("CurrentDB", currentDB)
	db = Database{
		Name:     viper.GetString(currentDB + ".POSTGRESQL_ADDON_DB"),
		Username: viper.GetString(currentDB + ".POSTGRESQL_ADDON_USER"),
		Password: viper.GetString(currentDB + ".POSTGRESQL_ADDON_PASSWORD"),
		Host:     viper.GetString(currentDB + ".POSTGRESQL_ADDON_HOST"),
		Port:     viper.GetInt(currentDB + ".POSTGRESQL_ADDON_PORT"),
	}
}
