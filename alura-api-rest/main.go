package main

import (
	"fmt"

	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest/database"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest/models"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "historia 2"},
	}

	database.ConectDB()

	fmt.Println("Iniciando serviço ...")
	routes.HandleRequest()
}
