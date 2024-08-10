package main

import (
	"encoding/json"
	"fmt"
	// "html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type Restaurant struct {
	ID int `json:"id"`
	Restoname string `json:"name"`
	Location string `json:"location"`
}

var users []User
var restaurants []Restaurant

func main() {
	//Using gorilla mux router
	router:=mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/home", homeHandler)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/createUser", createUser).Methods("POST")
	router.HandleFunc("/restaurants", getRestaurants).Methods("GET")
	router.HandleFunc("/createRestaurant", createRestaurant).Methods("POST")
	//Using net/hhtp
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/home", homeHandler)
	fmt.Println("http://localhost:8080/ is running")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello bud!")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to home")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(restaurants)
}

func createRestaurant(w http.ResponseWriter, r *http.Request) {
	var newRestaurant Restaurant
	_ = json.NewDecoder(r.Body).Decode(&newRestaurant)
	restaurants = append(restaurants, newRestaurant)
	json.NewEncoder(w).Encode(newRestaurant)
}