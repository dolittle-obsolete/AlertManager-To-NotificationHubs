package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ConfigUsage is the usage string for the configuration .yaml file
const ConfigUsage = "Config file USAGE:\n" +
	"serverAddr: \n" +
	"serverPort: \n" +
	"debug: \n"

// Config is a struct containing the necessary fields
type Config struct {
	ServerAddress string `yaml:"serverAddr"`
	ServerPort    int64  `yaml:"serverPort"`
	Debug         bool   `yaml:"debug"`
}

// ParseConfig returns the Config that's parsed from the yaml file at
// configPath
func ParseConfig(configPath string) Config {
	var config Config

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Could not read configuration file at path: %s", configPath)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Something went wrong while parsing configuration file %s\n%s", configPath, ConfigUsage)
	}

	log.Printf("Successfully parsed configuration file at path %s", configPath)
	return config

}
