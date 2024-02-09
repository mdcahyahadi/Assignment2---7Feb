package main

import (
	"assignment-golang8-7feb/database"
	"assignment-golang8-7feb/router"
)

var PORT = ":8080"

func main() {
	database.StartConnection()
	router.StartServer().Run(PORT)
}
