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
	Dates        []string
	Loca         []string
	Map          Maps
}

type Location struct {
	Locations []string
}

type Date struct {
	Dates []string
}

type Maps struct {
	Results []map[string]map[string]map[string]string
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
	/*
		for t := 0; t < len(tabData); t++ {
			tabLoca := tabData[t].Loca
			for j := 0; j < len(tabLoca); j++ {
				var mapi Maps
				query := tabLoca[j]
				url := "https://api.opencagedata.com/geocode/v1/json?q=" + query + "&key=ba772045bfb044078998edd6c4dc3c5a"
				containMap, _ := http.Get(url)

				defer containMap.Body.Close()

				bodyMap, _ := ioutil.ReadAll(containMap.Body)

				json.Unmarshal(bodyMap, &mapi)
				//fmt.Println(mapi.Results)

				//fmt.Println(mapi.Results[0])
			}

		}
	*/
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
