package internal

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

func FizzBuzz(nb uint) string {
	res, _ := FizzBuzzMap(nb, map[uint]string{3 : "Fizz", 5 : "Buzz"})
	return res
}

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
