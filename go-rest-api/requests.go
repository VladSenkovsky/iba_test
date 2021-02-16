package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func inTimeSpan(start, end, check time.Time) bool {
	fmt.Println(start, check, end)
	return check.After(start) && check.Before(end)
}

func getViolationsByDateAndSpeed(w http.ResponseWriter, r *http.Request) {
	configuration := GetConfig()
	currentH, currentM, currentS := time.Now().Clock()
	currentTime := strconv.Itoa(currentH) + ":" + strconv.Itoa(currentM) + ":" + strconv.Itoa(currentS)
	fmt.Println(configuration.ServerStart, currentTime, configuration.ServerShutdown)
	cTime, _ := time.Parse("15:04:05", currentTime)
	sStart, _ := time.Parse("15:04:05", configuration.ServerStart)
	sShut, _ := time.Parse("15:04:05", configuration.ServerShutdown)
	fmt.Println(sStart, cTime, sShut)
	if inTimeSpan(sStart, sShut, cTime) {
		shotsDate := mux.Vars(r)["date"]
		searchingSpeed, _ := strconv.ParseFloat(mux.Vars(r)["speed"], 64)
		for _, singleShot := range shots {
			shotSpeed, _ := strconv.ParseFloat(singleShot.Speed, 64)
			if singleShot.Date == shotsDate && shotSpeed > searchingSpeed {
				json.NewEncoder(w).Encode(singleShot)
			}
		}
	}
}
func getBoundaryValuesOfSpeed(w http.ResponseWriter, r *http.Request) {
	configuration := GetConfig()
	currentH, currentM, currentS := time.Now().Clock()
	currentTime := strconv.Itoa(currentH) + ":" + strconv.Itoa(currentM) + ":" + strconv.Itoa(currentS)
	fmt.Println(configuration.ServerStart, currentTime, configuration.ServerShutdown)
	cTime, _ := time.Parse("15:04:05", currentTime)
	sStart, _ := time.Parse("15:04:05", configuration.ServerStart)
	sShut, _ := time.Parse("15:04:05", configuration.ServerShutdown)
	if inTimeSpan(sStart, sShut, cTime) {
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
}
