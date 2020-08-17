package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create", CreateStudent).Methods("POST")              // http://localhost:3000/create
	router.HandleFunc("/getstudents", GetAllStudents).Methods("GET")         // http://localhost:3000/getstudents
	router.HandleFunc("/count", CountAllStudents).Methods("GET")             // http://localhost:3000/count
	router.HandleFunc("/getone/{id}", GetOneStudent).Methods("GET")          // http://localhost:3000/getone/1
	router.HandleFunc("/deleteone/{id}", DeleteOneStudent).Methods("DELETE") // http://localhost:3000/deleteone/1
	router.HandleFunc("/deleteall", DeleteAllStudents).Methods("DELETE")     // http://localhost:3000/deleteall
	router.HandleFunc("/update/{id}", UpdateStudent).Methods("PATCH")        // http://localhost:3000/update/3
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))

}
