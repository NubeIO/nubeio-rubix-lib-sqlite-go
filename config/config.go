package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)



type Configuration struct {
	Database struct {
		Driver   	string
		Name   		string
		Path 		string
		Logging   	bool

	}
}


func SqliteConfig(configFile string, generateFile bool) Configuration {
	file, err := os.Open(configFile)
	Config := Configuration{}
	if err != nil {
		Config.Database.Driver = "sqlite"
		Config.Database.Name = "test.db"
		Config.Database.Path = "./"
		Config.Database.Logging = false
		// generate config file if dont exist
		if generateFile {
			j, _ := json.Marshal(Config)
			err = ioutil.WriteFile(configFile, j, 0644)
		}
		return Config
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
		return Config
	}

}



