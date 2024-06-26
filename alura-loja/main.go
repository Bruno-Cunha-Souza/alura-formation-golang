package main

import (
	"net/http"

	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
