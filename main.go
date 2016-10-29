// time-stamp microservice receives timestamps and returns json form of natural/unix timestamp.
// Exercise by Free Code Camp.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Date struct returned as json.
type Date struct {
	Unix    *int64  `json:"unix"`
	Natural *string `json:"natural"`
}

func main() {
	http.HandleFunc("/", timestamp)

	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	fmt.Println("Listening on [", server.Addr, "]...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func timestamp(w http.ResponseWriter, r *http.Request) {

	tstr := r.URL.Path[1:]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("charset", "utf-8")

	t, err := returnTime(tstr)
	if err != nil {
		nilJSON, err := json.Marshal(Date{nil, nil})
		if err != nil {
			http.Error(w, "Json string failed to write", 500)
			return
		}
		w.Write(nilJSON)
		return
	}

	naturalDate := fmt.Sprintf("%v %v, %v", t.Month().String(), t.Day(), t.Year())
	naturalTime := t.Unix()

	date := Date{
		Unix:    &naturalTime,
		Natural: &naturalDate}

	responseJSON, err := json.Marshal(date)
	if err != nil {
		http.Error(w, "Json string failed to write", 500)
		return
	}

	w.Write(responseJSON)

}

// returnTime takes a string and returns the time.
// Can handle UNIX seconds input and general time in the form:
// "December 15, 2015"
func returnTime(t string) (time.Time, error) {

	// Optimistically try to parse Unix string
	if unixTime, err := strconv.Atoi(t); err == nil {
		return (time.Unix(int64(unixTime), 0)).UTC(), nil
	}

	// Assume general date form.
	genTime := strings.Split(t, " ")
	if len(genTime) != 3 {
		return time.Time{}, fmt.Errorf("general form of date not broken into 3 elements: %v", genTime)
	}

	// Re-arrange general time []string to allow time package to parse it.
	timeString := fmt.Sprintf("%s-%s-%02s", genTime[2], genTime[0][:3], genTime[1][:len(genTime[1])-1])

	parsedTime, err := time.Parse("2006-Jan-02", timeString)
	if err != nil {
		return time.Time{}, fmt.Errorf("cannot parse general time: %s", err)
	}

	return parsedTime.UTC(), nil
}
