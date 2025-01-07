package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Alice", Email: "alice@example.com"}

	// JSON Marshalling
	data, _ := json.Marshal(user)
	fmt.Println("JSON:", string(data))

	// JSON Unmarshalling
	var newUser User
	json.Unmarshal(data, &newUser)
	fmt.Println("User Struct:", newUser)

// JSON Marshalling: Converts Go struct to JSON string.
// JSON Unmarshalling: Converts JSON string back to Go struct.
}
