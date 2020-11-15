package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/ultra-utsav/FileUpload.git/handler"
)

func main() {
	//mux router
	r := mux.NewRouter()

	//routes
	r.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public/"))))
	// POST request to upload file.
	r.HandleFunc("/upload", handler.ResumableUpload)
	r.HandleFunc("/", handler.IndexHandler)

	//cors enable
	handler := cors.Default().Handler(r)

	// listening on port 8000
	log.Println("Starting server...")
	http.ListenAndServe(":8000", handler)
}
