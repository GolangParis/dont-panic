package main

// Pour lancer ce programme : go run 01_api_rest_get.go

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Distributeur struct {
	Cacolac    int // important de préfixer par une MAJ pour l'encodage JSON
	Kombucha   int // => tout champ préfixé par une minuscule est ignoré.
	BasqueCola int
}

var (
	distrib = Distributeur{
		Cacolac:    2,
		Kombucha:   1,
		BasqueCola: 3,
	}
)

func RéponseDistributeur(w http.ResponseWriter, r *http.Request) {

	jsonBytes, err := json.Marshal(distrib)
	if err != nil {
		log.Printf("Erreur json marshal: %v", err)
		return
	}

	// todo : régler le content-type serait plus conforme
	// cf. en-têtes (headers)

	// On convertir les octets de jsonBytes en chaîne de caractère
	// que l'on formatte simplement dans une chaîne de caractère
	// et fmt.Fprintf va utiliser le ResponseWriter pour retourner cette chaîne
	fmt.Fprintf(w, "%s", string(jsonBytes))
}

func main() {

	// On déclarer la fonction réponse : le handler auprès de la lib http :
	http.HandleFunc("/àboire", RéponseDistributeur)

	// lancer le serveur qui écoute sur le port 8080 sur toutes les interfaces
	// nil signifie que l'on ne désire pas passer d'instance d'un routeur (cf. gorilla ou gin par ex.)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Err ListenAndServe:", err)
	}

	// Pour tester :
	// curl localhost:8080/àboire
}
