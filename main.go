package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Food Model
type Food struct {
	Id       string  `json:"id"`
	FoodName string  `json:"foodname"`
	Foodtype string  `json:"foodtype"`
	Vendor   *Vendor `json:"vendor"`
}

//vendor model
type Vendor struct {
	VendorName     string `json:"vendorname"`
	VendorLocation string `json:"vendorlocation"`
}

// array or slice variable
var foods []Food

func main() {
	//init router
	router := mux.NewRouter()

	// fake database
	foods = append(foods, Food{
		Id:       "1",
		FoodName: "Fried Rice",
		Foodtype: "fast food",
		Vendor: &Vendor{
			VendorName:     "Stone Castle",
			VendorLocation: "Gida Dubu"}})
	foods = append(foods, Food{
		Id:       "2",
		FoodName: "Eba & Egusi",
		Foodtype: "African dish",
		Vendor: &Vendor{
			VendorName:     "Ego chops",
			VendorLocation: "Yalwawa"}})
	foods = append(foods, Food{
		Id:       "3",
		FoodName: "Fried yam",
		Foodtype: "fast food",
		Vendor: &Vendor{
			VendorName:     "Stone Castle",
			VendorLocation: "Gida Dubu"}})

	// Endpoints
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/getfoods", getFoods).Methods("GET")
	router.HandleFunc("/getfoods/{id}", getFoods).Methods("GET")
	router.HandleFunc("/getfood/{id}", getFood).Methods("GET")
	router.HandleFunc("/getfoods", createFood).Methods("POST")
	router.HandleFunc("/getfoods/{id}", updateFood).Methods("PUT")
	router.HandleFunc("/getfoods/{id}", deleteFood).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8000", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello Welcome to Enike's food endpoints.. \\getfoods - to get all food")
}
func getFoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}
func getFood(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("content-Type", "application/json")
	for _, item := range foods {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Food{})
}
func createFood(w http.ResponseWriter, r *http.Request) {

}
func updateFood(w http.ResponseWriter, r *http.Request) {

}
func deleteFood(w http.ResponseWriter, r *http.Request) {

}
