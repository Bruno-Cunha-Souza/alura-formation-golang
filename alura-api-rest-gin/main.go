package main

import (
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/database"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/routes"
)

func main() {
	database.ConectDB()
	routes.HandleRequests()
}
