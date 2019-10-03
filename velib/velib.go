package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func main() {

	// TODO faire un appel à une API de géoloc pour déterminer notre position
	// ici c'est la position du Fuz.re :
	lat := "48.853169"
	long := "2.402782"
	radius := "100"

	route := "https://opendata.paris.fr/api/records/1.0/search"

	// Construction basique d'une "query string" avec Sprintf
	// (pas très lisible : on pourrait stocker ces params dans une slice
	// et les assembler avec une boucle for)

	params := fmt.Sprintf("?dataset=velib-disponibilite-en-temps-reel&facet=overflowactivation&facet=creditcard&facet=kioskstate&facet=station_state&geofilter.distance=%s,%s,%s",
		lat, long, radius)

	// Assemblage route + query string avec encodage des caractères spéciaux
	url := fmt.Sprintf("%s/%s", route, html.EscapeString(params))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Err GET:", err)
		return
	}

	/*
		Variantes possibles

		* Utilisation de http.DefaultClient.Get
		OU
		1) Création d'une requête personnalisée : http.NewRequest
		2) puis appel à méthode Do sur http.DefaultClient
		OU
		Idem que précédemment + création de votre propre http.Client
		OU
		Utilisation d'une bibliothèque d'appel HTTP
	*/

	// Lecture du corps de la requête
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Err lecture du corps de la requête:", err)
	}

	velibStatus := VelibStatus{}

	if err = json.Unmarshal([]byte(jsonBytes /*jsonSample*/), &velibStatus); err != nil {
		fmt.Println("Err décodage JSON:", err)
		return
	}

	fmt.Printf("Réponse API Vélib OpenData:\n%v\n", velibStatus)
	// Vous pourriez équiper le type VelibStatus de la méthode String() string,
	// afin de contrôler l'impression d'une valeur VelibStatus par fmt
	//
	// Les fonctions de formattage de fmt testent la présence de valeurs Stringer,
	// l'ajout de cette méthode permet à VelibStatus de réaliser cette interface.
}
