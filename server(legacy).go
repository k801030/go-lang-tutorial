package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type MyResponse struct {
	Text string `json:"text,omitempty"`
	Date string `json:"date,omitempty"`
	Name string `json:"name,omitempty"`
}

type geometry interface {
	area()
	perim() float64
}

type rect struct{}

func (this rect) area() {
	return
}

type MyHandler struct {
}

func (this MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get name & get date
	date := time.Now().Format("01/02/2006 15:04:05")
	name := getNameFromQuery(r)

	// generate response
	s := MyResponse{
		Text: "Hello World",
		Date: date,
		Name: name,
	}

	// write response
	jsonString, err := json.Marshal(s)
	_ = err
	fmt.Fprintf(w, string(jsonString))

	// logging
	fmt.Printf("[%s] %s %s %s\n", date, r.Host, r.URL.Path, r.URL.RawQuery)
}

func mainLegacy() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     MyHandler{},
		ReadTimeout: 3 * time.Second,
	}

	fmt.Printf("[%s] Start http server.\n", time.Now().Format("01/02/2006 15:04:05"))
	log.Fatal(server.ListenAndServe())
}

func getNameFromQuery(r *http.Request) string {
	query, ok := r.URL.Query()["name"]
	name := ""
	if query != nil {
		name = query[0]
	}
	_ = ok
	return name
}
