package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Alpaca struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/alpaca/", CreateAlpaca)
	log.Printf("Alpaca server up and running")
	http.ListenAndServe(":8080", nil)
}

func CreateAlpaca(w http.ResponseWriter, r *http.Request) {
	alpaca := Alpaca{}
	err := json.NewDecoder(r.Body).Decode(&alpaca)
	if err != nil {
		log.Printf("Bad Request: %v", err)
	}
	fmt.Printf("Alpaca name is %v", alpaca.Name)
}
