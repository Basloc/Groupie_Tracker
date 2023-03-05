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
	Name         string
	Members      []string
	CreationDate int
	Image        string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Dates        Date
	Loca         Location
}

type Location struct {
	Locations []string
}

type Date struct {
	Dates []string
}

func ArtistPage(rw http.ResponseWriter, r *http.Request, data *[]Artist) {
	template, err := template.ParseFiles("./ArtistPage.html", "./template/whitebox.html", "./static/style.css", "./static/styles.css", "./template/header.html", "./template/footer.html", "./template/tempArtist.html")
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
	names, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer names.Body.Close()

	body, err := ioutil.ReadAll(names.Body)

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &ListArt)

	var tabData []Artist
	// possibilite de juste envoyer listart et non tabdata
	// --------------------------------------------------------------------

	for i := 0; i < len(ListArt); i++ {
		var containLoca Location
		urlo := ListArt[i].Locations
		fmt.Println(urlo)
		nome, err := http.Get(urlo)

		if err != nil {
			log.Fatal(err)
		}

		defer nome.Body.Close()

		body2, err := ioutil.ReadAll(nome.Body)

		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body2, &containLoca)
		fmt.Println(ListArt[i].Name)
		fmt.Println(i)
		fmt.Println("reel location : ", containLoca)
		ListArt[i].Loca = containLoca
		fmt.Println("location in struct : ", ListArt[i].Loca, "\n")
	}
	tabData = append(tabData, ListArt...)
	fmt.Println(ListArt[0].Loca)
	fmt.Println(ListArt[0].Name)

	// --------------------------------------------------------------------

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

	http.HandleFunc("/ArtistPage", func(rw http.ResponseWriter, r *http.Request) {
		ArtistPage(rw, r, &tabData)
	})

	http.HandleFunc("/calcul", func(rw http.ResponseWriter, r *http.Request) {
		input := ""
		input = r.FormValue("text")
		fmt.Println("/ArtistPage#" + input)
		http.Redirect(rw, r, "/ArtistPage#"+input, http.StatusFound)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

	ft := http.FileServer(http.Dir("./template/"))
	http.Handle("/template/", http.StripPrefix("/template/", ft))
}
