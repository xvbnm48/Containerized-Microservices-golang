package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	response := map[string]string{
		"status": "UP",
	}

	json.NewEncoder(w).Encode(response)
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hello, you have requested %s , with token %s \n ", r.URL.Path, r.URL.Query().Get("token"))
//
//}
//
//func main() {
//	http.HandleFunc("/", rootHandler)
//
//	fs := http.FileServer(http.Dir("static/"))
//	http.Handle("/static/", http.StripPrefix("/static/", fs))
//	log.Println("start")
//	http.ListenAndServe(":8080", nil)
//}
