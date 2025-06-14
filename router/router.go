package router

import (
	"github.com/Arnav-Agrawal-987/Libr_Prototype/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getall", controller.GetAllMessages).Methods("GET")
	router.HandleFunc("/get/{timestamp}", controller.GetMessage).Methods("GET")
	router.HandleFunc("/post", controller.PostMessage).Methods("POST")
	return router
}
