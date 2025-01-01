package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	helloWorld := HelloWorld{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(helloWorld)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorldHandler).Methods("GET")

	fmt.Println("Server running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
