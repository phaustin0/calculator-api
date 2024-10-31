package main

import (
	"log"
	"net/http"
)

func handleRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /add", add)
	mux.HandleFunc("POST /subtract", subtract)
	mux.HandleFunc("POST /multiply", multiply)
	mux.HandleFunc("POST /divide", divide)
	mux.HandleFunc("POST /sum", sum)
	return mux
}

func main() {
	routeHandler := handleRoutes()
	log.Fatal(http.ListenAndServe("localhost:8000", routeHandler))
}
