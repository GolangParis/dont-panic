package main

// Pour lancer ce programme : go run 02_slices.go

import "fmt"

func main() {
	// Ici on déclare un tableau dont la taille est déterminée
	// au moment de la compilation. Si on ajoute ou retire des élements,
	// pas besoin d'ajuster la taille du tableau, celle-ci est déterminé par le compilateur.
	tab := [...]int{1, 2, 3}

	// Une slice est un type référentiel qui pointe vers un tableau ou une slice,
	// ou bien son propre tampon mémoire "buffer".
	// Ici on déclare une slice qui couvre le tableau entier
	vals := tab[:]

	// D'après ce que vous savez,
	// qu'est ce qui sera imprimé par le bloc suivant ?
	fmt.Println("Démarrage:")
	fmt.Println("tab :", tab)
	fmt.Println("vals:", vals)
	fmt.Println()

	// Comme vous savez que les slice sont un type référentiel
	// qu'est ce qui sera imprimé par le bloc suivant ?
	// (vérifiez en lançant le programme)
	vals[2] = 4
	fmt.Println("Après écriture:")
	fmt.Println("tab :", tab)
	fmt.Println("vals:", vals)
	fmt.Println()

	vals = ConcatByVal(vals)
	// ConcatByAddr(&vals)

	// Qu'est ce qui sera imprimé par le bloc suivant ?
	fmt.Println("Après concat:")
	fmt.Println("tab :", tab)
	fmt.Println("vals:", vals)
	fmt.Println()

	// On refait quelquechose qu'on a déja fait ligne 28,
	// Et alors cette fois-ci, qu'est ce qui sera imprimé
	// par le bloc suivant ?
	vals[0] = 0

	fmt.Println("Après modif [0]:")
	fmt.Println("tab :", tab)
	fmt.Println("vals:", vals)
	fmt.Println()
	// Si vous êtes étonnés du résultat, le fait que le tableau
	// n'a pas été modifié vient du fait que la slice,
	// suite à la concaténation, a eu un dépassement de capacité,
	// et le runtime lui a donc réalloué un tampon mémoire distinct
	// du tableau sous-jacent.
	// Et donc la modification d'un élement de la slice en ligne 45
}

// Une slice est une valeur comme une autre, si on veut conserver
// l'effort de bord d'une opération sur une slice
// on peut la passer par adresse
func ConcatByAddr(vals *[]int) {
	*vals = append(*vals, 8)
}

// Mais généralement les modifications sur les slices
// sont retournées par valeur
func ConcatByVal(vals []int) []int {
	return append(vals, 8)
}
