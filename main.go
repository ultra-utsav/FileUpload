package main

import (
	"net/http"

	"github.com/rs/cors"

	"./handler"
	"github.com/gorilla/mux"
)

func main() {
	//mux router
	r := mux.NewRouter()

	//routes
	r.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public/"))))
	r.HandleFunc("/upload", handler.ResumableUpload)
	r.HandleFunc("/", handler.IndexHandler)

	//cors enable
	handler := cors.Default().Handler(r)

	// listening on port 8000
	http.ListenAndServe(":8000", handler)
}
