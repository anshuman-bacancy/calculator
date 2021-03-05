package main

import (
	"calculator/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/solve", controller.Solve).Methods("POST")

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/static/", router)

	log.Println("Server started at port... 8000")
	http.ListenAndServe(":8000", router)
}
