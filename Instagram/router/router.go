package router

import (
	"github.com/gorilla/mux"
	"github.com/kunalk27/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/posts", controller.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controller.GetPost).Methods("GET")
	router.HandleFunc("/posts/users/{id}", controller.ListPostsofUser).Methods("GET")
	return router
}
