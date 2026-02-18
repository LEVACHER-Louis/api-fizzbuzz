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
