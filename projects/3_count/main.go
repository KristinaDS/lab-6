package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var globCount int

func main() {
	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		log.Println(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(globCount)
}

func handlePost(w http.ResponseWriter, r *http.Request) {

	intCount, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("это не число"))
	} else {
		globCount += intCount
	}
}
