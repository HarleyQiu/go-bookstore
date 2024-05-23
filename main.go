package main

import (
	"github.com/HarleyQiu/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Println("Server is starting on localhost:9010...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
