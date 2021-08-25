package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//init router
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/getVendors", getVendors).Methods("GET")
	router.HandleFunc("/getVendors/{id}", getVendors).Methods("GET")
	router.HandleFunc("/getVendors", createVendor).Methods("POST")
	router.HandleFunc("/getVendors/{id}", updateVendor).Methods("PUT")
	router.HandleFunc("/getVendors/{id}", deleteVendor).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8000", router))

}

func index(w http.ResponseWriter, r *http.Request) {

}
func getVendors(w http.ResponseWriter, r *http.Request) {

}
func createVendor(w http.ResponseWriter, r *http.Request) {

}
func updateVendor(w http.ResponseWriter, r *http.Request) {

}
func deleteVendor(w http.ResponseWriter, r *http.Request) {

}
