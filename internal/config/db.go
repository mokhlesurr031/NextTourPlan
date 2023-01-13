package config

import (
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
	db = Database{
		Name:     viper.GetString("clever_cloud_db.POSTGRESQL_ADDON_DB"),
		Username: viper.GetString("clever_cloud_db.POSTGRESQL_ADDON_USER"),
		Password: viper.GetString("clever_cloud_db.POSTGRESQL_ADDON_PASSWORD"),
		Host:     viper.GetString("clever_cloud_db.POSTGRESQL_ADDON_HOST"),
		Port:     viper.GetInt("clever_cloud_db.POSTGRESQL_ADDON_PORT"),
	}
}
