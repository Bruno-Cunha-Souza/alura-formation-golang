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
	r.HandleFunc("/api/personalidades", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornPersona).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersona).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersona).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersona).Methods("Put")
	//
	log.Fatal(http.ListenAndServe(":8000", r))
}
