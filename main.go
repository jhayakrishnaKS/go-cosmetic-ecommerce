package main

import (
	"ecommerce/src/config"
	"ecommerce/src/routes"
	"ecommerce/src/utils/db"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	ENVDev = "dev"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = ENVDev
	}
	file, err := os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to get env file.Err:", err)
		os.Exit(1)
	}
	err = config.Parse(config.TypeJSON, file)
	if err != nil {
		log.Println("Unable to parse json env file.Err:", err)
		os.Exit(1)
	}
	db.Init()

	r := routes.GetRouter()
	r.Run(":" + config.Conf.Port)
	fmt.Printf("%+v", config.Conf)
}
