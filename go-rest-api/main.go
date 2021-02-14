package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

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

func createShot(w http.ResponseWriter, r *http.Request) {
	var newShot shot
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newShot)
	shots = append(shots, newShot)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newShot)
}

func getCarShots(w http.ResponseWriter, r *http.Request) {
	shotID := mux.Vars(r)["id"]

	for _, singleShot := range shots {
		if singleShot.ID == shotID {
			json.NewEncoder(w).Encode(singleShot)
		}
	}
}
func getAllShots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shots)
}

func getViolationsByDateAndSpeed(w http.ResponseWriter, r *http.Request) {
	shotsDate := mux.Vars(r)["date"]
	searchingSpeed, _ := strconv.ParseFloat(mux.Vars(r)["speed"], 64)
	for _, singleShot := range shots {
		shotSpeed, _ := strconv.ParseFloat(singleShot.Speed, 64)
		if singleShot.Date == shotsDate && shotSpeed > searchingSpeed {
			json.NewEncoder(w).Encode(singleShot)
		}
	}
}
func getBoundaryValuesOfSpeed(w http.ResponseWriter, r *http.Request) {
	shotsDate := mux.Vars(r)["date"]
	min := 1000.0
	max := 0.0
	var tempShots [2]shot
	for _, singleShot := range shots {
		shotSpeed, _ := strconv.ParseFloat(singleShot.Speed, 64)
		if singleShot.Date == shotsDate {
			if shotSpeed > max {
				max = shotSpeed
				tempShots[1] = singleShot
			}
			if shotSpeed < min {
				min = shotSpeed
				tempShots[0] = singleShot
			}
		}
	}
	json.NewEncoder(w).Encode(tempShots[0])
	json.NewEncoder(w).Encode(tempShots[1])
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/shot", createShot).Methods("POST")
	router.HandleFunc("/shots", getAllShots).Methods("GET")
	router.HandleFunc("/shots/{id}", getCarShots).Methods("GET")
	router.HandleFunc("/shots/{date}/{speed}", getViolationsByDateAndSpeed).Methods("GET")
	router.HandleFunc("/shots/minmax/{date}", getBoundaryValuesOfSpeed).Methods("VIEW")
	log.Fatal(http.ListenAndServe(":8080", router))
}
