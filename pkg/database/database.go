package database

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/sql_config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

)

var (
	DB    *gorm.DB
	err   error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB(models []interface{}) error {
	var db = DB
	commonConfig := sql_config.GetSqliteConfig()
	fmt.Println(commonConfig)
	dbName := commonConfig.DbName
	dbPath := commonConfig.DbPath
	logs := commonConfig.Logging
	path := fmt.Sprintf("%s%s?_foreign_keys=on", dbPath, dbName)
	fmt.Println(path)
	if logs {
		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		})
	}
	if err != nil {
		return errors.New("init db issue")
	}
	// Auto migrate project models
	for _, v := range models {
		err = db.AutoMigrate(v); if err != nil {
			return errors.New("db migrate issue")
		}
	}
	DB = db
	return nil
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
