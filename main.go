package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	time := r.URL.Path[1:]
	fmt.Fprintln(w, "time: ", time)
}
