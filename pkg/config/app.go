package config

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	db *gorm.DB
)

type Config struct {
	Database string `json:"database"`
}

func Connect() {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		panic("Failed to read config file")
	}

	var config Config
	json.Unmarshal(configFile, &config)

	d, err := gorm.Open("mysql", config.Database)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
