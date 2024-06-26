package routes

import (
	"log"
	"net/http"

	"github.com/Bruno-Cunha-Souza/alura-formation-golang/alura-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	// gorilla/mux
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonas).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornPersona).Methods("GET")
	//
	log.Fatal(http.ListenAndServe(":8000", r))
}
