package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Artist struct { // Mettre exactement les meme noms pour les attributs avec une majuscule pour que le json marche
	Locations []string
	Id        int
	Dates     string
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/locations/1"
	var ListArt []Artist
	names, err := http.Get(url) // API pour les artistes et le liens emmene au prememier artiste
	fmt.Println(names)
	if err != nil {
		log.Fatal(err)
	}

	defer names.Body.Close()

	body, err := ioutil.ReadAll(names.Body)

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &ListArt)
	fmt.Println(ListArt)
	//fmt.Println(listart[0].Name)
	for i := 0; i <= 4; i++ {
		fmt.Println(ListArt[i])
	}

}
