// Package internal contient la logique métier de l'application
package internal

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

// FizzBuzz implémente l'algorithme fizzbuzz.
// La fonction renvoie une chaine de caractères représentant les nombres de 1 à nb inclus.
// Tous les multiples de 3 sont remplacés par Fizz et les multiples de 5 par Buzz, les multiples de 3 et 5 sont remplacés par FizzBuzz
// Si nb vaut 0, la fonction renvoie une chaine vide
//
// paramètres :
// - nb : Le dernier nombre évalué
//
// retourne :
// La chaine de caractères après application de l'algorithme FizzBuzz
func FizzBuzz(nb uint) string {
	res, _ := FizzBuzzMap(nb, map[uint]string{3 : "Fizz", 5 : "Buzz"})
	return res
}

// FizzBuzzMap implémente l'algorithme fizzbuzz de manière générale
// La fonction renvoie une chaine de caractère contenant les nombres de 1 à nb inclus.
// Si un nombre est divisible par une clé de mots, il est remplacé par le mot correspondant dans mots
// Si un nombre est divisible par plusieurs clés, il est remplacé par la concaténation des mots par ordre croissant des clés.
// Si nb vaut 0, la fonction renvoie une chaine vide.
// Si une des clés de mots est 0, la fonction renvoie une erreur car un nombre ne peut être divisible par 0.
//
// paramètres :
// - nb : Le dernier nombre à évalué
// - mots : Le dictionnaire des mots avec comme clé le diviseur correspondant
//
// retourne :
// La chaine de caractère contenant les nombres évalués avec l'algorithme fizzbuzz
func FizzBuzzMap(nb uint, mots map[uint]string) (string, error) {
	if _, ok := mots[0]; ok {
		return "", errors.New("Clé interdite pour mots")
	}
	cles := make([]uint, 0, len(mots))
	for cle := range mots {
		cles = append(cles, cle)
	}
	slices.Sort(cles)
	res := []string{}
	for i := uint(1); i <= nb; i++ {
		nouvMot := ""
		for _, cle := range cles {
			if i % cle == 0 {
				nouvMot = nouvMot + mots[cle]
			}
		}
		if nouvMot == "" {
			nouvMot = strconv.FormatUint(uint64(i), 10)
		}
		res = append(res, nouvMot)
	}
	return strings.Join(res, " "), nil
}
