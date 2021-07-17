package sql_config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var (
	C Database
)


type Database struct {
		DbName   		string
		DbPath 		string
		Logging   	bool
}

type Params struct 	{
	UseConfigFile bool
	ConfigFile    string
	GenerateFile  bool
}

//SetSqliteConfig
// if args Params.GenerateFile is true this will create a json mqtt_config file and will disregard the mqtt_config Broker
// if args Params.UseConfigFile is true is will use a local mqtt_config file
func SetSqliteConfig (config Database, args Params) error {
	if !args.UseConfigFile {
		C = config
	} else {
		configFile := args.ConfigFile
		if configFile == "" {
			return errors.New("no valid config file passed in")
		}
		file, err := os.Open(configFile)
		_config := Database{}
		if err != nil {
			_config.DbName = "test.db"
			_config.DbPath = "./"
			_config.Logging = false
			// generate mqtt_config file if dont exist
			if args.GenerateFile {
				j, _ := json.Marshal(_config)
				err = ioutil.WriteFile(configFile, j, 0644)
			}
			C = _config
		} else {
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&_config)
			C = _config
		}
	}
	return nil
}

func GetSqliteConfig()  Database {
	return C
}




