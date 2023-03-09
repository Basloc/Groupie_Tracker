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

func UseApi(url string) []Artist {
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

	for i := 0; i < len(ListArt); i++ {
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
			//fmt.Println(containLoca.Locations[j])
			ListArt[i].Loca = append(ListArt[i].Loca, containLoca.Locations[j])
		}

	}

	for i := 0; i < len(ListArt); i++ {
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

func main() {

	tabData := UseApi("https://groupietrackers.herokuapp.com/api/artists")
	//fmt.Println(tabData[0].Map)
	for i := 0; i < len(tabData); i++ {
		tabData[i].Hidden = ""
		tabData[i].NotHidden = "true"
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Home(rw, r)
	})

	http.HandleFunc("/ArtistPage", func(rw http.ResponseWriter, r *http.Request) {
		ArtistPage(rw, r, &tabData)
	})

	http.HandleFunc("/calcul", func(rw http.ResponseWriter, r *http.Request) {
		/*    en faire un fonction car fonctionne   */
		input := ""
		input = r.FormValue("text")
		fmt.Println("/ArtistPage#" + input)
		for i := 0; i < len(tabData); i++ {
			if strings.ToLower(input) == strings.ToLower(tabData[i].Name) {
				http.Redirect(rw, r, "/ArtistPage#"+tabData[i].Name, http.StatusFound)
			}
			for j := 0; j < len(tabData[i].Members); j++ {
				if strings.ToLower(input) == strings.ToLower(tabData[i].Members[j]) {
					http.Redirect(rw, r, "/ArtistPage#"+tabData[i].Name, http.StatusFound)
				}
			}
		}
		/*    fin de la fonction      */

		/* faire en fonction car foctionne*/

		for i := 0; i < len(tabData); i++ {
			tabData[i].Hidden = ""
			tabData[i].NotHidden = "true"
		}

		checkboxe1, checkboxe2, checkboxe3, checkboxe4, checkboxe5, checkboxe6 := r.Form["check1"], r.Form["check2"], r.Form["check3"], r.Form["check4"], r.Form["check5"], r.Form["check6"]
		fmt.Println(checkboxe1, checkboxe2, checkboxe3, checkboxe4, checkboxe5, checkboxe6)
		fmt.Println(len(checkboxe1))
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
		http.Redirect(rw, r, "/ArtistPage", http.StatusFound)
		/*  fin de la fonct    */

		/*   faire une fonction avec  */
		date := r.FormValue("date")
		fmt.Println(date)
		for i := 0; i < len(tabData); i++ {
			for j := 0; j < len(tabData[i].Dates); j++ {
				//faire le test savoir si la taille vaut 11 ou 10 pour savoir si il y a une etoile ou non
			}
		}
		a := "100"
		b := 0
		_, err := fmt.Sscan(a, &b)
		fmt.Println(a, b, err)
		/*   fin de la fonction      */
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

	ft := http.FileServer(http.Dir("./template/"))
	http.Handle("/template/", http.StripPrefix("/template/", ft))
}
