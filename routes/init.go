package routes

import (
	"github.com/gorilla/mux"
	home "github.com/oasiscse/go-docker-deployment/api"
	"log"
	"net/http"
)

func GetRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
