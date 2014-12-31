package config

import (
	"fmt"
	"encoding/json"
	"os"
)

type MapType struct {
	Incoming_Url string
	Url_Pool []string
}

type Configuration struct {
	Maps []MapType
}

var fileName = "./conf.json"

func LoadConfig () Configuration {
	configuration := Configuration{}

	if _, err := os.Stat(fileName); err == nil {
		file, _ := os.Open(fileName)
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&configuration)
		if err != nil {
			fmt.Println("error:", err)
		}
	}

	return configuration
}
