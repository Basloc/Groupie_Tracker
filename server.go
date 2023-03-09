package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Artist struct { // Struct pour acceuillir les infos de l api groupie
	Name         string
	Members      []string
	CreationDate int
	Image        string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Dates        []string
	Loca         []string
	Hidden       string
	NotHidden    string
}

type Location struct { // Struct pour receuillir les locations de concert
	Locations []string
}

type Date struct { // struct pour receuillir les dates de concert
	Dates []string
}

func ArtistPage(rw http.ResponseWriter, r *http.Request, data *[]Artist) { // fonction contenant toutes les templates et fichier pour créer la page artiste
	template, err := template.ParseFiles("./ArtistPage.html", "./template/whitebox.html", "./static/style.css", "./static/styles.css", "./template/header.html", "./template/footer.html", "./template/tempArtist.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, data)
}

func Home(rw http.ResponseWriter, r *http.Request) { // fonction contenant toute les templates et fichier pour creer la home page
	template, err := template.ParseFiles("./homepage.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, r)

}

func UseApi(url string) []Artist { // fonction permettant de créer tout les artistes
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

	for i := 0; i < len(ListArt); i++ { // boucle pour receuillir les localisation
		var containLoca Location
		urlLoca := ListArt[i].Locations
		location, err := http.Get(urlLoca)

		if err != nil {
			log.Fatal(err)
		}

		defer location.Body.Close()

		body2, err := ioutil.ReadAll(location.Body)

		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body2, &containLoca)
		for j := 0; j < len(containLoca.Locations); j++ {
			ListArt[i].Loca = append(ListArt[i].Loca, containLoca.Locations[j])
		}

	}

	for i := 0; i < len(ListArt); i++ { // boucles pour receuillir les dates
		var containeDate Date
		urlDate := ListArt[i].ConcertDates
		date, err := http.Get(urlDate)

		if err != nil {
			log.Fatal(err)
		}

		defer date.Body.Close()

		bodyDate, err := ioutil.ReadAll(date.Body)

		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(bodyDate, &containeDate)
		for j := 0; j < len(containeDate.Dates); j++ {
			ListArt[i].Dates = append(ListArt[i].Dates, containeDate.Dates[j])
		}

	}

	tabData = append(tabData, ListArt...)
	return tabData
}

func searchBar(input string, tabData []Artist) string { // fonction pour la barre de recherche
	var redirect string
	for i := 0; i < len(tabData); i++ {
		if strings.ToLower(input) == strings.ToLower(tabData[i].Name) {
			redirect = tabData[i].Name
		}
		for j := 0; j < len(tabData[i].Members); j++ {
			if strings.ToLower(input) == strings.ToLower(tabData[i].Members[j]) {
				redirect = tabData[i].Name
			}
		}
	}
	return redirect
}

func filterMember(checkboxe1 []string, checkboxe2 []string, checkboxe3 []string, checkboxe4 []string, checkboxe5 []string, checkboxe6 []string, tabData []Artist) { // fonction pour les filtre membre
	if len(checkboxe1) != 0 {
		for i := 0; i < len(tabData); i++ {
			if len(tabData[i].Members) > 1 {
				tabData[i].Hidden = "true"
				tabData[i].NotHidden = ""

			}
		}
	} else if len(checkboxe2) != 0 {
		for i := 0; i < len(tabData); i++ {
			if len(tabData[i].Members) > 2 {
				tabData[i].Hidden = "true"
				tabData[i].NotHidden = ""

			}
		}
	} else if len(checkboxe3) != 0 {
		for i := 0; i < len(tabData); i++ {
			if len(tabData[i].Members) > 3 {
				tabData[i].Hidden = "true"
				tabData[i].NotHidden = ""

			}
		}
	} else if len(checkboxe4) != 0 {
		for i := 0; i < len(tabData); i++ {
			if len(tabData[i].Members) > 4 {
				tabData[i].Hidden = "true"
				tabData[i].NotHidden = ""

			}
		}
	} else if len(checkboxe5) != 0 {
		for i := 0; i < len(tabData); i++ {
			if len(tabData[i].Members) > 5 {
				tabData[i].Hidden = "true"
				tabData[i].NotHidden = ""

			}
		}
	} else if len(checkboxe6) != 0 {
		for i := 0; i < len(tabData); i++ {
			if len(tabData[i].Members) > 6 {
				tabData[i].Hidden = "true"
				tabData[i].NotHidden = ""

			}
		}
	} else {
		for i := 0; i < len(tabData); i++ {
			tabData[i].Hidden = ""
			tabData[i].NotHidden = "true"
		}

	}
}

func filterDate(dateInt int, tabData []Artist) { // fonction pour le filtre date de concert
	for i := 0; i < len(tabData); i++ {
		for j := 0; j < len(tabData[i].Dates); j++ {
			if len(tabData[i].Dates[j]) == 11 {
				str := tabData[i].Dates[j][7:]
				newInt := 0
				_, err := fmt.Sscan(str, &newInt)
				if err != nil {
					fmt.Println(err)
				}
				if dateInt < newInt {
					tabData[i].Hidden = "true"
					tabData[i].NotHidden = ""
				} else {
					tabData[i].Hidden = ""
					tabData[i].NotHidden = "true"
				}
			} else if len(tabData[i].Dates[j]) == 10 {
				str := tabData[i].Dates[j][6:]
				newInt := 0
				_, err := fmt.Sscan(str, &newInt)
				if err != nil {
					fmt.Println(err)
				}
				if dateInt < newInt {
					tabData[i].Hidden = "true"
					tabData[i].NotHidden = ""
				} else {
					tabData[i].Hidden = ""
					tabData[i].NotHidden = "true"
				}
			}
		}
	}
}

func filterCreation(tabData []Artist, creationInt int) { // fonction du filtre de date de creation
	for i := 0; i < len(tabData); i++ {
		if creationInt < tabData[i].CreationDate {
			tabData[i].Hidden = "true"
			tabData[i].NotHidden = ""
		} else {
			tabData[i].Hidden = ""
			tabData[i].NotHidden = "true"
		}
	}
}

func main() {

	tabData := UseApi("https://groupietrackers.herokuapp.com/api/artists") // creation du tableau d artist
	//fmt.Println(tabData[0].Map)
	for i := 0; i < len(tabData); i++ { // Initialisation des variable Hidden et NotHidden pour afficher tout les artiste par default
		tabData[i].Hidden = ""
		tabData[i].NotHidden = "true"
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { // creation de la route principale de la homepage
		Home(rw, r)
	})

	http.HandleFunc("/ArtistPage", func(rw http.ResponseWriter, r *http.Request) { // creation de la route artist page
		ArtistPage(rw, r, &tabData)
	})

	http.HandleFunc("/calcul", func(rw http.ResponseWriter, r *http.Request) { // creation d'une route calcul appeler pour tout ce qui est filtre et serchbar
		for i := 0; i < len(tabData); i++ {
			tabData[i].Hidden = ""
			tabData[i].NotHidden = "true"
		}
		input := ""
		input = r.FormValue("text")
		fmt.Println("/ArtistPage#" + input)
		if input != "" { // verification si un information est rentrer dans la barre de recherche
			redirect := searchBar(input, tabData)
			http.Redirect(rw, r, "/ArtistPage#"+redirect, http.StatusFound)
		}

		// creation de varaiable qui donneront l état de nos checkboxes
		checkboxe1, checkboxe2, checkboxe3, checkboxe4, checkboxe5, checkboxe6 := r.Form["check1"], r.Form["check2"], r.Form["check3"], r.Form["check4"], r.Form["check5"], r.Form["check6"]
		filterMember(checkboxe1, checkboxe2, checkboxe3, checkboxe4, checkboxe5, checkboxe6, tabData)

		date := r.FormValue("date") // recuperation de la valeur sur le filtre des dates de concert
		dateInt := 0                // valeur receptacle

		_, err := fmt.Sscan(date, &dateInt) //transformation de la string date en int dans le receptacle
		fmt.Println(dateInt)
		if dateInt != 0 { // test si le filtre est utiliser ou non
			filterDate(dateInt, tabData)
		}
		if err != nil {
			fmt.Println(" erreur :", err)
		}

		creation := r.FormValue("creation")
		creationInt := 0
		_, err = fmt.Sscan(creation, &creationInt)
		if creationInt != 0 {
			filterCreation(tabData, creationInt)
		}

		http.Redirect(rw, r, "/ArtistPage", http.StatusFound) // redirige sur la page artist avec les changements des etat visible ou non
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	ft := http.FileServer(http.Dir("./template/"))
	http.Handle("/template/", http.StripPrefix("/template/", ft))

	http.ListenAndServe(":8080", nil)
}
