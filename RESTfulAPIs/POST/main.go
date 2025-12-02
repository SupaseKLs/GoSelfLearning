package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "User created: %s, Age: %d", user.Name, user.Age)
}

func main() {
	http.HandleFunc("/user", createUserHandler)
	http.ListenAndServe(":8080", nil)
}
