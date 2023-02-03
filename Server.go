package main

import (
	"log"
	"net/http"
	"text/template"
)

type Artist struct {
}

func Home(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./home.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, r)

}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

	http.ListenAndServe(":8080", nil)
}
