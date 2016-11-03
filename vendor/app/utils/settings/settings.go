package settings

import (
	"io/ioutil"
	"log"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// LoadConfig ...
func LoadConfig(configName string, p Parser) {
	// We need to parse config json file into Config struct
	configFile, err := ioutil.ReadFile(configName)
	if err = p.ParseJSON(configFile); err != nil {
		log.Fatalf("Couldn`t parse %s: %v", configName, configFile)
	}
}
