package config

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

var c config

type config struct {
	DB         Database   `toml:"database"`
	Time       Time       `toml:"time"`
	Foursquare Foursquare `toml:"foursquare"`
}

// Database ...
type Database struct {
	Driver string `toml:"driver"`
	Source string `toml:"source"`
}

// Time ...
type Time struct {
	Zone string `toml:"zone"`
}

// Foursquare ...
type Foursquare struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}

// GetDB return database configuration
func GetDB() Database {
	return c.DB
}

// GetTime return time configuration
func GetTime() Time {
	return c.Time
}

// GetFoursquare return foursquare configuration
func GetFoursquare() Foursquare {
	return c.Foursquare
}

func GetMode() string {
	return os.Getenv("BREAK_MODE")
}

func DebugMode() bool {
	return GetMode() == "debug"
}

func TestMode() bool {
	return GetMode() == "test"
}

func RleaseMode() bool {
	return GetMode() == "release"
}

func GetRootPath() string {
	return os.Getenv("BREAK_ROOT")
}

// Load ...
func Load() {
	tmlPath := path.Join(GetRootPath(), "config", GetMode(), "config.tml")
	log.WithFields(log.Fields{
		"Path": tmlPath,
	}).Infoln("Config:Load")
	_, err := toml.DecodeFile(tmlPath, &c)
	if err != nil {
		log.WithFields(log.Fields{
			"Path":  tmlPath,
			"Error": err,
		}).Panicln("Config:Load toml.DecodeFile")
	}
}
