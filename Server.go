package main

import (
	"log"
	"net/http"
	"text/template"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./homepage.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, r)

}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

	fs := http.FileServer(http.Dir("./css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	fi := http.FileServer(http.Dir("./template/"))
	http.Handle("/template/", http.StripPrefix("/template/", fi))

	fu := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fu))

	http.ListenAndServe(":8080", nil)
}
