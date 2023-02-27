package main

import (
	"log"
	"net/http"
	"text/template"
)

type Artist struct {
	Name string
	Age  int
	Img  string
}

func ArtistPage(rw http.ResponseWriter, r *http.Request, data *[]Artist) {
	template, err := template.ParseFiles("./ArtistPage.html", "./templates/whitebox.html", "./static/style.css", "./static/styles.css")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, data)
}

func Home(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./home.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, r)

}

func main() {
	var tabData []Artist
	//Data := &Artist{"XXXtentacion", 21}
	tabData = append(tabData, Artist{"XXXtentacion", 21, "./static/téléchargé.png"})
	tabData = append(tabData, Artist{"Lil Peep", 20, "./static/téléchargé.png"})
	tabData = append(tabData, Artist{"columbine", 21, "./static/téléchargé.png"})
	tabData = append(tabData, Artist{"lorenzo ", 21, "./static/téléchargé.png"})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

	http.HandleFunc("/ArtistPage", func(rw http.ResponseWriter, r *http.Request) {
		ArtistPage(rw, r, &tabData) // data = struct pour les artist
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
