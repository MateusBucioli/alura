package routes

import (
	"goRest/controllers"
	"goRest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadeById).Methods("Get")

	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods("Post")

	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("Delete")

	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(handlers.AllowedOrigins([]string{"localhost"}))(r)))
}
