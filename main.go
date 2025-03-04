/*
Premier programme d'essai du langage Go, essai de Tri Bubble et essai du Tri proposé par le package sort

Points abordés :
  - Les imports
  - Les fonctions
  - La transmission des tableaux, ici de chaînes de caractères, indirectement par référence
  - Le retour de valeurs par une fonction
  - L'absence de while compensée par for
  - La portée des variables et les deux modes de déclaration, := et var
*/
package main

import "fmt"
import "sort"

func Greetings() {
	fmt.Println("Buenos días señora/señorita/señor")
}

func AfficherListes(ListeA []string, ListeB []string) {
	limite := len(ListeA)

	var LaPlusGrande []string
	if limite > len(ListeB) {
		limite = len(ListeB)
		LaPlusGrande = ListeA
	} else if len(ListeA) != len(ListeB) {
		LaPlusGrande = ListeB
	}

	fmt.Println("\nLes deux listes en même temps, sauf la fin de la plus grande")
	for i := 0; i < limite; i++ {
		fmt.Println(i, ListeA[i], ListeB[i])
	}
	if len(LaPlusGrande) != 0 {
		fmt.Println("\nReste de la liste la plus grande")
		for i := limite; i < len(LaPlusGrande); i++ {
			fmt.Println(i, LaPlusGrande[i])
		}
	}
}

func PreparationBubbleSort(ListeA []string, ListeB []string) []string {
	// Preparation for Simple Bubble Sort
	fmt.Println("\nConcaténées")
	ListeConcatenee := append(ListeA, ListeB...)
	for i, val := range ListeConcatenee {
		fmt.Println(i, val)
	}

	return ListeConcatenee
}

func TriBulle(ListeTriee []string) {
	fmt.Println("\nTri")
	limite := len(ListeTriee)
	for found := 1; found == 1; {
		found = 0
		for i := 1; i < limite; i++ {
			if ListeTriee[i-1] > ListeTriee[i] {
				swap := ListeTriee[i-1]
				ListeTriee[i-1] = ListeTriee[i]
				ListeTriee[i] = swap
				found = 1
			}
		}
	}
}

func DisplayResults(ListeTriee []string, Message string) {
	fmt.Println("\nRésultat du ", Message)
	for i, val := range ListeTriee {
		fmt.Println(i, val)
	}
}

// ======================================================================
func main() {
	ListeA := []string{"01A", "11A", "00A", "02A"}
	ListeB := []string{"03B", "10B", "07B", "05B", "64A", "00B", "99W"}

	Greetings()

	AfficherListes(ListeA, ListeB)

	// Preparation for Simple Bubble Sort
	ListeTriee := PreparationBubbleSort(ListeA, ListeB)
	// Sort
	TriBulle(ListeTriee) // Parameters are transmitted by value but contain a reference to underlying dara when it's an array
	DisplayResults(ListeTriee, "tri bulle")

	// Preparation for Go optimized Sort
	ListeTriee = PreparationBubbleSort(ListeA, ListeB)
	// Go optimized sort
	sort.Strings(ListeTriee)
	DisplayResults(ListeTriee, "tri optimisé de Go")
}

/*

➜  PremierProjet go run main.go
Buenos días señora/señorita/señor

Les deux listes en même temps, sauf la fin de la plus grande
0 01A 03B
1 11A 10B
2 00A 07B
3 02A 05B

Reste de la liste la plus grande
4 64A
5 00B
6 99W

Concaténées
0 01A
1 11A
2 00A
3 02A
4 03B
5 10B
6 07B
7 05B
8 64A
9 00B
10 99W

Tri

Résultat du  tri bulle
0 00A
1 00B
2 01A
3 02A
4 03B
5 05B
6 07B
7 10B
8 11A
9 64A
10 99W

Concaténées
0 01A
1 11A
2 00A
3 02A
4 03B
5 10B
6 07B
7 05B
8 64A
9 00B
10 99W

Résultat du  tri optimisé de Go
0 00A
1 00B
2 01A
3 02A
4 03B
5 05B
6 07B
7 10B
8 11A
9 64A
10 99W

*/
