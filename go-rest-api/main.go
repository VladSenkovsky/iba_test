package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type shot struct {
	Date  string `json:"Date"`
	Time  string `json:"Time"`
	ID    string `json:"ID"`
	Speed string `json:"Speed"`
}

type allShots []shot

var shots = allShots{
	{
		Date:  "01.01.2000",
		Time:  "00:00:01",
		ID:    "1234 QW-1",
		Speed: "40",
	}, {
		Date:  "10.10.2002",
		Time:  "18:20:15",
		ID:    "1234 QW-1",
		Speed: "67",
	}, {
		Date:  "01.01.2000",
		Time:  "00:00:01",
		ID:    "1000 CS-1",
		Speed: "39",
	}, {
		Date:  "01.01.2000",
		Time:  "05:00:01",
		ID:    "1234 QW-1",
		Speed: "61.5",
	}, {
		Date:  "15.03.2020",
		Time:  "00:00:01",
		ID:    "7667 GF-5",
		Speed: "50",
	}, {
		Date:  "01.01.2010",
		Time:  "00:00:01",
		ID:    "1234 QW-1",
		Speed: "30",
	},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shots/{date}", getBoundaryValuesOfSpeed).Methods("GET")
	router.HandleFunc("/shots/{date}/{speed}", getViolationsByDateAndSpeed).Methods("GET")
	time.AfterFunc(4*time.Second, func() { log.Fatal(http.ListenAndServe(":8080", router)) })
	writeNewShot(file)
}
