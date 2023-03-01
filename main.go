package main

import (
	"fmt"
	"log"

	"github.com/go-gin-api/pkg/infrastructure/datastore"
	"github.com/go-gin-api/pkg/infrastructure/router"
	"github.com/go-gin-api/pkg/registry"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {

	// Load the configuration values from the file
	viper.SetConfigName("config") // name of the config file without extension
	viper.SetConfigType("json")   // type of the config file
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Create an orm instance and connect to the database.
	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	// Resolve dependendies between each layer.
	r := registry.NewRegistry(db)

	// Create controllers and routers to hanlde the requests.
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	serverAddress := viper.GetString("server.address")

	fmt.Println("Server listen at http://localhost:" + serverAddress)
	if err := e.Start(":" + serverAddress); err != nil {
		log.Fatalln(err)
	}
}
