package routes

import (
	"github.com/gorilla/mux"
	"github.com/satishgowda28/go_projects/go_book-management_db/pkg/controllers"
)


var RegisterRoutes = func (r *mux.Router) {
	r.HandleFunc("/book", controllers.CreatBook).Methods("POST")
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}