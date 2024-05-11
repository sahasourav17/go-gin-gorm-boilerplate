package main

import (
	"example/explore-go/config"
	"example/explore-go/schemas"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&schemas.Todo{})
}
