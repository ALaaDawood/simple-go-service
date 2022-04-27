package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

const (
	port = ":8000"
)

var users []User

func main() {
	router := mux.NewRouter()

	users = append(users, User{ID: "1", Firstname: "Alaa", Lastname: "Dawood", Address: &Address{Country: "Egypt", City: "Tanta", Street: "Mahkama st"}})
	users = append(users, User{ID: "2", Firstname: "Yahya", Lastname: "Qandel", Address: &Address{Country: "Egypt", City: "Qalyub", Street: "Mahkama st"}})

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Printf("listening on port: %v", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var updatedUser User
			_ = json.NewDecoder(r.Body).Decode(&updatedUser)
			updatedUser.ID = params["id"]
			users = append(users, updatedUser)
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
