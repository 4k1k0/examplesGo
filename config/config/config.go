package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	if config == nil {
		log.Println("Loading config file...")
		file, _ := os.Open("config.json")
		defer file.Close()

		decoder := json.NewDecoder(file)
		var conf Config
		err := decoder.Decode(&conf)
		if err != nil {
			conf.Name = "Default Name"
			conf.Age = 21
		}

		once.Do(func() {
			config = &Config{
				Name: conf.Name,
				Age:  conf.Age,
			}
		})
	}

	return config
}
