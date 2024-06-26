package main

import (
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/models"
	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Buno", CPF: "00000000000", RG: "470000000"},
		{Nome: "Pedro", CPF: "01010000000", RG: "690000000"},
	}
	routes.HandleRequests()
}
