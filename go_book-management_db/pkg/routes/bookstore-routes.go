package routes

import (
	"github.com/gorilla/mux"
	"github.com/satishgowda28/go_projects/go_book-management_db/pkg/controllers"
)


var RegisterRoutes = func (r *mux.Route) {
	r.HandlerFunc("/book/", controllers.CreatBook).Methods("POST")
	r.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandlerFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandlerFunc("/book/{id}", controllers.UpdateById).Methods("PUT")
	r.HandlerFunc("/book/{id}", controllers.DeleteById).Methods("DELETE")
}