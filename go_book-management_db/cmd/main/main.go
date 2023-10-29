package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satishgowda28/go_projects/go_book-management_db/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}