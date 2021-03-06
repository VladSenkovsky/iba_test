package main

import (
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func formatDate(year int, month string, day int) string {
	if month == "January" {
		month = "01"
	} else if month == "February" {
		month = "02"
	} else if month == "March" {
		month = "03"
	} else if month == "April" {
		month = "04"
	} else if month == "May" {
		month = "05"
	} else if month == "June" {
		month = "06"
	} else if month == "July" {
		month = "07"
	} else if month == "August" {
		month = "08"
	} else if month == "September" {
		month = "09"
	} else if month == "October" {
		month = "10"
	} else if month == "November" {
		month = "11"
	} else if month == "December" {
		month = "12"
	}
	return strconv.Itoa(day) + "." + month + "." + strconv.Itoa(year)
}

func formatTime(hour int, min int, sec int) string {
	return strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
}

func generateID() string {
	return strconv.Itoa(rand.Intn(8999)+1000) + " PP-7"
}

func generateSpeed() string {
	return strconv.FormatFloat(math.Round(((float64(rand.Intn(700)+200))*0.1)*100)/100, 'f', 1, 64)
}

//constantly adding new shots with some interval while server is running
func writeNewShot(f *os.File) {
	byteValue, _ := ioutil.ReadAll(f)
	if byteValue != nil {
		f.WriteString(", {\n\t\"Date\": \"")
	} else {
		f.WriteString("{\n\t\"Date\": \"")
	}
	year, month, day := time.Now().Date()
	dateString := formatDate(year, month.String(), day)
	f.WriteString(dateString)
	hour, min, sec := time.Now().Clock()
	timeString := formatTime(hour, min, sec)
	f.WriteString("\",\n\t\"Time\": \"")
	f.WriteString(timeString)
	f.WriteString("\",\n\t\"ID\": \"")
	idString := generateID()
	f.WriteString(idString)
	f.WriteString("\",\n\t\"Speed\": \"")
	speedString := generateSpeed()
	f.WriteString(speedString)
	f.WriteString("\"\n}")
	time.Sleep(300 * time.Second)
	writeNewShot(f)
}
