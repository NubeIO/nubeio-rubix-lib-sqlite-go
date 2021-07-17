package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/pkg/database"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/sql_config"
	"gorm.io/gorm"
	"log"
)



func main() {

	type TestTable struct {
		Name        	string
		Description 	string
	}

	type TestTable2nd struct {
		Name       		string
		Description 	string
	}

	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}

	var models = []interface{}{
		&TestTable{},  &TestTable2nd{},  &Product{},
	}
	var args sql_config.Params
	args.UseConfigFile = false

	var config sql_config.Database
	config.DbName = "test.db"
	config.DbPath = "./"

	err := sql_config.SetSqliteConfig(config, args); if err != nil {
		log.Println(err)
		return
	}
	err = database.SetupDB(models)
	if err != nil {
		log.Println(2222, err)
	}
	var db = database.GetDB()
	db.Create(&Product{Code: "D42", Price: 100})
	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	db.First(&product, 1) // find product with integer primary key
	fmt.Printf("%+v\n", product)

}
