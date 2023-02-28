package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Artist struct {
	Name  string
	Age   int
	Image string
}

func ArtistPage(rw http.ResponseWriter, r *http.Request, data *[]Artist) {
	template, err := template.ParseFiles("./ArtistPage.html", "./template/whitebox.html", "./static/style.css", "./static/styles.css", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, data)
}

func Home(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./homepage.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, r)

}

func main() {

	url := "https://groupietrackers.herokuapp.com/api/artists"
	var ListArt []Artist
	names, err := http.Get(url) // API pour les artistes et le liens emmene au prememier artiste

	if err != nil {
		log.Fatal(err)
	}

	defer names.Body.Close()

	body, err := ioutil.ReadAll(names.Body)

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &ListArt)
	//fmt.Println(listart[0].Name)
	for i := 0; i <= 4; i++ {
		fmt.Println(ListArt[i].Name)

		var tabData []Artist
		tabData = append(tabData, ListArt...)
		// //Data := &Artist{"XXXtentacion", 21}
		// tabData = append(tabData, Artist{"XXXtentacion", 21, "./static/téléchargé.png"})
		// tabData = append(tabData, Artist{"Lil Peep", 20, "./static/téléchargé.png"})
		// tabData = append(tabData, Artist{"columbine", 21, "./static/téléchargé.png"})
		// tabData = append(tabData, Artist{"lorenzo ", 21, "./static/téléchargé.png"})

		http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			Home(rw, r)
		})

		http.HandleFunc("/ArtistPage", func(rw http.ResponseWriter, r *http.Request) {
			ArtistPage(rw, r, &tabData) // data = struct pour les artist
		})

		fs := http.FileServer(http.Dir("./static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		http.ListenAndServe(":8080", nil)

		fi := http.FileServer(http.Dir("./template/"))
		http.Handle("/template/", http.StripPrefix("/template/", fi))
	}
}
