package database

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB(configFile string, generateConfigFile bool, models []interface{}) {
	var db = DB
	commonConfig := config.SqliteConfig(configFile, generateConfigFile)
	driver := commonConfig.Database.Driver
	dbPath := commonConfig.Database.Path
	dbName := commonConfig.Database.Name
	logs := commonConfig.Database.Logging
	fmt.Println(222222, dbPath, dbName)
	path := fmt.Sprintf("%s%s?_foreign_keys=on", dbPath, dbName)
	fmt.Println(222222, path)



	if driver == "sqlite" {
		if logs {
			db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		} else {
			db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
			})
		}
		if err != nil {
			DBErr = err
			log.Println("db err: ", err)
		}
	}

	type TestTable struct {
		Name        	string
		Description 	string
	}

	type TestTable2nd struct {
		Name       		string
		Description 	string
	}

	//var autoMigrate = []interface{}{
	//	&TestTable{},  &TestTable2nd{},
	//}
	// Auto migrate project models
	for i, v := range models {
		log.Println("v ",i, v)
		db.AutoMigrate(v)
	}

	if err != nil {
		return
	}
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
// GetDBError helps you to get a connection
func GetDBError() error {
	return DBErr
}
