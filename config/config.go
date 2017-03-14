package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	EnvMode    string
	GoMaxProcs int

	Cluster                   string
	Database                  string
	DriversLocationCollection string

	DefaltRadius        int64
	DefaultResultsLimit int64
	DriverIdStart       int64
	DriverIdEnd         int64

	TestDatabase string
}

func (conf *config) setDefaultConfig() {
	conf.EnvMode = "dev"
	conf.GoMaxProcs = 2
	conf.Cluster = "localhost"
	conf.Database = "mydriver"
	conf.DriversLocationCollection = "driver"
	conf.DefaltRadius = 500
	conf.DefaultResultsLimit = 10
	conf.DriverIdStart = 1
	conf.DriverIdEnd = 50000
	conf.TestDatabase = "testdb"
}

var Setting config

func init() {
	log.Println("Loading configurations...")
	dir, _ := os.Getwd()

	// order in which to search for config file
	files := []string{
		dir + "/dev.ini",
		dir + "/config.ini",
		dir + "/conf/dev.ini",
		dir + "/conf/config.ini",
	}

	for _, f := range files {
		if _, err := toml.DecodeFile(f, &Setting); err == nil {
			log.Printf("Loaded configuration %s", f)
			return
		}
	}

	if len(Setting.EnvMode) < 3 {
		log.Println("Configuration files are not loaded properly, problem in finding Environment Mode to run application.")
		log.Println("Loading default configurations")
		Setting.setDefaultConfig()
		log.Printf("%+v\n", Setting)
	}

}
