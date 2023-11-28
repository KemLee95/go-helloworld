package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResWithData struct {
	Response
	Data interface{} `json:"data"`
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response = Response{
		Status:  true,
		Message: "Access Successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
