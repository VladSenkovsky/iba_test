package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	file, err := os.OpenFile("data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shot", createShot).Methods("POST")
	router.HandleFunc("/shots/{date}", getBoundaryValuesOfSpeed).Methods("GET")
	router.HandleFunc("/shots/{date}/{speed}", getViolationsByDateAndSpeed).Methods("GET")
	time.AfterFunc(time.Second, func() { log.Fatal(http.ListenAndServe(":8080", router)) })
	writeNewShot(file)
}
