package main

// Pour lancer ce programme : go run 00_système_de_type.go

// go build échouera car plusieurs programmes (et donc plusieurs fonction main)
// cohabitent dans ce répertoire.
// Suggestion : essayez de repartir d'un autre répertoire vide en partageant le code
// dans plusieurs fichiers et construisez votre programme avec go build.

// En Go on écrit soit des programmes ou des bibliothèques.
// Un programme composé de plusieurs fichiers appartenant au "paquetage" principal le package main.
// Tout fichier doit démarrer par le nom de package auquel il se rapporte.
// Un package est un ensemble de fichier dans un répertoire

import (
	// ici on importe fmt la bibliothèque de formattage de chaînes de caractères
	// depuis la bibliothèque standard :
	"fmt"
)

// Moteur est un type nommé basé sur une structure (type composite),
// comportant un membre nommé rpm (rotation par minute) de type entier signé (int)
// auquel on pourra adjoindre une ou des méthodes.
type Moteur struct {
	rpm int

	// mais ces méthodes devront être déclarées en dehors du corps de cette struct
	// soit dans ce fichier, ou dans un autre fichier de ce même package
}

func (m Moteur) Démarrer() {
	// v s'appelle le receveur,
	// ici il est passé par valeur, cad qu'il ne sera pas affecté par toute modification
	// sur lui ayant lieu pdt le déroulement de la fonction
	//
	// pour le passer par référence on placerait une étoile devant le type Moteur: v *Moteur
	fmt.Println("Démarré à : ", m.rpm)
}

// Tout type nommé peut se baser sur un autre type existant
// La première qui pourrait vous venir à l'esprit est quelquechose comme :
type M Moteur

// où M est basé sur Moteur : tout M sera équipé des mêmes membres que Moteur
// mais en donnant un nom à M on le distingue du type Moteur et dans ce cas
// M ne pourrait appeler les méthodes Moteur

// L'autre manière est de composer un type avec une instance de Moteur
type LocoVapeur struct {
	m Moteur
}

// Si l'on veut reproduire le concept de réutilisation connu dans le monde orienté objet,
// on peut faire appel à une construction intérressante : l'inclusion anonyme

type MoteurVapeur struct {
	Moteur
}

// Etrangement ici on a bien un membre mais il n'est pas nommé.
// Le compilateur autorise cela et déclenche la promotion des méthodes Moteur
// sur le type MoteurVapeur

func Démarrer(mm []MoteurVapeur) {
	for _, m := range mm {
		m.Démarrer()
	}
}

func main() {
	var mv0 MoteurVapeur

	mv1 := MoteurVapeur{
		Moteur{
			rpm: 50,
		},
	}

	// On déclare un tableau de taille déterminée
	// au moment de la compilation, d'où le ...
	moteurs := [...]MoteurVapeur{
		mv0,
		MoteurVapeur{Moteur{70}},
		mv1,
	}

	// On transforme le tableau moteur en slice
	// qui est une sorte de sous-tableau dynamique
	// balayant une plage déterminée.
	// Ici sans spécification, la slice couvre le tableau entier
	Démarrer(moteurs[:])
}
