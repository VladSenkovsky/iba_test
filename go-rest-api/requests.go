package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Shot struct {
	Date  string `json:"Date"`
	Time  string `json:"Time"`
	ID    string `json:"ID"`
	Speed string `json:"Speed"`
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func createShot(w http.ResponseWriter, r *http.Request) {
	var newShot Shot
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newShot)
	post, _ := json.MarshalIndent(newShot, "", " ")
	file, err := os.OpenFile("data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(", " + string(post))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Thanks for new shot!\n"))
	json.NewEncoder(w).Encode(newShot)
}

//all shots that surpassed certain speed
func getViolationsByDateAndSpeed(w http.ResponseWriter, r *http.Request) {
	configuration := GetConfigTime()
	currentH, currentM, currentS := time.Now().Clock()
	currentTime := strconv.Itoa(currentH) + ":" + strconv.Itoa(currentM) + ":" + strconv.Itoa(currentS)
	cTime, _ := time.Parse("15:04:05", currentTime)
	sStart, _ := time.Parse("15:04:05", configuration.ServerStart)
	sShut, _ := time.Parse("15:04:05", configuration.ServerShutdown)
	if inTimeSpan(sStart, sShut, cTime) {
		shotsDate := mux.Vars(r)["date"]
		searchingSpeed, _ := strconv.ParseFloat(mux.Vars(r)["speed"], 64)
		var data []byte
		data, _ = ioutil.ReadFile("data.json")
		validStr := "[\n" + string(data) + "\n]"
		var shots []Shot
		if err := json.Unmarshal([]byte(validStr), &shots); err != nil {
			panic(err)
		}
		for _, singleShot := range shots {
			shotSpeed, _ := strconv.ParseFloat(singleShot.Speed, 64)
			if singleShot.Date == shotsDate && shotSpeed > searchingSpeed {
				json.NewEncoder(w).Encode(singleShot)
			}
		}
	} else {
		w.WriteHeader(503)
		w.Write([]byte("Sorry but your request is currently unavailable.\nServer schedule:\n" + configuration.ServerStart + " - " + configuration.ServerShutdown))
	}
}

//min and max speed
func getBoundaryValuesOfSpeed(w http.ResponseWriter, r *http.Request) {
	configuration := GetConfigTime()
	currentH, currentM, currentS := time.Now().Clock()
	currentTime := strconv.Itoa(currentH) + ":" + strconv.Itoa(currentM) + ":" + strconv.Itoa(currentS)
	cTime, _ := time.Parse("15:04:05", currentTime)
	sStart, _ := time.Parse("15:04:05", configuration.ServerStart)
	sShut, _ := time.Parse("15:04:05", configuration.ServerShutdown)
	if inTimeSpan(sStart, sShut, cTime) {
		shotsDate := mux.Vars(r)["date"]
		min := 1000.0
		max := 0.0
		var tempShots [2]Shot
		var data []byte
		data, _ = ioutil.ReadFile("data.json")
		validStr := "[\n" + string(data) + "\n]"
		var shots []Shot
		if err := json.Unmarshal([]byte(validStr), &shots); err != nil {
			panic(err)
		}
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
	} else {
		w.WriteHeader(503)
		w.Write([]byte("Sorry but your request is currently unavailable.\nServer schedule:\n" + configuration.ServerStart + " - " + configuration.ServerShutdown))
	}
}
