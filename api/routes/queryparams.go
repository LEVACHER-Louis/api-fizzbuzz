package routes

import (
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
)

// defaultQueryUInt permet de récupérer un query params naturel.
// Le résultat est def si le paramètre n'est pas défini.
// La fonction renvoie une erreur si le paramètre n'est pas un naturel.
//
// paramètres :
// - c : Le gin.Context de l'appel contenant le paramètre
// - champ : le nom du champ à consulter
// - def : La valeur par défaut si le paramère n'est pas défini
//
// retourne :
// La valeur du paramètre ou la valeur par défaut
func defaultQueryUInt(c *gin.Context, champ string, def uint) (uint, error) {
	valStr := c.Query(champ)
	if valStr == "" {
		return def, nil
	}
	val, conversionOk := strconv.ParseUint(valStr, 10, 64)
	if conversionOk != nil {
		return 0, errors.New("parametre invalide")
	}
	return uint(val), nil
}

// getQueryMapUintString permet de récupérer un query params de uint vers string
// Le résultat sera une map vide nil si le parametre n'est pas présent
// La fonction renvoie une erreur si les cles ne sont pas des naturels
//
// parametres : 
// - c : Le gin.Context de l'appel contenant le parametre
// - champ : Le nom du champ du dictionnaire
//
// retourne : 
// Un map contenant les liens clées valeurs
func getQueryMapUintString(c *gin.Context, champ string) (map[uint]string, error) {
	mapRaw := c.QueryMap(champ)
	if len(mapRaw) == 0 {
		return nil, nil
	}
	res := map[uint]string{}
	for cleRaw, valeur := range mapRaw {
		cle, conversionOk := strconv.ParseUint(cleRaw, 10, 64)
		if conversionOk != nil {
			return nil, errors.New("parametre invalide")
		}
		res[uint(cle)] = valeur
	}
	return res, nil
}
